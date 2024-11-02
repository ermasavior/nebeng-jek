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
func (m *MockRidesLocationRepository) AddAvailableDriver(ctx context.Context, driverID int64, location model.Coordinate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAvailableDriver", ctx, driverID, location)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAvailableDriver indicates an expected call of AddAvailableDriver.
func (mr *MockRidesLocationRepositoryMockRecorder) AddAvailableDriver(ctx, driverID, location interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAvailableDriver", reflect.TypeOf((*MockRidesLocationRepository)(nil).AddAvailableDriver), ctx, driverID, location)
}

// GetNearestAvailableDrivers mocks base method.
func (m *MockRidesLocationRepository) GetNearestAvailableDrivers(ctx context.Context, location model.Coordinate) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNearestAvailableDrivers", ctx, location)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNearestAvailableDrivers indicates an expected call of GetNearestAvailableDrivers.
func (mr *MockRidesLocationRepositoryMockRecorder) GetNearestAvailableDrivers(ctx, location interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNearestAvailableDrivers", reflect.TypeOf((*MockRidesLocationRepository)(nil).GetNearestAvailableDrivers), ctx, location)
}

// GetRidePath mocks base method.
func (m *MockRidesLocationRepository) GetRidePath(ctx context.Context, rideID, driverID int64) ([]model.Coordinate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRidePath", ctx, rideID, driverID)
	ret0, _ := ret[0].([]model.Coordinate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRidePath indicates an expected call of GetRidePath.
func (mr *MockRidesLocationRepositoryMockRecorder) GetRidePath(ctx, rideID, driverID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRidePath", reflect.TypeOf((*MockRidesLocationRepository)(nil).GetRidePath), ctx, rideID, driverID)
}

// RemoveAvailableDriver mocks base method.
func (m *MockRidesLocationRepository) RemoveAvailableDriver(ctx context.Context, driverID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAvailableDriver", ctx, driverID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAvailableDriver indicates an expected call of RemoveAvailableDriver.
func (mr *MockRidesLocationRepositoryMockRecorder) RemoveAvailableDriver(ctx, driverID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAvailableDriver", reflect.TypeOf((*MockRidesLocationRepository)(nil).RemoveAvailableDriver), ctx, driverID)
}

// TrackUserLocation mocks base method.
func (m *MockRidesLocationRepository) TrackUserLocation(ctx context.Context, req model.TrackUserLocationRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TrackUserLocation", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// TrackUserLocation indicates an expected call of TrackUserLocation.
func (mr *MockRidesLocationRepositoryMockRecorder) TrackUserLocation(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TrackUserLocation", reflect.TypeOf((*MockRidesLocationRepository)(nil).TrackUserLocation), ctx, req)
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

// GetDriverDataByID mocks base method.
func (m *MockRidesRepository) GetDriverDataByID(ctx context.Context, driverID int64) (model.DriverData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDriverDataByID", ctx, driverID)
	ret0, _ := ret[0].(model.DriverData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDriverDataByID indicates an expected call of GetDriverDataByID.
func (mr *MockRidesRepositoryMockRecorder) GetDriverDataByID(ctx, driverID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDriverDataByID", reflect.TypeOf((*MockRidesRepository)(nil).GetDriverDataByID), ctx, driverID)
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

// GetRideData mocks base method.
func (m *MockRidesRepository) GetRideData(ctx context.Context, id int64) (model.RideData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRideData", ctx, id)
	ret0, _ := ret[0].(model.RideData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRideData indicates an expected call of GetRideData.
func (mr *MockRidesRepositoryMockRecorder) GetRideData(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRideData", reflect.TypeOf((*MockRidesRepository)(nil).GetRideData), ctx, id)
}

// GetRiderDataByID mocks base method.
func (m *MockRidesRepository) GetRiderDataByID(ctx context.Context, riderID int64) (model.RiderData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRiderDataByID", ctx, riderID)
	ret0, _ := ret[0].(model.RiderData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRiderDataByID indicates an expected call of GetRiderDataByID.
func (mr *MockRidesRepositoryMockRecorder) GetRiderDataByID(ctx, riderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRiderDataByID", reflect.TypeOf((*MockRidesRepository)(nil).GetRiderDataByID), ctx, riderID)
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

// StoreRideCommission mocks base method.
func (m *MockRidesRepository) StoreRideCommission(ctx context.Context, req model.StoreRideCommissionRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreRideCommission", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreRideCommission indicates an expected call of StoreRideCommission.
func (mr *MockRidesRepositoryMockRecorder) StoreRideCommission(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreRideCommission", reflect.TypeOf((*MockRidesRepository)(nil).StoreRideCommission), ctx, req)
}

// UpdateRideData mocks base method.
func (m *MockRidesRepository) UpdateRideData(ctx context.Context, req model.UpdateRideDataRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRideData", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRideData indicates an expected call of UpdateRideData.
func (mr *MockRidesRepositoryMockRecorder) UpdateRideData(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRideData", reflect.TypeOf((*MockRidesRepository)(nil).UpdateRideData), ctx, req)
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

// MockPaymentRepository is a mock of PaymentRepository interface.
type MockPaymentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPaymentRepositoryMockRecorder
}

// MockPaymentRepositoryMockRecorder is the mock recorder for MockPaymentRepository.
type MockPaymentRepositoryMockRecorder struct {
	mock *MockPaymentRepository
}

// NewMockPaymentRepository creates a new mock instance.
func NewMockPaymentRepository(ctrl *gomock.Controller) *MockPaymentRepository {
	mock := &MockPaymentRepository{ctrl: ctrl}
	mock.recorder = &MockPaymentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPaymentRepository) EXPECT() *MockPaymentRepositoryMockRecorder {
	return m.recorder
}

// AddCredit mocks base method.
func (m *MockPaymentRepository) AddCredit(arg0 context.Context, arg1 model.AddCreditRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCredit", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCredit indicates an expected call of AddCredit.
func (mr *MockPaymentRepositoryMockRecorder) AddCredit(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCredit", reflect.TypeOf((*MockPaymentRepository)(nil).AddCredit), arg0, arg1)
}

// DeductCredit mocks base method.
func (m *MockPaymentRepository) DeductCredit(arg0 context.Context, arg1 model.DeductCreditRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeductCredit", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeductCredit indicates an expected call of DeductCredit.
func (mr *MockPaymentRepositoryMockRecorder) DeductCredit(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeductCredit", reflect.TypeOf((*MockPaymentRepository)(nil).DeductCredit), arg0, arg1)
}
