package usecase

import (
	"context"
	"errors"
	"math"
	"testing"

	"nebeng-jek/internal/pkg/constants"
	pkgContext "nebeng-jek/internal/pkg/context"
	pkgLocation "nebeng-jek/internal/pkg/location"
	"nebeng-jek/internal/rides/model"
	mockRepo "nebeng-jek/mock/repository"
	"nebeng-jek/pkg/configs"
	pkgError "nebeng-jek/pkg/error"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUsecase_DriverEndRide(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg := &configs.Config{
		RidePricePerKm:    3000,
		RideFeePercentage: 30,
	}
	locationRepoMock := mockRepo.NewMockRidesLocationRepository(ctrl)
	ridesRepoMock := mockRepo.NewMockRidesRepository(ctrl)
	ridesPubsubMock := mockRepo.NewMockRidesPubsubRepository(ctrl)
	usecaseMock := NewUsecase(cfg, locationRepoMock, ridesRepoMock, ridesPubsubMock, nil)

	var (
		driverID = int64(7777)
		rideID   = int64(111)
		rideData = model.RideData{
			RideID:    rideID,
			RiderID:   6666,
			DriverID:  &driverID,
			StatusNum: model.StatusNumRideStarted,
			PickupLocation: pkgLocation.Coordinate{
				Latitude:  1,
				Longitude: 2,
			},
			Destination: pkgLocation.Coordinate{
				Latitude:  1,
				Longitude: 2,
			},
		}
		path = []pkgLocation.Coordinate{
			{Longitude: 1, Latitude: 1},
			{Longitude: 2, Latitude: 2},
			{Longitude: 3, Latitude: 3},
		}
		ridePathRes = model.GetRidePathResponse{
			DriverPath: path,
			RiderPath:  path,
		}

		distance = calculateTotalDistance(path)
		fare     = math.Ceil(distance) * cfg.RidePricePerKm

		expectedRes = model.RideData{
			RideID:    rideID,
			RiderID:   6666,
			DriverID:  &driverID,
			StatusNum: model.StatusNumRideEnded,
			PickupLocation: pkgLocation.Coordinate{
				Latitude:  1,
				Longitude: 2,
			},
			Destination: pkgLocation.Coordinate{
				Latitude:  1,
				Longitude: 2,
			},
			Distance: &distance,
			Fare:     &fare,
		}

		ridePathReq = model.GetRidePathRequest{
			RideID: rideID, DriverID: driverID, RiderID: rideData.RiderID,
		}
		req = model.DriverEndRideRequest{
			RideID: rideID,
		}
	)

	ctx := context.Background()
	ctx = pkgContext.SetDriverIDToContext(ctx, driverID)

	t.Run("success - should create new ride and publish message broadcast", func(t *testing.T) {
		ridesRepoMock.EXPECT().GetRideData(ctx, rideID).Return(rideData, nil)
		locationRepoMock.EXPECT().GetRidePath(ctx, model.GetRidePathRequest{
			RideID: rideID, RiderID: rideData.RiderID, DriverID: driverID,
		}).Return(ridePathRes, nil)
		ridesRepoMock.EXPECT().UpdateRideData(ctx, gomock.AssignableToTypeOf(model.UpdateRideDataRequest{})).
			Return(nil)

		ridesPubsubMock.EXPECT().BroadcastMessage(ctx, constants.TopicRideEnded, model.RideEndedMessage{
			RideID:   rideID,
			Distance: distance,
			Fare:     fare,
			RiderID:  rideData.RiderID,
		}).Return(nil).AnyTimes()

		actual, err := usecaseMock.DriverEndRide(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, expectedRes, actual)
	})

	t.Run("failed - ride data is not found", func(t *testing.T) {
		ridesRepoMock.EXPECT().GetRideData(ctx, req.RideID).Return(model.RideData{}, constants.ErrorDataNotFound)

		_, err := usecaseMock.DriverEndRide(ctx, req)
		assert.Equal(t, pkgError.ErrResourceNotFoundCode, err.GetCode())
	})

	t.Run("failed - invalid ride data", func(t *testing.T) {
		invalidRide := rideData
		invalidRide.StatusNum = model.StatusNumRideCancelled
		ridesRepoMock.EXPECT().GetRideData(ctx, req.RideID).Return(invalidRide, nil)
		locationRepoMock.EXPECT().GetRidePath(ctx, model.GetRidePathRequest{
			RideID: rideID, RiderID: rideData.RiderID, DriverID: driverID,
		}).Return(ridePathRes, nil)

		_, err := usecaseMock.DriverEndRide(ctx, req)
		assert.Equal(t, pkgError.ErrForbiddenCode, err.GetCode())
	})

	t.Run("failed - should return error when get ride data is failed", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetRideData(ctx, rideID).Return(model.RideData{}, expectedErr)

		_, err := usecaseMock.DriverEndRide(ctx, req)
		assert.Equal(t, pkgError.ErrInternalErrorCode, err.GetCode())
	})

	t.Run("failed - should return error when get ride path failed", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetRideData(ctx, rideID).Return(rideData, nil)
		locationRepoMock.EXPECT().GetRidePath(ctx, ridePathReq).Return(model.GetRidePathResponse{}, expectedErr)

		_, err := usecaseMock.DriverEndRide(ctx, req)
		assert.Equal(t, pkgError.ErrInternalErrorCode, err.GetCode())
	})

	t.Run("failed - should return error when update ride data is failed", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetRideData(ctx, rideID).Return(rideData, nil)
		locationRepoMock.EXPECT().GetRidePath(ctx, ridePathReq).Return(ridePathRes, nil)
		ridesRepoMock.EXPECT().UpdateRideData(ctx, gomock.AssignableToTypeOf(model.UpdateRideDataRequest{})).
			Return(expectedErr)

		_, err := usecaseMock.DriverEndRide(ctx, req)
		assert.Equal(t, pkgError.ErrInternalErrorCode, err.GetCode())
	})

	t.Run("ignore - broadcast data is failed", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetRideData(ctx, rideID).Return(rideData, nil)
		locationRepoMock.EXPECT().GetRidePath(ctx, ridePathReq).Return(ridePathRes, nil)
		ridesRepoMock.EXPECT().UpdateRideData(ctx, gomock.AssignableToTypeOf(model.UpdateRideDataRequest{})).
			Return(nil)

		ridesPubsubMock.EXPECT().BroadcastMessage(ctx, constants.TopicRideEnded, model.RideEndedMessage{
			RideID:   rideID,
			Distance: distance,
			Fare:     fare,
			RiderID:  rideData.RiderID,
		}).Return(expectedErr).AnyTimes()

		_, err := usecaseMock.DriverEndRide(ctx, req)
		assert.Nil(t, err)
	})
}
