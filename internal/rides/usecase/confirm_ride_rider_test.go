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

func TestUsecase_ConfirmRideRider(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	locationRepoMock := mockRepo.NewMockRidesLocationRepository(ctrl)
	ridesRepoMock := mockRepo.NewMockRidesRepository(ctrl)
	ridesPubsubMock := mockRepo.NewMockRidesPubsubRepository(ctrl)
	usecaseMock := NewUsecase(locationRepoMock, ridesRepoMock, ridesPubsubMock, nil)

	var (
		riderMSISDN  = "0811111"
		driverMSISDN = "0821111"
		riderData    = model.RiderData{
			ID:     1111,
			Name:   "Agus",
			MSISDN: "0811111",
		}
		rideData = model.RideData{
			RideID:   111,
			RiderID:  666,
			DriverID: 1111,
			PickupLocation: model.Coordinate{
				Latitude:  1,
				Longitude: 2,
			},
			Destination: model.Coordinate{
				Latitude:  1,
				Longitude: 2,
			},
		}
		req = model.ConfirmRideRiderRequest{
			RideID:   111,
			IsAccept: true,
		}
	)

	ctx := context.Background()
	ctx = pkgContext.SetMSISDNToContext(ctx, riderMSISDN)

	t.Run("success - should confirm ride rider and broadcast message", func(t *testing.T) {
		ridesRepoMock.EXPECT().GetRiderDataByMSISDN(ctx, riderMSISDN).Return(riderData, nil)
		ridesRepoMock.EXPECT().ConfirmRideRider(ctx, model.ConfirmRideRiderRequest{
			RiderID:  riderData.ID,
			RideID:   req.RideID,
			IsAccept: true,
		}).Return(rideData, nil)

		ridesRepoMock.EXPECT().GetDriverMSISDNByID(ctx, rideData.DriverID).Return(driverMSISDN, nil)

		ridesPubsubMock.EXPECT().BroadcastMessage(ctx, constants.TopicRideReadyToPickup, model.RideReadyToPickupMessage{
			RideID:         rideData.RideID,
			PickupLocation: rideData.PickupLocation,
			Destination:    rideData.Destination,
			RiderMSISDN:    riderMSISDN,
			DriverMSISDN:   driverMSISDN,
		}).Return(nil)

		err := usecaseMock.ConfirmRideRider(ctx, req)
		assert.Nil(t, err)
	})

	t.Run("success - driver not accepting ride request", func(t *testing.T) {
		err := usecaseMock.ConfirmRideRider(ctx, model.ConfirmRideRiderRequest{
			RiderID:  req.RiderID,
			RideID:   req.RideID,
			IsAccept: false,
		})
		assert.Nil(t, err)
	})

	t.Run("failed - get rider data returns error", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetRiderDataByMSISDN(ctx, riderMSISDN).Return(model.RiderData{}, expectedErr)

		err := usecaseMock.ConfirmRideRider(ctx, req)
		assert.Error(t, err, pkgError.NewInternalServerError(expectedErr, "error get rider data"))
	})

	t.Run("failed - confirm ride returns error", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetRiderDataByMSISDN(ctx, riderMSISDN).Return(riderData, nil)
		ridesRepoMock.EXPECT().ConfirmRideRider(ctx, model.ConfirmRideRiderRequest{
			RiderID:  riderData.ID,
			RideID:   req.RideID,
			IsAccept: true,
		}).Return(model.RideData{}, expectedErr)

		err := usecaseMock.ConfirmRideRider(ctx, req)
		assert.Error(t, err, pkgError.NewInternalServerError(expectedErr, "error confirm ride by driver"))
	})

	t.Run("failed - get rider msisdn by id returns error", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetRiderDataByMSISDN(ctx, riderMSISDN).Return(riderData, nil)
		ridesRepoMock.EXPECT().ConfirmRideRider(ctx, model.ConfirmRideRiderRequest{
			RiderID:  riderData.ID,
			RideID:   req.RideID,
			IsAccept: true,
		}).Return(rideData, nil)

		ridesRepoMock.EXPECT().GetDriverMSISDNByID(ctx, rideData.DriverID).Return("", expectedErr)

		err := usecaseMock.ConfirmRideRider(ctx, req)
		assert.Error(t, err, pkgError.NewInternalServerError(expectedErr, "error get driver msisdn"))
	})

	t.Run("failed - broadcast message returns error", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetRiderDataByMSISDN(ctx, riderMSISDN).Return(riderData, nil)
		ridesRepoMock.EXPECT().ConfirmRideRider(ctx, model.ConfirmRideRiderRequest{
			RiderID:  riderData.ID,
			RideID:   req.RideID,
			IsAccept: true,
		}).Return(rideData, nil)

		ridesRepoMock.EXPECT().GetDriverMSISDNByID(ctx, rideData.DriverID).Return(driverMSISDN, nil)

		ridesPubsubMock.EXPECT().BroadcastMessage(ctx, constants.TopicRideReadyToPickup, model.RideReadyToPickupMessage{
			RideID:         rideData.RideID,
			PickupLocation: rideData.PickupLocation,
			Destination:    rideData.Destination,
			RiderMSISDN:    riderMSISDN,
			DriverMSISDN:   driverMSISDN,
		}).Return(expectedErr)

		err := usecaseMock.ConfirmRideRider(ctx, req)
		assert.Error(t, err, pkgError.NewInternalServerError(expectedErr, "error broadcasting ride ready to pickup"))
	})
}
