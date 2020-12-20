// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_search is a generated GoMock package.
package mock_search

import (
	domain "github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIRepository is a mock of IRepository interface
type MockIRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIRepositoryMockRecorder
}

// MockIRepositoryMockRecorder is the mock recorder for MockIRepository
type MockIRepositoryMockRecorder struct {
	mock *MockIRepository
}

// NewMockIRepository creates a new mock instance
func NewMockIRepository(ctrl *gomock.Controller) *MockIRepository {
	mock := &MockIRepository{ctrl: ctrl}
	mock.recorder = &MockIRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIRepository) EXPECT() *MockIRepositoryMockRecorder {
	return m.recorder
}

// GetUsersByName mocks base method
func (m *MockIRepository) GetUsersByName(username string, last int) ([]domain.UserSearch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersByName", username, last)
	ret0, _ := ret[0].([]domain.UserSearch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersByName indicates an expected call of GetUsersByName
func (mr *MockIRepositoryMockRecorder) GetUsersByName(username, last interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersByName", reflect.TypeOf((*MockIRepository)(nil).GetUsersByName), username, last)
}

// GetPinsByTitle mocks base method
func (m *MockIRepository) GetPinsByTitle(title string, last int) ([]domain.Pin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPinsByTitle", title, last)
	ret0, _ := ret[0].([]domain.Pin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPinsByTitle indicates an expected call of GetPinsByTitle
func (mr *MockIRepositoryMockRecorder) GetPinsByTitle(title, last interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPinsByTitle", reflect.TypeOf((*MockIRepository)(nil).GetPinsByTitle), title, last)
}

// GetBoardsByTitle mocks base method
func (m *MockIRepository) GetBoardsByTitle(title string, last int) ([]domain.Board, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBoardsByTitle", title, last)
	ret0, _ := ret[0].([]domain.Board)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBoardsByTitle indicates an expected call of GetBoardsByTitle
func (mr *MockIRepositoryMockRecorder) GetBoardsByTitle(title, last interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBoardsByTitle", reflect.TypeOf((*MockIRepository)(nil).GetBoardsByTitle), title, last)
}
