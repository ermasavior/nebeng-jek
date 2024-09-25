// Code generated by MockGen. DO NOT EDIT.
// Source: internal/rides/usecase/type.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	model "nebeng-jek/internal/rides/model"
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

// SetDriverAvailability mocks base method.
func (m *MockRidesUsecase) SetDriverAvailability(arg0 context.Context, arg1 model.SetDriverAvailabilityRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDriverAvailability", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDriverAvailability indicates an expected call of SetDriverAvailability.
func (mr *MockRidesUsecaseMockRecorder) SetDriverAvailability(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDriverAvailability", reflect.TypeOf((*MockRidesUsecase)(nil).SetDriverAvailability), arg0, arg1)
}