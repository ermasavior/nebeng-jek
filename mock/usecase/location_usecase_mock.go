// Code generated by MockGen. DO NOT EDIT.
// Source: internal/location/usecase/type.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	model "nebeng-jek/internal/location/model"
	location "nebeng-jek/internal/pkg/location"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLocationUsecase is a mock of LocationUsecase interface.
type MockLocationUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockLocationUsecaseMockRecorder
}

// MockLocationUsecaseMockRecorder is the mock recorder for MockLocationUsecase.
type MockLocationUsecaseMockRecorder struct {
	mock *MockLocationUsecase
}

// NewMockLocationUsecase creates a new mock instance.
func NewMockLocationUsecase(ctrl *gomock.Controller) *MockLocationUsecase {
	mock := &MockLocationUsecase{ctrl: ctrl}
	mock.recorder = &MockLocationUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLocationUsecase) EXPECT() *MockLocationUsecaseMockRecorder {
	return m.recorder
}

// AddAvailableDriver mocks base method.
func (m *MockLocationUsecase) AddAvailableDriver(ctx context.Context, driverID int64, location location.Coordinate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAvailableDriver", ctx, driverID, location)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAvailableDriver indicates an expected call of AddAvailableDriver.
func (mr *MockLocationUsecaseMockRecorder) AddAvailableDriver(ctx, driverID, location interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAvailableDriver", reflect.TypeOf((*MockLocationUsecase)(nil).AddAvailableDriver), ctx, driverID, location)
}

// GetNearestAvailableDrivers mocks base method.
func (m *MockLocationUsecase) GetNearestAvailableDrivers(ctx context.Context, location location.Coordinate) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNearestAvailableDrivers", ctx, location)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNearestAvailableDrivers indicates an expected call of GetNearestAvailableDrivers.
func (mr *MockLocationUsecaseMockRecorder) GetNearestAvailableDrivers(ctx, location interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNearestAvailableDrivers", reflect.TypeOf((*MockLocationUsecase)(nil).GetNearestAvailableDrivers), ctx, location)
}

// GetRidePath mocks base method.
func (m *MockLocationUsecase) GetRidePath(ctx context.Context, rideID, driverID int64) ([]location.Coordinate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRidePath", ctx, rideID, driverID)
	ret0, _ := ret[0].([]location.Coordinate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRidePath indicates an expected call of GetRidePath.
func (mr *MockLocationUsecaseMockRecorder) GetRidePath(ctx, rideID, driverID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRidePath", reflect.TypeOf((*MockLocationUsecase)(nil).GetRidePath), ctx, rideID, driverID)
}

// RemoveAvailableDriver mocks base method.
func (m *MockLocationUsecase) RemoveAvailableDriver(ctx context.Context, driverID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAvailableDriver", ctx, driverID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAvailableDriver indicates an expected call of RemoveAvailableDriver.
func (mr *MockLocationUsecaseMockRecorder) RemoveAvailableDriver(ctx, driverID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAvailableDriver", reflect.TypeOf((*MockLocationUsecase)(nil).RemoveAvailableDriver), ctx, driverID)
}

// TrackUserLocation mocks base method.
func (m *MockLocationUsecase) TrackUserLocation(ctx context.Context, req model.TrackUserLocationRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TrackUserLocation", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// TrackUserLocation indicates an expected call of TrackUserLocation.
func (mr *MockLocationUsecaseMockRecorder) TrackUserLocation(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrackUserLocation", reflect.TypeOf((*MockLocationUsecase)(nil).TrackUserLocation), ctx, req)
}
