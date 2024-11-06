package usecase

import (
	"context"
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

func TestUsecase_GetRideData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ridesRepoMock := mockRepo.NewMockRidesRepository(ctrl)
	usecaseMock := NewUsecase(nil, ridesRepoMock, nil, nil)

	var (
		rideID   = int64(666)
		riderID  = int64(9999)
		driverID = int64(1111)
		rideData = model.RideData{
			RideID:    rideID,
			RiderID:   riderID,
			DriverID:  &driverID,
			StatusNum: model.StatusNumRideMatchedDriver,
			PickupLocation: pkgLocation.Coordinate{
				Latitude:  1,
				Longitude: 2,
			},
			Destination: pkgLocation.Coordinate{
				Latitude:  1,
				Longitude: 2,
			},
		}
	)

	ctx := context.Background()
	ctx = pkgContext.SetRiderIDToContext(ctx, riderID)

	t.Run("success - should return ride data", func(t *testing.T) {
		ridesRepoMock.EXPECT().GetRideData(ctx, rideID).Return(rideData, nil)
		data, err := usecaseMock.GetRideData(ctx, rideID)
		assert.Nil(t, err)
		assert.Equal(t, rideData, data)
	})

	t.Run("error - should return not found", func(t *testing.T) {
		ridesRepoMock.EXPECT().GetRideData(ctx, rideID).Return(model.RideData{}, constants.ErrorDataNotFound)
		data, err := usecaseMock.GetRideData(ctx, rideID)
		assert.Equal(t, pkgError.ErrResourceNotFoundCode, err.GetCode())
		assert.Equal(t, model.RideData{}, data)
	})

	t.Run("error - should return unauthorized", func(t *testing.T) {
		ctx = pkgContext.SetRiderIDToContext(ctx, 0)
		ridesRepoMock.EXPECT().GetRideData(ctx, rideID).Return(rideData, nil)
		data, err := usecaseMock.GetRideData(ctx, rideID)
		assert.Equal(t, pkgError.ErrUnauthorizedCode, err.GetCode())
		assert.Equal(t, model.RideData{}, data)
	})
}
