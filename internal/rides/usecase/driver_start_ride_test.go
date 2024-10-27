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

func TestUsecase_DriverStartRide(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	locationRepoMock := mockRepo.NewMockRidesLocationRepository(ctrl)
	ridesRepoMock := mockRepo.NewMockRidesRepository(ctrl)
	ridesPubsubMock := mockRepo.NewMockRidesPubsubRepository(ctrl)
	usecaseMock := NewUsecase(locationRepoMock, ridesRepoMock, ridesPubsubMock, nil)

	var (
		driverID = int64(1111)
		rideData = model.RideData{
			RideID:    111,
			RiderID:   666,
			DriverID:  &driverID,
			StatusNum: model.StatusNumRideReadyToPickup,
			PickupLocation: model.Coordinate{
				Latitude:  1,
				Longitude: 2,
			},
			Destination: model.Coordinate{
				Latitude:  1,
				Longitude: 2,
			},
		}
		req = model.DriverStartRideRequest{
			RideID:   111,
			DriverID: 1111,
		}
	)

	ctx := context.Background()
	ctx = pkgContext.SetDriverIDToContext(ctx, driverID)

	t.Run("success - should confirm ride driver and broadcast to rider", func(t *testing.T) {
		ridesRepoMock.EXPECT().GetRideData(ctx, req.RideID).Return(rideData, nil)
		ridesRepoMock.EXPECT().UpdateRideData(ctx, model.UpdateRideDataRequest{
			RideID: req.RideID,
			Status: model.StatusNumRideStarted,
		}).Return(nil)

		locationRepoMock.EXPECT().RemoveAvailableDriver(ctx, driverID).Return(nil)

		ridesPubsubMock.EXPECT().BroadcastMessage(ctx, constants.TopicRideStarted, model.RideStartedMessage{
			RideID:  rideData.RideID,
			RiderID: rideData.RiderID,
		}).Return(nil)

		res, err := usecaseMock.DriverStartRide(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, model.StatusNumRideStarted, res.StatusNum)
	})

	t.Run("failed - get ride data returns error", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetRideData(ctx, req.RideID).Return(model.RideData{}, expectedErr)

		res, err := usecaseMock.DriverStartRide(ctx, req)
		assert.Error(t, err, pkgError.NewInternalServerError("error get ride data"))
		assert.Equal(t, model.RideData{}, res)
	})

	t.Run("failed - update ride returns error", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetRideData(ctx, req.RideID).Return(rideData, nil)
		ridesRepoMock.EXPECT().UpdateRideData(ctx, model.UpdateRideDataRequest{
			RideID: req.RideID,
			Status: model.StatusNumRideStarted,
		}).Return(expectedErr)

		res, err := usecaseMock.DriverStartRide(ctx, req)
		assert.Equal(t, err.GetCode(), pkgError.ErrInternalErrorCode)
		assert.Equal(t, model.RideData{}, res)
	})

	t.Run("failed - remove available driver returns error", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetRideData(ctx, req.RideID).Return(rideData, nil)
		ridesRepoMock.EXPECT().UpdateRideData(ctx, model.UpdateRideDataRequest{
			RideID: req.RideID,
			Status: model.StatusNumRideStarted,
		}).Return(nil)
		locationRepoMock.EXPECT().RemoveAvailableDriver(ctx, driverID).Return(expectedErr)

		res, err := usecaseMock.DriverStartRide(ctx, req)
		assert.Equal(t, err.GetCode(), pkgError.ErrInternalErrorCode)
		assert.Equal(t, model.RideData{}, res)
	})

	t.Run("failed - broadcast message returns error", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetRideData(ctx, req.RideID).Return(rideData, nil)
		ridesRepoMock.EXPECT().UpdateRideData(ctx, model.UpdateRideDataRequest{
			RideID: req.RideID,
			Status: model.StatusNumRideStarted,
		}).Return(nil)
		locationRepoMock.EXPECT().RemoveAvailableDriver(ctx, driverID).Return(nil)

		ridesPubsubMock.EXPECT().BroadcastMessage(ctx, constants.TopicRideStarted, model.RideStartedMessage{
			RideID:  rideData.RideID,
			RiderID: rideData.RiderID,
		}).Return(expectedErr)

		res, err := usecaseMock.DriverStartRide(ctx, req)
		assert.Equal(t, err.GetCode(), pkgError.ErrInternalErrorCode)
		assert.Equal(t, model.RideData{}, res)
	})
}
