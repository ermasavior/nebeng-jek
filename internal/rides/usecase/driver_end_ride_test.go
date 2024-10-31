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

func TestUsecase_DriverEndRide(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	locationRepoMock := mockRepo.NewMockRidesLocationRepository(ctrl)
	ridesRepoMock := mockRepo.NewMockRidesRepository(ctrl)
	ridesPubsubMock := mockRepo.NewMockRidesPubsubRepository(ctrl)
	usecaseMock := NewUsecase(locationRepoMock, ridesRepoMock, ridesPubsubMock, nil)

	var (
		driverID = int64(7777)
		rideID   = int64(111)
		rideData = model.RideData{
			RideID:    rideID,
			RiderID:   6666,
			DriverID:  &driverID,
			StatusNum: model.StatusNumRideStarted,
			PickupLocation: model.Coordinate{
				Latitude:  1,
				Longitude: 2,
			},
			Destination: model.Coordinate{
				Latitude:  1,
				Longitude: 2,
			},
		}
		ridePath = []model.Coordinate{
			{
				Longitude: 1,
				Latitude:  1,
			},
			{
				Longitude: 2,
				Latitude:  2,
			},
			{
				Longitude: 3,
				Latitude:  3,
			},
		}

		distance = calculateTotalDistance(ridePath)
		fare     = calculateRideFare(distance)

		expectedRes = model.RideData{
			RideID:    rideID,
			RiderID:   6666,
			DriverID:  &driverID,
			StatusNum: model.StatusNumRideEnded,
			PickupLocation: model.Coordinate{
				Latitude:  1,
				Longitude: 2,
			},
			Destination: model.Coordinate{
				Latitude:  1,
				Longitude: 2,
			},
			Distance: &distance,
			Fare:     &fare,
		}

		req = model.DriverEndRideRequest{
			RideID: rideID,
		}
	)

	ctx := context.Background()
	ctx = pkgContext.SetDriverIDToContext(ctx, driverID)

	t.Run("success - should create new ride and publish message broadcast", func(t *testing.T) {
		locationRepoMock.EXPECT().GetRidePath(ctx, rideID, driverID).Return(ridePath, nil)
		ridesRepoMock.EXPECT().GetRideData(ctx, rideID).Return(rideData, nil)
		ridesRepoMock.EXPECT().UpdateRideData(ctx, model.UpdateRideDataRequest{
			RideID:   req.RideID,
			Status:   model.StatusNumRideEnded,
			Fare:     &fare,
			Distance: &distance,
		}).Return(nil)

		ridesPubsubMock.EXPECT().BroadcastMessage(ctx, constants.TopicRideEnded, model.RideEndedMessage{
			RideID:   rideID,
			Distance: distance,
			Fare:     fare,
			RiderID:  rideData.RiderID,
		}).Return(nil)

		actual, err := usecaseMock.DriverEndRide(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, expectedRes, actual)
	})

	t.Run("failed - should return error when get ride path failed", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		locationRepoMock.EXPECT().GetRidePath(ctx, rideID, driverID).Return([]model.Coordinate{}, expectedErr)

		_, err := usecaseMock.DriverEndRide(ctx, req)
		assert.Equal(t, err.GetCode(), pkgError.ErrInternalErrorCode)
	})

	t.Run("failed - should return error when get ride data is failed", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		locationRepoMock.EXPECT().GetRidePath(ctx, rideID, driverID).Return(ridePath, nil)
		ridesRepoMock.EXPECT().GetRideData(ctx, rideID).Return(model.RideData{}, expectedErr)

		_, err := usecaseMock.DriverEndRide(ctx, req)
		assert.Equal(t, err.GetCode(), pkgError.ErrInternalErrorCode)
	})

	t.Run("failed - should return error when update ride data is failed", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		locationRepoMock.EXPECT().GetRidePath(ctx, rideID, driverID).Return(ridePath, nil)
		ridesRepoMock.EXPECT().GetRideData(ctx, rideID).Return(rideData, nil)
		ridesRepoMock.EXPECT().UpdateRideData(ctx, model.UpdateRideDataRequest{
			RideID:   req.RideID,
			Status:   model.StatusNumRideEnded,
			Fare:     &fare,
			Distance: &distance,
		}).Return(expectedErr)

		_, err := usecaseMock.DriverEndRide(ctx, req)
		assert.Equal(t, err.GetCode(), pkgError.ErrInternalErrorCode)
	})

	t.Run("failed - should return error when broadcast data is failed", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		locationRepoMock.EXPECT().GetRidePath(ctx, rideID, driverID).Return(ridePath, nil)
		ridesRepoMock.EXPECT().GetRideData(ctx, rideID).Return(rideData, nil)
		ridesRepoMock.EXPECT().UpdateRideData(ctx, model.UpdateRideDataRequest{
			RideID:   req.RideID,
			Status:   model.StatusNumRideEnded,
			Fare:     &fare,
			Distance: &distance,
		}).Return(nil)

		ridesPubsubMock.EXPECT().BroadcastMessage(ctx, constants.TopicRideEnded, model.RideEndedMessage{
			RideID:   rideID,
			Distance: distance,
			Fare:     fare,
			RiderID:  rideData.RiderID,
		}).Return(expectedErr)

		_, err := usecaseMock.DriverEndRide(ctx, req)
		assert.Equal(t, err.GetCode(), pkgError.ErrInternalErrorCode)
	})
}