// Copyright 2019-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package roles :
package roles

import (
	"context"
	"sync"

	"github.com/onosproject/onos-lib-go/pkg/logging"
	"github.com/onosproject/onos-lib-go/pkg/northbound"
	"github.com/onosproject/onos-ztp/api/admin"
	"github.com/onosproject/onos-ztp/pkg/manager"
	"google.golang.org/grpc"
)

var log = logging.GetLogger("roles")

// Service is a Service implementation for administration.
type Service struct {
	northbound.Service
}

// Register registers the Service with the gRPC server.
func (s Service) Register(r *grpc.Server) {
	server := &Server{}
	go server.dispatch()
	admin.RegisterDeviceRoleServiceServer(r, server)
}

// Server implements the gRPC service for zero-touch provisioning facilities.
type Server struct {
	dispatcherStarted bool
	subscriberCount   int
	subscribers       sync.Map
}

func (s *Server) register(changes *chan admin.DeviceRoleChange) int {
	// TODO: add hashing or mutex
	s.subscriberCount++
	key := s.subscriberCount
	s.subscribers.Store(key, changes)
	return key
}

func (s *Server) unregister(key int) {
	s.subscribers.Delete(key)
}

// Set provides means to add, update or delete device role configuration.
func (s *Server) Set(ctx context.Context, r *admin.DeviceRoleChangeRequest) (*admin.DeviceRoleChangeResponse, error) {
	var err error
	var cfg = r.GetConfig()
	var changeType = admin.DeviceRoleChange_UPDATED

	switch r.GetChange() {
	case admin.DeviceRoleChangeRequest_ADD:
		changeType = admin.DeviceRoleChange_ADDED
		log.Infof("Adding new role %s", cfg.GetRole())
		err = manager.GetManager().RoleStore.WriteRole(cfg, false)
	case admin.DeviceRoleChangeRequest_UPDATE:
		log.Infof("Updating role %s", cfg.GetRole())
		err = manager.GetManager().RoleStore.WriteRole(cfg, true)
	case admin.DeviceRoleChangeRequest_DELETE:
		changeType = admin.DeviceRoleChange_DELETED
		log.Infof("Removing role %s", cfg.GetRole())
		cfg, err = manager.GetManager().RoleStore.DeleteRole(cfg.Role)
	}

	if err != nil {
		return nil, err
	}

	change := admin.DeviceRoleChange{
		Change: changeType,
		Config: cfg,
	}
	log.Infof("Responding to request with %v", change)

	// Queue up device role change for subscribers
	manager.GetManager().ChangesChannel <- change

	// ...and return it to the caller
	return &admin.DeviceRoleChangeResponse{
		Change: &change,
	}, nil
}

// Get provides means to query device role configuration.
func (s *Server) Get(req *admin.DeviceRoleRequest, stream admin.DeviceRoleService_GetServer) error {
	roleName := req.GetRole()
	if len(roleName) > 0 {
		return s.sendRoleConfig(roleName, stream)
	}

	roles, err := manager.GetManager().RoleStore.ListRoles()
	for i := 0; err == nil && i < len(roles); i++ {
		err = s.sendRoleConfig(roles[i], stream)
	}
	return err
}

func (s *Server) sendRoleConfig(roleName string, stream admin.DeviceRoleService_GetServer) error {
	role, err := manager.GetManager().RoleStore.ReadRole(roleName)
	if err == nil {
		err = stream.Send(role)
	}
	return err
}

// Subscribe provides means to monitor changes in the device role configuration.
func (s *Server) Subscribe(req *admin.DeviceRoleRequest, stream admin.DeviceRoleService_SubscribeServer) error {
	// Create and register a channel on which to receive notifications
	changeChan := make(chan admin.DeviceRoleChange)
	key := s.register(&changeChan)
	defer s.unregister(key)

	// ... then go to listen on it
	for {
		event, ok := <-changeChan
		if !ok {
			break
		}
		err := stream.Send(&event)
		if err != nil {
			log.Error("Unable to send notification for", event)
			break
		}
	}
	return nil
}

func (s *Server) dispatch() {
	changes := manager.GetManager().ChangesChannel
	log.Info("Dispatcher started")
	for {
		change, ok := <-changes
		if !ok {
			break
		}
		s.subscribers.Range(func(key, sub interface{}) bool {
			ch := sub.(*chan admin.DeviceRoleChange)
			*ch <- change
			return true
		})
	}
	log.Info("Dispatcher terminated")
	s.dispatcherStarted = false
}
