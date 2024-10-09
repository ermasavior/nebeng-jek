// Code generated by MockGen. DO NOT EDIT.
// Source: internal/rides/repository/type.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	model "nebeng-jek/internal/rides/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRidesLocationRepository is a mock of RidesLocationRepository interface.
type MockRidesLocationRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRidesLocationRepositoryMockRecorder
}

// MockRidesLocationRepositoryMockRecorder is the mock recorder for MockRidesLocationRepository.
type MockRidesLocationRepositoryMockRecorder struct {
	mock *MockRidesLocationRepository
}

// NewMockRidesLocationRepository creates a new mock instance.
func NewMockRidesLocationRepository(ctrl *gomock.Controller) *MockRidesLocationRepository {
	mock := &MockRidesLocationRepository{ctrl: ctrl}
	mock.recorder = &MockRidesLocationRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRidesLocationRepository) EXPECT() *MockRidesLocationRepositoryMockRecorder {
	return m.recorder
}

// AddAvailableDriver mocks base method.
func (m *MockRidesLocationRepository) AddAvailableDriver(arg0 context.Context, arg1 string, arg2 model.Coordinate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAvailableDriver", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAvailableDriver indicates an expected call of AddAvailableDriver.
func (mr *MockRidesLocationRepositoryMockRecorder) AddAvailableDriver(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAvailableDriver", reflect.TypeOf((*MockRidesLocationRepository)(nil).AddAvailableDriver), arg0, arg1, arg2)
}

// GetNearestAvailableDrivers mocks base method.
func (m *MockRidesLocationRepository) GetNearestAvailableDrivers(arg0 context.Context, arg1 model.Coordinate) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNearestAvailableDrivers", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNearestAvailableDrivers indicates an expected call of GetNearestAvailableDrivers.
func (mr *MockRidesLocationRepositoryMockRecorder) GetNearestAvailableDrivers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNearestAvailableDrivers", reflect.TypeOf((*MockRidesLocationRepository)(nil).GetNearestAvailableDrivers), arg0, arg1)
}

// GetRidePath mocks base method.
func (m *MockRidesLocationRepository) GetRidePath(arg0 context.Context, arg1 int64, arg2 string) ([]model.Coordinate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRidePath", arg0, arg1, arg2)
	ret0, _ := ret[0].([]model.Coordinate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRidePath indicates an expected call of GetRidePath.
func (mr *MockRidesLocationRepositoryMockRecorder) GetRidePath(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRidePath", reflect.TypeOf((*MockRidesLocationRepository)(nil).GetRidePath), arg0, arg1, arg2)
}

// RemoveAvailableDriver mocks base method.
func (m *MockRidesLocationRepository) RemoveAvailableDriver(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAvailableDriver", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAvailableDriver indicates an expected call of RemoveAvailableDriver.
func (mr *MockRidesLocationRepositoryMockRecorder) RemoveAvailableDriver(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAvailableDriver", reflect.TypeOf((*MockRidesLocationRepository)(nil).RemoveAvailableDriver), arg0, arg1)
}

// MockRidesRepository is a mock of RidesRepository interface.
type MockRidesRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRidesRepositoryMockRecorder
}

// MockRidesRepositoryMockRecorder is the mock recorder for MockRidesRepository.
type MockRidesRepositoryMockRecorder struct {
	mock *MockRidesRepository
}

// NewMockRidesRepository creates a new mock instance.
func NewMockRidesRepository(ctrl *gomock.Controller) *MockRidesRepository {
	mock := &MockRidesRepository{ctrl: ctrl}
	mock.recorder = &MockRidesRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRidesRepository) EXPECT() *MockRidesRepositoryMockRecorder {
	return m.recorder
}

// ConfirmRideDriver mocks base method.
func (m *MockRidesRepository) ConfirmRideDriver(ctx context.Context, req model.ConfirmRideDriverRequest) (model.RideData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfirmRideDriver", ctx, req)
	ret0, _ := ret[0].(model.RideData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfirmRideDriver indicates an expected call of ConfirmRideDriver.
func (mr *MockRidesRepositoryMockRecorder) ConfirmRideDriver(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfirmRideDriver", reflect.TypeOf((*MockRidesRepository)(nil).ConfirmRideDriver), ctx, req)
}

// ConfirmRideRider mocks base method.
func (m *MockRidesRepository) ConfirmRideRider(ctx context.Context, req model.ConfirmRideRiderRequest) (model.RideData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfirmRideRider", ctx, req)
	ret0, _ := ret[0].(model.RideData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfirmRideRider indicates an expected call of ConfirmRideRider.
func (mr *MockRidesRepositoryMockRecorder) ConfirmRideRider(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfirmRideRider", reflect.TypeOf((*MockRidesRepository)(nil).ConfirmRideRider), ctx, req)
}

// CreateNewRide mocks base method.
func (m *MockRidesRepository) CreateNewRide(arg0 context.Context, arg1 model.CreateNewRideRequest) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewRide", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNewRide indicates an expected call of CreateNewRide.
func (mr *MockRidesRepositoryMockRecorder) CreateNewRide(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewRide", reflect.TypeOf((*MockRidesRepository)(nil).CreateNewRide), arg0, arg1)
}

// GetDriverDataByMSISDN mocks base method.
func (m *MockRidesRepository) GetDriverDataByMSISDN(ctx context.Context, msisdn string) (model.DriverData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDriverDataByMSISDN", ctx, msisdn)
	ret0, _ := ret[0].(model.DriverData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDriverDataByMSISDN indicates an expected call of GetDriverDataByMSISDN.
func (mr *MockRidesRepositoryMockRecorder) GetDriverDataByMSISDN(ctx, msisdn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDriverDataByMSISDN", reflect.TypeOf((*MockRidesRepository)(nil).GetDriverDataByMSISDN), ctx, msisdn)
}

// GetDriverMSISDNByID mocks base method.
func (m *MockRidesRepository) GetDriverMSISDNByID(ctx context.Context, id int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDriverMSISDNByID", ctx, id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDriverMSISDNByID indicates an expected call of GetDriverMSISDNByID.
func (mr *MockRidesRepositoryMockRecorder) GetDriverMSISDNByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDriverMSISDNByID", reflect.TypeOf((*MockRidesRepository)(nil).GetDriverMSISDNByID), ctx, id)
}

// GetRiderDataByMSISDN mocks base method.
func (m *MockRidesRepository) GetRiderDataByMSISDN(ctx context.Context, msisdn string) (model.RiderData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRiderDataByMSISDN", ctx, msisdn)
	ret0, _ := ret[0].(model.RiderData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRiderDataByMSISDN indicates an expected call of GetRiderDataByMSISDN.
func (mr *MockRidesRepositoryMockRecorder) GetRiderDataByMSISDN(ctx, msisdn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRiderDataByMSISDN", reflect.TypeOf((*MockRidesRepository)(nil).GetRiderDataByMSISDN), ctx, msisdn)
}

// GetRiderMSISDNByID mocks base method.
func (m *MockRidesRepository) GetRiderMSISDNByID(ctx context.Context, id int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRiderMSISDNByID", ctx, id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRiderMSISDNByID indicates an expected call of GetRiderMSISDNByID.
func (mr *MockRidesRepositoryMockRecorder) GetRiderMSISDNByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRiderMSISDNByID", reflect.TypeOf((*MockRidesRepository)(nil).GetRiderMSISDNByID), ctx, id)
}

// UpdateRideByDriver mocks base method.
func (m *MockRidesRepository) UpdateRideByDriver(ctx context.Context, req model.UpdateRideByDriverRequest) (model.RideData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRideByDriver", ctx, req)
	ret0, _ := ret[0].(model.RideData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRideByDriver indicates an expected call of UpdateRideByDriver.
func (mr *MockRidesRepositoryMockRecorder) UpdateRideByDriver(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRideByDriver", reflect.TypeOf((*MockRidesRepository)(nil).UpdateRideByDriver), ctx, req)
}

// MockRidesPubsubRepository is a mock of RidesPubsubRepository interface.
type MockRidesPubsubRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRidesPubsubRepositoryMockRecorder
}

// MockRidesPubsubRepositoryMockRecorder is the mock recorder for MockRidesPubsubRepository.
type MockRidesPubsubRepositoryMockRecorder struct {
	mock *MockRidesPubsubRepository
}

// NewMockRidesPubsubRepository creates a new mock instance.
func NewMockRidesPubsubRepository(ctrl *gomock.Controller) *MockRidesPubsubRepository {
	mock := &MockRidesPubsubRepository{ctrl: ctrl}
	mock.recorder = &MockRidesPubsubRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRidesPubsubRepository) EXPECT() *MockRidesPubsubRepositoryMockRecorder {
	return m.recorder
}

// BroadcastMessage mocks base method.
func (m *MockRidesPubsubRepository) BroadcastMessage(arg0 context.Context, arg1 string, arg2 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BroadcastMessage", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// BroadcastMessage indicates an expected call of BroadcastMessage.
func (mr *MockRidesPubsubRepositoryMockRecorder) BroadcastMessage(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BroadcastMessage", reflect.TypeOf((*MockRidesPubsubRepository)(nil).BroadcastMessage), arg0, arg1, arg2)
}
