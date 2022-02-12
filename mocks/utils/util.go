// Code generated by MockGen. DO NOT EDIT.
// Source: /home/darllan/go/src/github.com/darllantissei/api-cartesian/application/utils/interface.go

// Package mock_utils is a generated GoMock package.
package mock_utils

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIUtilsService is a mock of IUtilsService interface.
type MockIUtilsService struct {
	ctrl     *gomock.Controller
	recorder *MockIUtilsServiceMockRecorder
}

// MockIUtilsServiceMockRecorder is the mock recorder for MockIUtilsService.
type MockIUtilsServiceMockRecorder struct {
	mock *MockIUtilsService
}

// NewMockIUtilsService creates a new mock instance.
func NewMockIUtilsService(ctrl *gomock.Controller) *MockIUtilsService {
	mock := &MockIUtilsService{ctrl: ctrl}
	mock.recorder = &MockIUtilsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUtilsService) EXPECT() *MockIUtilsServiceMockRecorder {
	return m.recorder
}

// FileExists mocks base method.
func (m *MockIUtilsService) FileExists(fileName string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FileExists", fileName)
	ret0, _ := ret[0].(bool)
	return ret0
}

// FileExists indicates an expected call of FileExists.
func (mr *MockIUtilsServiceMockRecorder) FileExists(fileName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FileExists", reflect.TypeOf((*MockIUtilsService)(nil).FileExists), fileName)
}