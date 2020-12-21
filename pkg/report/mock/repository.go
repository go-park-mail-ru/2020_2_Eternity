// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_report is a generated GoMock package.
package mock_report

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

// ReportPin mocks base method
func (m *MockIRepository) ReportPin(userId int, rep *domain.ReportReq) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReportPin", userId, rep)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReportPin indicates an expected call of ReportPin
func (mr *MockIRepositoryMockRecorder) ReportPin(userId, rep interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportPin", reflect.TypeOf((*MockIRepository)(nil).ReportPin), userId, rep)
}

// GetReportsByPinId mocks base method
func (m *MockIRepository) GetReportsByPinId(pinId int) ([]domain.Report, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReportsByPinId", pinId)
	ret0, _ := ret[0].([]domain.Report)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReportsByPinId indicates an expected call of GetReportsByPinId
func (mr *MockIRepositoryMockRecorder) GetReportsByPinId(pinId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReportsByPinId", reflect.TypeOf((*MockIRepository)(nil).GetReportsByPinId), pinId)
}

// GetReportsByUsername mocks base method
func (m *MockIRepository) GetReportsByUsername(username string) ([]domain.Report, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReportsByUsername", username)
	ret0, _ := ret[0].([]domain.Report)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReportsByUsername indicates an expected call of GetReportsByUsername
func (mr *MockIRepositoryMockRecorder) GetReportsByUsername(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReportsByUsername", reflect.TypeOf((*MockIRepository)(nil).GetReportsByUsername), username)
}
