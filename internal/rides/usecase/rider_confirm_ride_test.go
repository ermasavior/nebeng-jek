package usecase

import (
	"context"
	"errors"
	"testing"

	"nebeng-jek/internal/pkg/constants"
	pkgContext "nebeng-jek/internal/pkg/context"
	"nebeng-jek/internal/rides/model"
	mockRepo "nebeng-jek/mock/repository"
	pkgError "nebeng-jek/pkg/error"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUsecase_RiderConfirmRide(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	locationRepoMock := mockRepo.NewMockRidesLocationRepository(ctrl)
	ridesRepoMock := mockRepo.NewMockRidesRepository(ctrl)
	ridesPubsubMock := mockRepo.NewMockRidesPubsubRepository(ctrl)
	usecaseMock := NewUsecase(locationRepoMock, ridesRepoMock, ridesPubsubMock, nil)

	var (
		riderID  = int64(9999)
		driverID = int64(1111)
		rideData = model.RideData{
			RideID:    111,
			RiderID:   riderID,
			DriverID:  &driverID,
			StatusNum: model.StatusNumRideMatchedDriver,
			PickupLocation: model.Coordinate{
				Latitude:  1,
				Longitude: 2,
			},
			Destination: model.Coordinate{
				Latitude:  1,
				Longitude: 2,
			},
		}
		req = model.RiderConfirmRideRequest{
			RideID:   111,
			IsAccept: true,
		}
	)

	ctx := context.Background()
	ctx = pkgContext.SetRiderIDToContext(ctx, riderID)

	t.Run("success - should confirm ride rider and broadcast message", func(t *testing.T) {
		ridesRepoMock.EXPECT().GetRideData(ctx, req.RideID).Return(rideData, nil)
		ridesRepoMock.EXPECT().UpdateRideData(ctx, model.UpdateRideDataRequest{
			RideID: req.RideID,
			Status: model.StatusNumRideReadyToPickup,
		}).Return(nil)

		ridesPubsubMock.EXPECT().BroadcastMessage(ctx, constants.TopicRideReadyToPickup, model.RideReadyToPickupMessage{
			RideID:   req.RideID,
			RiderID:  riderID,
			DriverID: driverID,
		}).Return(nil)

		data, err := usecaseMock.RiderConfirmRide(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, model.StatusNumRideReadyToPickup, data.StatusNum)
	})

	t.Run("success - rider not accepting ride request", func(t *testing.T) {
		ridesRepoMock.EXPECT().GetRideData(ctx, req.RideID).Return(rideData, nil)
		ridesRepoMock.EXPECT().UpdateRideData(ctx, model.UpdateRideDataRequest{
			RideID: req.RideID,
			Status: model.StatusNumRideCancelled,
		}).Return(nil)

		data, err := usecaseMock.RiderConfirmRide(ctx, model.RiderConfirmRideRequest{
			RideID:   req.RideID,
			IsAccept: false,
		})
		assert.Nil(t, err)
		assert.Equal(t, model.StatusNumRideCancelled, data.StatusNum)
	})

	t.Run("failed - invalid ride data", func(t *testing.T) {
		invalidRide := rideData
		invalidRide.StatusNum = model.StatusNumRideCancelled
		ridesRepoMock.EXPECT().GetRideData(ctx, req.RideID).Return(invalidRide, nil)

		_, err := usecaseMock.RiderConfirmRide(ctx, req)
		assert.Equal(t, pkgError.ErrForbiddenCode, err.GetCode())
	})

	t.Run("failed - ride data is not found", func(t *testing.T) {
		ridesRepoMock.EXPECT().GetRideData(ctx, req.RideID).Return(model.RideData{}, constants.ErrorDataNotFound)

		_, err := usecaseMock.RiderConfirmRide(ctx, req)
		assert.Equal(t, pkgError.ErrResourceNotFoundCode, err.GetCode())
	})

	t.Run("failed - get ride data returns error", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetRideData(ctx, req.RideID).Return(model.RideData{}, expectedErr)

		data, err := usecaseMock.RiderConfirmRide(ctx, req)
		assert.Equal(t, err.GetCode(), pkgError.ErrInternalErrorCode)
		assert.Equal(t, model.RideData{}, data)
	})

	t.Run("failed - update ride returns error", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetRideData(ctx, req.RideID).Return(rideData, nil)
		ridesRepoMock.EXPECT().UpdateRideData(ctx, model.UpdateRideDataRequest{
			RideID: req.RideID,
			Status: model.StatusNumRideReadyToPickup,
		}).Return(expectedErr)

		data, err := usecaseMock.RiderConfirmRide(ctx, req)
		assert.Equal(t, err.GetCode(), pkgError.ErrInternalErrorCode)
		assert.Equal(t, model.RideData{}, data)
	})

	t.Run("failed - broadcast message returns error", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetRideData(ctx, req.RideID).Return(rideData, nil)
		ridesRepoMock.EXPECT().UpdateRideData(ctx, model.UpdateRideDataRequest{
			RideID: req.RideID,
			Status: model.StatusNumRideReadyToPickup,
		}).Return(nil)

		ridesPubsubMock.EXPECT().BroadcastMessage(ctx, constants.TopicRideReadyToPickup, model.RideReadyToPickupMessage{
			RideID:   rideData.RideID,
			RiderID:  riderID,
			DriverID: driverID,
		}).Return(expectedErr)

		data, err := usecaseMock.RiderConfirmRide(ctx, req)
		assert.Equal(t, err.GetCode(), pkgError.ErrInternalErrorCode)
		assert.Equal(t, model.RideData{}, data)
	})
}
