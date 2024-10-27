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

func TestUsecase_DriverConfirmRide(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	locationRepoMock := mockRepo.NewMockRidesLocationRepository(ctrl)
	ridesRepoMock := mockRepo.NewMockRidesRepository(ctrl)
	ridesPubsubMock := mockRepo.NewMockRidesPubsubRepository(ctrl)
	usecaseMock := NewUsecase(locationRepoMock, ridesRepoMock, ridesPubsubMock, nil)

	var (
		driverID   = int64(1111)
		driverData = model.DriverData{
			ID:           1111,
			Name:         "Agus",
			MSISDN:       "0811111",
			VehicleType:  "CAR",
			VehiclePlate: "B11111A",
		}
		rideData = model.RideData{
			RideID:   111,
			RiderID:  666,
			DriverID: 1111,
			Status:   model.StatusRideWaitingForDriver,
			PickupLocation: model.Coordinate{
				Latitude:  1,
				Longitude: 2,
			},
			Destination: model.Coordinate{
				Latitude:  1,
				Longitude: 2,
			},
		}
		req = model.DriverConfirmRideRequest{
			RideID:   111,
			IsAccept: true,
		}
	)

	ctx := context.Background()
	ctx = pkgContext.SetDriverIDToContext(ctx, driverID)

	t.Run("success - should confirm ride driver and broadcast to rider", func(t *testing.T) {
		ridesRepoMock.EXPECT().GetDriverDataByID(ctx, driverID).Return(driverData, nil)
		ridesRepoMock.EXPECT().GetRideData(ctx, req.RideID).Return(rideData, nil)
		ridesRepoMock.EXPECT().UpdateRideData(ctx, model.UpdateRideDataRequest{
			RideID:   req.RideID,
			DriverID: driverID,
			Status:   model.StatusNumRideDriverMatched,
		}).Return(nil)

		ridesPubsubMock.EXPECT().BroadcastMessage(ctx, constants.TopicRideMatchedDriver, model.MatchedRideMessage{
			RideID:  rideData.RideID,
			Driver:  driverData,
			RiderID: rideData.RiderID,
		}).Return(nil)

		err := usecaseMock.DriverConfirmRide(ctx, req)
		assert.Nil(t, err)
	})

	t.Run("success - driver not accepting ride request", func(t *testing.T) {
		err := usecaseMock.DriverConfirmRide(ctx, model.DriverConfirmRideRequest{
			RideID:   req.RideID,
			IsAccept: false,
		})
		assert.Nil(t, err)
	})

	t.Run("failed - get driver data returns error", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetDriverDataByID(ctx, driverID).Return(model.DriverData{}, expectedErr)

		err := usecaseMock.DriverConfirmRide(ctx, req)
		assert.Equal(t, err.GetCode(), pkgError.ErrInternalErrorCode)
	})

	t.Run("failed - get ride data returns error", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetDriverDataByID(ctx, driverID).Return(driverData, nil)
		ridesRepoMock.EXPECT().GetRideData(ctx, req.RideID).Return(model.RideData{}, expectedErr)

		err := usecaseMock.DriverConfirmRide(ctx, req)
		assert.Equal(t, err.GetCode(), pkgError.ErrInternalErrorCode)
	})

	t.Run("failed - update ride data returns error", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetDriverDataByID(ctx, driverID).Return(driverData, nil)
		ridesRepoMock.EXPECT().GetRideData(ctx, req.RideID).Return(rideData, nil)
		ridesRepoMock.EXPECT().UpdateRideData(ctx, model.UpdateRideDataRequest{
			RideID:   req.RideID,
			DriverID: driverID,
			Status:   model.StatusNumRideDriverMatched,
		}).Return(expectedErr)

		err := usecaseMock.DriverConfirmRide(ctx, req)
		assert.Equal(t, err.GetCode(), pkgError.ErrInternalErrorCode)
	})

	t.Run("failed - broadcast message returns error", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetDriverDataByID(ctx, driverID).Return(driverData, nil)
		ridesRepoMock.EXPECT().GetRideData(ctx, req.RideID).Return(rideData, nil)
		ridesRepoMock.EXPECT().UpdateRideData(ctx, model.UpdateRideDataRequest{
			RideID:   req.RideID,
			DriverID: driverID,
			Status:   model.StatusNumRideDriverMatched,
		}).Return(nil)

		ridesPubsubMock.EXPECT().BroadcastMessage(ctx, constants.TopicRideMatchedDriver, model.MatchedRideMessage{
			RideID:  rideData.RideID,
			Driver:  driverData,
			RiderID: rideData.RiderID,
		}).Return(expectedErr)

		err := usecaseMock.DriverConfirmRide(ctx, req)
		assert.Equal(t, err.GetCode(), pkgError.ErrInternalErrorCode)
	})
}
