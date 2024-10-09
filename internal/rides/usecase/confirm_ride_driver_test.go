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

func TestUsecase_ConfirmRideDriver(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	locationRepoMock := mockRepo.NewMockRidesLocationRepository(ctrl)
	ridesRepoMock := mockRepo.NewMockRidesRepository(ctrl)
	ridesPubsubMock := mockRepo.NewMockRidesPubsubRepository(ctrl)
	usecaseMock := NewUsecase(locationRepoMock, ridesRepoMock, ridesPubsubMock)

	var (
		driverMsisdn = "0811111"
		riderMSISDN  = "0821111"
		driverData   = model.DriverData{
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
			PickupLocation: model.Coordinate{
				Latitude:  1,
				Longitude: 2,
			},
			Destination: model.Coordinate{
				Latitude:  1,
				Longitude: 2,
			},
		}
		req = model.ConfirmRideDriverRequest{
			RideID:   111,
			IsAccept: true,
		}
	)

	ctx := context.Background()
	ctx = pkgContext.SetMSISDNToContext(ctx, driverMsisdn)

	t.Run("success - should confirm ride driver and broadcast to rider", func(t *testing.T) {
		ridesRepoMock.EXPECT().GetDriverDataByMSISDN(ctx, driverMsisdn).Return(driverData, nil)
		ridesRepoMock.EXPECT().ConfirmRideDriver(ctx, model.ConfirmRideDriverRequest{
			DriverID: driverData.ID,
			RideID:   req.RideID,
			IsAccept: true,
		}).Return(rideData, nil)

		ridesRepoMock.EXPECT().GetRiderMSISDNByID(ctx, rideData.RiderID).Return(riderMSISDN, nil)

		ridesPubsubMock.EXPECT().BroadcastMessage(ctx, constants.DriverAcceptedRideExchange, model.MatchedRideMessage{
			RideID:      rideData.RideID,
			Driver:      driverData,
			RiderMSISDN: riderMSISDN,
		}).Return(nil)

		err := usecaseMock.ConfirmRideDriver(ctx, req)
		assert.Nil(t, err)
	})

	t.Run("success - driver not accepting ride request", func(t *testing.T) {
		err := usecaseMock.ConfirmRideDriver(ctx, model.ConfirmRideDriverRequest{
			DriverID: req.DriverID,
			RideID:   req.RideID,
			IsAccept: false,
		})
		assert.Nil(t, err)
	})

	t.Run("failed - get driver data returns error", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetDriverDataByMSISDN(ctx, driverMsisdn).Return(model.DriverData{}, expectedErr)

		err := usecaseMock.ConfirmRideDriver(ctx, req)
		assert.Error(t, err, pkgError.NewInternalServerError(expectedErr, "error get driver data"))
	})

	t.Run("failed - confirm ride returns error", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetDriverDataByMSISDN(ctx, driverMsisdn).Return(driverData, nil)
		ridesRepoMock.EXPECT().ConfirmRideDriver(ctx, model.ConfirmRideDriverRequest{
			DriverID: driverData.ID,
			RideID:   req.RideID,
			IsAccept: true,
		}).Return(model.RideData{}, expectedErr)

		err := usecaseMock.ConfirmRideDriver(ctx, req)
		assert.Error(t, err, pkgError.NewInternalServerError(expectedErr, "error get driver data"))
	})

	t.Run("failed - get rider msisdn by id returns error", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetDriverDataByMSISDN(ctx, driverMsisdn).Return(driverData, nil)
		ridesRepoMock.EXPECT().ConfirmRideDriver(ctx, model.ConfirmRideDriverRequest{
			DriverID: driverData.ID,
			RideID:   req.RideID,
			IsAccept: true,
		}).Return(rideData, nil)

		ridesRepoMock.EXPECT().GetRiderMSISDNByID(ctx, rideData.RiderID).Return("", expectedErr)

		err := usecaseMock.ConfirmRideDriver(ctx, req)
		assert.Error(t, err, pkgError.NewInternalServerError(expectedErr, "error get rider msisdn"))
	})

	t.Run("failed - broadcast message returns error", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetDriverDataByMSISDN(ctx, driverMsisdn).Return(driverData, nil)
		ridesRepoMock.EXPECT().ConfirmRideDriver(ctx, model.ConfirmRideDriverRequest{
			DriverID: driverData.ID,
			RideID:   req.RideID,
			IsAccept: true,
		}).Return(rideData, nil)

		ridesRepoMock.EXPECT().GetRiderMSISDNByID(ctx, rideData.RiderID).Return(riderMSISDN, nil)

		ridesPubsubMock.EXPECT().BroadcastMessage(ctx, constants.DriverAcceptedRideExchange, model.MatchedRideMessage{
			RideID:      rideData.RideID,
			Driver:      driverData,
			RiderMSISDN: riderMSISDN,
		}).Return(expectedErr)

		err := usecaseMock.ConfirmRideDriver(ctx, req)
		assert.Error(t, err, pkgError.NewInternalServerError(expectedErr, "error broadcasting matched ride to rider"))
	})
}
