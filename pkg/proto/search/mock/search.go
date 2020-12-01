// Code generated by MockGen. DO NOT EDIT.
// Source: search.pb.go

// Package mock_search is a generated GoMock package.
package mock_search

import (
	context "context"
	search "github.com/go-park-mail-ru/2020_2_Eternity/pkg/proto/search"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockSearchServiceClient is a mock of SearchServiceClient interface
type MockSearchServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockSearchServiceClientMockRecorder
}

// MockSearchServiceClientMockRecorder is the mock recorder for MockSearchServiceClient
type MockSearchServiceClientMockRecorder struct {
	mock *MockSearchServiceClient
}

// NewMockSearchServiceClient creates a new mock instance
func NewMockSearchServiceClient(ctrl *gomock.Controller) *MockSearchServiceClient {
	mock := &MockSearchServiceClient{ctrl: ctrl}
	mock.recorder = &MockSearchServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSearchServiceClient) EXPECT() *MockSearchServiceClientMockRecorder {
	return m.recorder
}

// GetUsersByName mocks base method
func (m *MockSearchServiceClient) GetUsersByName(ctx context.Context, in *search.UserSearch, opts ...grpc.CallOption) (*search.Users, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUsersByName", varargs...)
	ret0, _ := ret[0].(*search.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersByName indicates an expected call of GetUsersByName
func (mr *MockSearchServiceClientMockRecorder) GetUsersByName(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersByName", reflect.TypeOf((*MockSearchServiceClient)(nil).GetUsersByName), varargs...)
}

// GetPinsByTitle mocks base method
func (m *MockSearchServiceClient) GetPinsByTitle(ctx context.Context, in *search.PinSearch, opts ...grpc.CallOption) (*search.Pins, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPinsByTitle", varargs...)
	ret0, _ := ret[0].(*search.Pins)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPinsByTitle indicates an expected call of GetPinsByTitle
func (mr *MockSearchServiceClientMockRecorder) GetPinsByTitle(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPinsByTitle", reflect.TypeOf((*MockSearchServiceClient)(nil).GetPinsByTitle), varargs...)
}

// MockSearchServiceServer is a mock of SearchServiceServer interface
type MockSearchServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockSearchServiceServerMockRecorder
}

// MockSearchServiceServerMockRecorder is the mock recorder for MockSearchServiceServer
type MockSearchServiceServerMockRecorder struct {
	mock *MockSearchServiceServer
}

// NewMockSearchServiceServer creates a new mock instance
func NewMockSearchServiceServer(ctrl *gomock.Controller) *MockSearchServiceServer {
	mock := &MockSearchServiceServer{ctrl: ctrl}
	mock.recorder = &MockSearchServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSearchServiceServer) EXPECT() *MockSearchServiceServerMockRecorder {
	return m.recorder
}

// GetUsersByName mocks base method
func (m *MockSearchServiceServer) GetUsersByName(arg0 context.Context, arg1 *search.UserSearch) (*search.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersByName", arg0, arg1)
	ret0, _ := ret[0].(*search.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersByName indicates an expected call of GetUsersByName
func (mr *MockSearchServiceServerMockRecorder) GetUsersByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersByName", reflect.TypeOf((*MockSearchServiceServer)(nil).GetUsersByName), arg0, arg1)
}

// GetPinsByTitle mocks base method
func (m *MockSearchServiceServer) GetPinsByTitle(arg0 context.Context, arg1 *search.PinSearch) (*search.Pins, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPinsByTitle", arg0, arg1)
	ret0, _ := ret[0].(*search.Pins)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPinsByTitle indicates an expected call of GetPinsByTitle
func (mr *MockSearchServiceServerMockRecorder) GetPinsByTitle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPinsByTitle", reflect.TypeOf((*MockSearchServiceServer)(nil).GetPinsByTitle), arg0, arg1)
}
