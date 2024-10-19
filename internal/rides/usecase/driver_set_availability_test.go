package usecase

import (
	"context"
	"errors"
	"testing"

	pkgContext "nebeng-jek/internal/pkg/context"
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
		driverID = int64(1111)

		req = model.DriverSetAvailabilityRequest{
			CurrentLocation: model.Coordinate{
				Longitude: 11,
				Latitude:  11,
			},
		}
	)

	ctx := context.Background()
	ctx = pkgContext.SetDriverIDToContext(ctx, driverID)

	t.Run("success - should add available driver", func(t *testing.T) {
		repoRidesMock.EXPECT().GetDriverDataByID(ctx, driverID).Return(model.DriverData{}, nil)

		req.IsAvailable = true
		repoLocMock.EXPECT().AddAvailableDriver(ctx, driverID, req.CurrentLocation).
			Return(nil)

		err := usecaseMock.DriverSetAvailability(ctx, req)
		assert.Nil(t, err)
	})

	t.Run("success - should remove available driver", func(t *testing.T) {
		repoRidesMock.EXPECT().GetDriverDataByID(ctx, driverID).Return(model.DriverData{}, nil)

		req.IsAvailable = false
		repoLocMock.EXPECT().RemoveAvailableDriver(ctx, driverID).
			Return(nil)

		err := usecaseMock.DriverSetAvailability(ctx, req)
		assert.Nil(t, err)
	})

	t.Run("failed - should return error when add available driver failed", func(t *testing.T) {
		repoRidesMock.EXPECT().GetDriverDataByID(ctx, driverID).Return(model.DriverData{}, nil)

		req.IsAvailable = true
		expectedErr := errors.New("error from repo")

		repoLocMock.EXPECT().AddAvailableDriver(ctx, driverID, req.CurrentLocation).
			Return(expectedErr)

		err := usecaseMock.DriverSetAvailability(ctx, req)
		assert.Equal(t, pkgError.NewInternalServerError(expectedErr, "error adding available driver"), err)
	})

	t.Run("failed - should return error when remove available driver failed", func(t *testing.T) {
		repoRidesMock.EXPECT().GetDriverDataByID(ctx, driverID).Return(model.DriverData{}, nil)

		req.IsAvailable = false
		expectedErr := errors.New("error from repo")

		repoLocMock.EXPECT().RemoveAvailableDriver(ctx, driverID).
			Return(expectedErr)

		err := usecaseMock.DriverSetAvailability(ctx, req)
		assert.Equal(t, pkgError.NewInternalServerError(expectedErr, "error removing available driver"), err)
	})
}
