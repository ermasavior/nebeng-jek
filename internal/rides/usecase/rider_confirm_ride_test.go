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
			RideID:   111,
			RiderID:  riderID,
			DriverID: driverID,
			Status:   model.StatusRideDriverMatched,
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
			Status: model.StatusNumRideWaitingForPickup,
		}).Return(nil)

		ridesPubsubMock.EXPECT().BroadcastMessage(ctx, constants.TopicRideReadyToPickup, model.RideReadyToPickupMessage{
			RideID:   req.RideID,
			RiderID:  riderID,
			DriverID: driverID,
		}).Return(nil)

		err := usecaseMock.RiderConfirmRide(ctx, req)
		assert.Nil(t, err)
	})

	t.Run("success - rider not accepting ride request", func(t *testing.T) {
		ridesRepoMock.EXPECT().GetRideData(ctx, req.RideID).Return(rideData, nil)
		ridesRepoMock.EXPECT().UpdateRideData(ctx, model.UpdateRideDataRequest{
			RideID: req.RideID,
			Status: model.StatusNumRideCancelled,
		}).Return(nil)

		err := usecaseMock.RiderConfirmRide(ctx, model.RiderConfirmRideRequest{
			RiderID:  req.RiderID,
			RideID:   req.RideID,
			IsAccept: false,
		})
		assert.Nil(t, err)
	})

	t.Run("failed - get ride data returns error", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetRideData(ctx, req.RideID).Return(model.RideData{}, expectedErr)

		err := usecaseMock.RiderConfirmRide(ctx, req)
		assert.Equal(t, err.GetCode(), pkgError.ErrInternalErrorCode)
	})

	t.Run("failed - update ride returns error", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetRideData(ctx, req.RideID).Return(rideData, nil)
		ridesRepoMock.EXPECT().UpdateRideData(ctx, model.UpdateRideDataRequest{
			RideID: req.RideID,
			Status: model.StatusNumRideWaitingForPickup,
		}).Return(expectedErr)

		err := usecaseMock.RiderConfirmRide(ctx, req)
		assert.Equal(t, err.GetCode(), pkgError.ErrInternalErrorCode)
	})

	t.Run("failed - broadcast message returns error", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetRideData(ctx, req.RideID).Return(rideData, nil)
		ridesRepoMock.EXPECT().UpdateRideData(ctx, model.UpdateRideDataRequest{
			RideID: req.RideID,
			Status: model.StatusNumRideWaitingForPickup,
		}).Return(nil)

		ridesPubsubMock.EXPECT().BroadcastMessage(ctx, constants.TopicRideReadyToPickup, model.RideReadyToPickupMessage{
			RideID:   rideData.RideID,
			RiderID:  riderID,
			DriverID: driverID,
		}).Return(expectedErr)

		err := usecaseMock.RiderConfirmRide(ctx, req)
		assert.Equal(t, err.GetCode(), pkgError.ErrInternalErrorCode)
	})
}
