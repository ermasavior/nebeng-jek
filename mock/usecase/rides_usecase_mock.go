// Code generated by MockGen. DO NOT EDIT.
// Source: internal/rides/usecase/type.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	model "nebeng-jek/internal/rides/model"
	error "nebeng-jek/pkg/error"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRidesUsecase is a mock of RidesUsecase interface.
type MockRidesUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockRidesUsecaseMockRecorder
}

// MockRidesUsecaseMockRecorder is the mock recorder for MockRidesUsecase.
type MockRidesUsecaseMockRecorder struct {
	mock *MockRidesUsecase
}

// NewMockRidesUsecase creates a new mock instance.
func NewMockRidesUsecase(ctrl *gomock.Controller) *MockRidesUsecase {
	mock := &MockRidesUsecase{ctrl: ctrl}
	mock.recorder = &MockRidesUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRidesUsecase) EXPECT() *MockRidesUsecaseMockRecorder {
	return m.recorder
}

// DriverConfirmPrice mocks base method.
func (m *MockRidesUsecase) DriverConfirmPrice(arg0 context.Context, arg1 model.DriverConfirmPriceRequest) (model.RideData, error.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DriverConfirmPrice", arg0, arg1)
	ret0, _ := ret[0].(model.RideData)
	ret1, _ := ret[1].(error.AppError)
	return ret0, ret1
}

// DriverConfirmPrice indicates an expected call of DriverConfirmPrice.
func (mr *MockRidesUsecaseMockRecorder) DriverConfirmPrice(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DriverConfirmPrice", reflect.TypeOf((*MockRidesUsecase)(nil).DriverConfirmPrice), arg0, arg1)
}

// DriverConfirmRide mocks base method.
func (m *MockRidesUsecase) DriverConfirmRide(arg0 context.Context, arg1 model.DriverConfirmRideRequest) error.AppError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DriverConfirmRide", arg0, arg1)
	ret0, _ := ret[0].(error.AppError)
	return ret0
}

// DriverConfirmRide indicates an expected call of DriverConfirmRide.
func (mr *MockRidesUsecaseMockRecorder) DriverConfirmRide(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DriverConfirmRide", reflect.TypeOf((*MockRidesUsecase)(nil).DriverConfirmRide), arg0, arg1)
}

// DriverEndRide mocks base method.
func (m *MockRidesUsecase) DriverEndRide(arg0 context.Context, arg1 model.DriverEndRideRequest) (model.RideData, error.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DriverEndRide", arg0, arg1)
	ret0, _ := ret[0].(model.RideData)
	ret1, _ := ret[1].(error.AppError)
	return ret0, ret1
}

// DriverEndRide indicates an expected call of DriverEndRide.
func (mr *MockRidesUsecaseMockRecorder) DriverEndRide(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DriverEndRide", reflect.TypeOf((*MockRidesUsecase)(nil).DriverEndRide), arg0, arg1)
}

// DriverSetAvailability mocks base method.
func (m *MockRidesUsecase) DriverSetAvailability(arg0 context.Context, arg1 model.DriverSetAvailabilityRequest) error.AppError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DriverSetAvailability", arg0, arg1)
	ret0, _ := ret[0].(error.AppError)
	return ret0
}

// DriverSetAvailability indicates an expected call of DriverSetAvailability.
func (mr *MockRidesUsecaseMockRecorder) DriverSetAvailability(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DriverSetAvailability", reflect.TypeOf((*MockRidesUsecase)(nil).DriverSetAvailability), arg0, arg1)
}

// DriverStartRide mocks base method.
func (m *MockRidesUsecase) DriverStartRide(arg0 context.Context, arg1 model.DriverStartRideRequest) (model.RideData, error.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DriverStartRide", arg0, arg1)
	ret0, _ := ret[0].(model.RideData)
	ret1, _ := ret[1].(error.AppError)
	return ret0, ret1
}

// DriverStartRide indicates an expected call of DriverStartRide.
func (mr *MockRidesUsecaseMockRecorder) DriverStartRide(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DriverStartRide", reflect.TypeOf((*MockRidesUsecase)(nil).DriverStartRide), arg0, arg1)
}

// RiderConfirmRide mocks base method.
func (m *MockRidesUsecase) RiderConfirmRide(arg0 context.Context, arg1 model.RiderConfirmRideRequest) error.AppError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RiderConfirmRide", arg0, arg1)
	ret0, _ := ret[0].(error.AppError)
	return ret0
}

// RiderConfirmRide indicates an expected call of RiderConfirmRide.
func (mr *MockRidesUsecaseMockRecorder) RiderConfirmRide(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RiderConfirmRide", reflect.TypeOf((*MockRidesUsecase)(nil).RiderConfirmRide), arg0, arg1)
}

// RiderCreateNewRide mocks base method.
func (m *MockRidesUsecase) RiderCreateNewRide(arg0 context.Context, arg1 model.CreateNewRideRequest) (int64, error.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RiderCreateNewRide", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error.AppError)
	return ret0, ret1
}

// RiderCreateNewRide indicates an expected call of RiderCreateNewRide.
func (mr *MockRidesUsecaseMockRecorder) RiderCreateNewRide(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RiderCreateNewRide", reflect.TypeOf((*MockRidesUsecase)(nil).RiderCreateNewRide), arg0, arg1)
}

// TrackUserLocation mocks base method.
func (m *MockRidesUsecase) TrackUserLocation(ctx context.Context, req model.TrackUserLocationRequest) error.AppError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TrackUserLocation", ctx, req)
	ret0, _ := ret[0].(error.AppError)
	return ret0
}

// TrackUserLocation indicates an expected call of TrackUserLocation.
func (mr *MockRidesUsecaseMockRecorder) TrackUserLocation(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrackUserLocation", reflect.TypeOf((*MockRidesUsecase)(nil).TrackUserLocation), ctx, req)
}
