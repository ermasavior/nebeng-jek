package usecase

import (
	"context"
	"errors"
	"testing"

	"nebeng-jek/internal/pkg/constants"
	pkgContext "nebeng-jek/internal/pkg/context"
	pkgLocation "nebeng-jek/internal/pkg/location"
	"nebeng-jek/internal/rides/model"
	mockRepo "nebeng-jek/mock/repository"
	pkgError "nebeng-jek/pkg/error"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUsecase_DriverSetAvailability(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoLocMock := mockRepo.NewMockRidesLocationRepository(ctrl)
	repoRidesMock := mockRepo.NewMockRidesRepository(ctrl)
	usecaseMock := NewUsecase(repoLocMock, repoRidesMock, nil, nil)

	var (
		driverID     = int64(1111)
		driverMSISDN = "081111"
		req          = model.DriverSetAvailabilityRequest{
			CurrentLocation: pkgLocation.Coordinate{
				Longitude: 11,
				Latitude:  11,
			},
		}
	)

	ctx := context.Background()
	ctx = pkgContext.SetDriverIDToContext(ctx, driverID)

	t.Run("success - should add available driver", func(t *testing.T) {
		repoRidesMock.EXPECT().GetDriverMSISDNByID(ctx, driverID).Return(driverMSISDN, nil)

		req.IsAvailable = true
		repoLocMock.EXPECT().AddAvailableDriver(ctx, driverID, req.CurrentLocation).
			Return(nil)
		repoRidesMock.EXPECT().UpdateDriverStatus(ctx, model.UpdateDriverStatusRequest{
			DriverID: driverID, Status: model.StatusDriverAvailable,
		}).Return(nil)

		err := usecaseMock.DriverSetAvailability(ctx, req)
		assert.Nil(t, err)
	})

	t.Run("success - should remove available driver", func(t *testing.T) {
		repoRidesMock.EXPECT().GetDriverMSISDNByID(ctx, driverID).Return(driverMSISDN, nil)

		req.IsAvailable = false
		repoLocMock.EXPECT().RemoveAvailableDriver(ctx, driverID).
			Return(nil)
		repoRidesMock.EXPECT().UpdateDriverStatus(ctx, model.UpdateDriverStatusRequest{
			DriverID: driverID, Status: model.StatusDriverOff,
		}).Return(nil)

		err := usecaseMock.DriverSetAvailability(ctx, req)
		assert.Nil(t, err)
	})

	t.Run("failed - driver msisdn is not found", func(t *testing.T) {
		repoRidesMock.EXPECT().GetDriverMSISDNByID(ctx, driverID).Return("", constants.ErrorDataNotFound)

		err := usecaseMock.DriverSetAvailability(ctx, req)
		assert.Equal(t, pkgError.ErrUnauthorizedCode, err.GetCode())
	})

	t.Run("failed - should return error when add available driver failed", func(t *testing.T) {
		repoRidesMock.EXPECT().GetDriverMSISDNByID(ctx, driverID).Return(driverMSISDN, nil)

		req.IsAvailable = true
		expectedErr := errors.New("error from repo")

		repoLocMock.EXPECT().AddAvailableDriver(ctx, driverID, req.CurrentLocation).
			Return(expectedErr)

		err := usecaseMock.DriverSetAvailability(ctx, req)
		assert.Equal(t, err.GetCode(), pkgError.ErrInternalErrorCode)
	})

	t.Run("failed - should return error when remove available driver failed", func(t *testing.T) {
		repoRidesMock.EXPECT().GetDriverMSISDNByID(ctx, driverID).Return(driverMSISDN, nil)

		req.IsAvailable = false
		expectedErr := errors.New("error from repo")

		repoLocMock.EXPECT().RemoveAvailableDriver(ctx, driverID).
			Return(expectedErr)

		err := usecaseMock.DriverSetAvailability(ctx, req)
		assert.Equal(t, err.GetCode(), pkgError.ErrInternalErrorCode)
	})

	t.Run("failed - should return error when update driver status failed", func(t *testing.T) {
		repoRidesMock.EXPECT().GetDriverMSISDNByID(ctx, driverID).Return(driverMSISDN, nil)

		req.IsAvailable = true
		expectedErr := errors.New("error from repo")

		repoLocMock.EXPECT().AddAvailableDriver(ctx, driverID, req.CurrentLocation).
			Return(nil)
		repoRidesMock.EXPECT().UpdateDriverStatus(ctx, model.UpdateDriverStatusRequest{
			DriverID: driverID, Status: model.StatusDriverAvailable,
		}).Return(expectedErr)

		err := usecaseMock.DriverSetAvailability(ctx, req)
		assert.Equal(t, err.GetCode(), pkgError.ErrInternalErrorCode)
	})
}
