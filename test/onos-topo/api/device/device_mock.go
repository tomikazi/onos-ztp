// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/onosproject/onos-topo/api/device (interfaces: DeviceServiceClient,DeviceService_ListClient)

// Package mock_device is a generated GoMock package.
package device

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	device "github.com/onosproject/onos-topo/api/device"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	reflect "reflect"
)

// MockDeviceServiceClient is a mock of DeviceServiceClient interface
type MockDeviceServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockDeviceServiceClientMockRecorder
}

// MockDeviceServiceClientMockRecorder is the mock recorder for MockDeviceServiceClient
type MockDeviceServiceClientMockRecorder struct {
	mock *MockDeviceServiceClient
}

// NewMockDeviceServiceClient creates a new mock instance
func NewMockDeviceServiceClient(ctrl *gomock.Controller) *MockDeviceServiceClient {
	mock := &MockDeviceServiceClient{ctrl: ctrl}
	mock.recorder = &MockDeviceServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDeviceServiceClient) EXPECT() *MockDeviceServiceClientMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockDeviceServiceClient) Add(arg0 context.Context, arg1 *device.AddRequest, arg2 ...grpc.CallOption) (*device.AddResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Add", varargs...)
	ret0, _ := ret[0].(*device.AddResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add
func (mr *MockDeviceServiceClientMockRecorder) Add(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockDeviceServiceClient)(nil).Add), varargs...)
}

// Get mocks base method
func (m *MockDeviceServiceClient) Get(arg0 context.Context, arg1 *device.GetRequest, arg2 ...grpc.CallOption) (*device.GetResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*device.GetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockDeviceServiceClientMockRecorder) Get(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDeviceServiceClient)(nil).Get), varargs...)
}

// List mocks base method
func (m *MockDeviceServiceClient) List(arg0 context.Context, arg1 *device.ListRequest, arg2 ...grpc.CallOption) (device.DeviceService_ListClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(device.DeviceService_ListClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockDeviceServiceClientMockRecorder) List(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockDeviceServiceClient)(nil).List), varargs...)
}

// Remove mocks base method
func (m *MockDeviceServiceClient) Remove(arg0 context.Context, arg1 *device.RemoveRequest, arg2 ...grpc.CallOption) (*device.RemoveResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Remove", varargs...)
	ret0, _ := ret[0].(*device.RemoveResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Remove indicates an expected call of Remove
func (mr *MockDeviceServiceClientMockRecorder) Remove(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockDeviceServiceClient)(nil).Remove), varargs...)
}

// Update mocks base method
func (m *MockDeviceServiceClient) Update(arg0 context.Context, arg1 *device.UpdateRequest, arg2 ...grpc.CallOption) (*device.UpdateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(*device.UpdateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockDeviceServiceClientMockRecorder) Update(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDeviceServiceClient)(nil).Update), varargs...)
}

// MockDeviceService_ListClient is a mock of DeviceService_ListClient interface
type MockDeviceService_ListClient struct {
	ctrl     *gomock.Controller
	recorder *MockDeviceService_ListClientMockRecorder
}

// MockDeviceService_ListClientMockRecorder is the mock recorder for MockDeviceService_ListClient
type MockDeviceService_ListClientMockRecorder struct {
	mock *MockDeviceService_ListClient
}

// NewMockDeviceService_ListClient creates a new mock instance
func NewMockDeviceService_ListClient(ctrl *gomock.Controller) *MockDeviceService_ListClient {
	mock := &MockDeviceService_ListClient{ctrl: ctrl}
	mock.recorder = &MockDeviceService_ListClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDeviceService_ListClient) EXPECT() *MockDeviceService_ListClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method
func (m *MockDeviceService_ListClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockDeviceService_ListClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockDeviceService_ListClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockDeviceService_ListClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockDeviceService_ListClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockDeviceService_ListClient)(nil).Context))
}

// Header mocks base method
func (m *MockDeviceService_ListClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockDeviceService_ListClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockDeviceService_ListClient)(nil).Header))
}

// Recv mocks base method
func (m *MockDeviceService_ListClient) Recv() (*device.ListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*device.ListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockDeviceService_ListClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockDeviceService_ListClient)(nil).Recv))
}

// RecvMsg mocks base method
func (m *MockDeviceService_ListClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockDeviceService_ListClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockDeviceService_ListClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method
func (m *MockDeviceService_ListClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockDeviceService_ListClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockDeviceService_ListClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method
func (m *MockDeviceService_ListClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockDeviceService_ListClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockDeviceService_ListClient)(nil).Trailer))
}
