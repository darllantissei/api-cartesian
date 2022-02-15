// Code generated by MockGen. DO NOT EDIT.
// Source: /home/darllan/go/src/github.com/darllantissei/api-cartesian/application/coordinate/interface.go

// Package mock_coordinate is a generated GoMock package.
package mock_coordinate

import (
	reflect "reflect"

	models "github.com/darllantissei/api-cartesian/application/models"
	gomock "github.com/golang/mock/gomock"
)

// MockICoordinateService is a mock of ICoordinateService interface.
type MockICoordinateService struct {
	ctrl     *gomock.Controller
	recorder *MockICoordinateServiceMockRecorder
}

// MockICoordinateServiceMockRecorder is the mock recorder for MockICoordinateService.
type MockICoordinateServiceMockRecorder struct {
	mock *MockICoordinateService
}

// NewMockICoordinateService creates a new mock instance.
func NewMockICoordinateService(ctrl *gomock.Controller) *MockICoordinateService {
	mock := &MockICoordinateService{ctrl: ctrl}
	mock.recorder = &MockICoordinateServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICoordinateService) EXPECT() *MockICoordinateServiceMockRecorder {
	return m.recorder
}

// Proccess mocks base method.
func (m *MockICoordinateService) Proccess(coordX, coordY, distance int64) ([]models.Way, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Proccess", coordX, coordY, distance)
	ret0, _ := ret[0].([]models.Way)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Proccess indicates an expected call of Proccess.
func (mr *MockICoordinateServiceMockRecorder) Proccess(coordX, coordY, distance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Proccess", reflect.TypeOf((*MockICoordinateService)(nil).Proccess), coordX, coordY, distance)
}

// MockICoordinateReaderFile is a mock of ICoordinateReaderFile interface.
type MockICoordinateReaderFile struct {
	ctrl     *gomock.Controller
	recorder *MockICoordinateReaderFileMockRecorder
}

// MockICoordinateReaderFileMockRecorder is the mock recorder for MockICoordinateReaderFile.
type MockICoordinateReaderFileMockRecorder struct {
	mock *MockICoordinateReaderFile
}

// NewMockICoordinateReaderFile creates a new mock instance.
func NewMockICoordinateReaderFile(ctrl *gomock.Controller) *MockICoordinateReaderFile {
	mock := &MockICoordinateReaderFile{ctrl: ctrl}
	mock.recorder = &MockICoordinateReaderFileMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICoordinateReaderFile) EXPECT() *MockICoordinateReaderFileMockRecorder {
	return m.recorder
}

// ListPoints mocks base method.
func (m *MockICoordinateReaderFile) ListPoints() (models.Points, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPoints")
	ret0, _ := ret[0].(models.Points)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPoints indicates an expected call of ListPoints.
func (mr *MockICoordinateReaderFileMockRecorder) ListPoints() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPoints", reflect.TypeOf((*MockICoordinateReaderFile)(nil).ListPoints))
}

// MockICoordinatePersistenceFile is a mock of ICoordinatePersistenceFile interface.
type MockICoordinatePersistenceFile struct {
	ctrl     *gomock.Controller
	recorder *MockICoordinatePersistenceFileMockRecorder
}

// MockICoordinatePersistenceFileMockRecorder is the mock recorder for MockICoordinatePersistenceFile.
type MockICoordinatePersistenceFileMockRecorder struct {
	mock *MockICoordinatePersistenceFile
}

// NewMockICoordinatePersistenceFile creates a new mock instance.
func NewMockICoordinatePersistenceFile(ctrl *gomock.Controller) *MockICoordinatePersistenceFile {
	mock := &MockICoordinatePersistenceFile{ctrl: ctrl}
	mock.recorder = &MockICoordinatePersistenceFileMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICoordinatePersistenceFile) EXPECT() *MockICoordinatePersistenceFileMockRecorder {
	return m.recorder
}

// ListPoints mocks base method.
func (m *MockICoordinatePersistenceFile) ListPoints() (models.Points, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPoints")
	ret0, _ := ret[0].(models.Points)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPoints indicates an expected call of ListPoints.
func (mr *MockICoordinatePersistenceFileMockRecorder) ListPoints() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPoints", reflect.TypeOf((*MockICoordinatePersistenceFile)(nil).ListPoints))
}
