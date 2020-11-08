// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_pin is a generated GoMock package.
package mock_pin

import (
	domain "github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	gomock "github.com/golang/mock/gomock"
	multipart "mime/multipart"
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

// StorePin mocks base method
func (m *MockIRepository) StorePin(p *domain.Pin) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorePin", p)
	ret0, _ := ret[0].(error)
	return ret0
}

// StorePin indicates an expected call of StorePin
func (mr *MockIRepositoryMockRecorder) StorePin(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorePin", reflect.TypeOf((*MockIRepository)(nil).StorePin), p)
}

// GetPin mocks base method
func (m *MockIRepository) GetPin(id int) (domain.Pin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPin", id)
	ret0, _ := ret[0].(domain.Pin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPin indicates an expected call of GetPin
func (mr *MockIRepositoryMockRecorder) GetPin(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPin", reflect.TypeOf((*MockIRepository)(nil).GetPin), id)
}

// GetPinList mocks base method
func (m *MockIRepository) GetPinList(username string) ([]domain.Pin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPinList", username)
	ret0, _ := ret[0].([]domain.Pin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPinList indicates an expected call of GetPinList
func (mr *MockIRepositoryMockRecorder) GetPinList(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPinList", reflect.TypeOf((*MockIRepository)(nil).GetPinList), username)
}

// GetPinBoardList mocks base method
func (m *MockIRepository) GetPinBoardList(boardId int) ([]domain.Pin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPinBoardList", boardId)
	ret0, _ := ret[0].([]domain.Pin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPinBoardList indicates an expected call of GetPinBoardList
func (mr *MockIRepositoryMockRecorder) GetPinBoardList(boardId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPinBoardList", reflect.TypeOf((*MockIRepository)(nil).GetPinBoardList), boardId)
}

// MockIStorage is a mock of IStorage interface
type MockIStorage struct {
	ctrl     *gomock.Controller
	recorder *MockIStorageMockRecorder
}

// MockIStorageMockRecorder is the mock recorder for MockIStorage
type MockIStorageMockRecorder struct {
	mock *MockIStorage
}

// NewMockIStorage creates a new mock instance
func NewMockIStorage(ctrl *gomock.Controller) *MockIStorage {
	mock := &MockIStorage{ctrl: ctrl}
	mock.recorder = &MockIStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIStorage) EXPECT() *MockIStorageMockRecorder {
	return m.recorder
}

// SaveUploadedFile mocks base method
func (m *MockIStorage) SaveUploadedFile(file *multipart.FileHeader, filename string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveUploadedFile", file, filename)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveUploadedFile indicates an expected call of SaveUploadedFile
func (mr *MockIStorageMockRecorder) SaveUploadedFile(file, filename interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveUploadedFile", reflect.TypeOf((*MockIStorage)(nil).SaveUploadedFile), file, filename)
}
