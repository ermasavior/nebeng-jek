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

func TestUsecase_StartRideDriver(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	locationRepoMock := mockRepo.NewMockRidesLocationRepository(ctrl)
	ridesRepoMock := mockRepo.NewMockRidesRepository(ctrl)
	ridesPubsubMock := mockRepo.NewMockRidesPubsubRepository(ctrl)
	usecaseMock := NewUsecase(locationRepoMock, ridesRepoMock, ridesPubsubMock, nil)

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
		req = model.StartRideDriverRequest{
			RideID:   111,
			DriverID: 1111,
		}
	)

	ctx := context.Background()
	ctx = pkgContext.SetMSISDNToContext(ctx, driverMsisdn)

	t.Run("success - should confirm ride driver and broadcast to rider", func(t *testing.T) {
		ridesRepoMock.EXPECT().GetDriverDataByMSISDN(ctx, driverMsisdn).Return(driverData, nil)
		ridesRepoMock.EXPECT().UpdateRideByDriver(ctx, model.UpdateRideByDriverRequest{
			DriverID: driverData.ID,
			RideID:   req.RideID,
			Status:   model.StatusNumRideStarted,
		}).Return(rideData, nil)

		ridesRepoMock.EXPECT().GetRiderMSISDNByID(ctx, rideData.RiderID).Return(riderMSISDN, nil)
		locationRepoMock.EXPECT().RemoveAvailableDriver(ctx, driverMsisdn).Return(nil)

		ridesPubsubMock.EXPECT().BroadcastMessage(ctx, constants.TopicRideStarted, model.RideStartedMessage{
			RideID:      rideData.RideID,
			RiderMSISDN: riderMSISDN,
		}).Return(nil)

		res, err := usecaseMock.StartRideDriver(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, rideData, res)
	})

	t.Run("failed - get driver data returns error", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetDriverDataByMSISDN(ctx, driverMsisdn).Return(model.DriverData{}, expectedErr)

		res, err := usecaseMock.StartRideDriver(ctx, req)
		assert.Error(t, err, pkgError.NewInternalServerError(expectedErr, "error get driver data"))
		assert.Equal(t, model.RideData{}, res)
	})

	t.Run("failed - confirm ride returns error", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetDriverDataByMSISDN(ctx, driverMsisdn).Return(driverData, nil)
		ridesRepoMock.EXPECT().UpdateRideByDriver(ctx, model.UpdateRideByDriverRequest{
			DriverID: driverData.ID,
			RideID:   req.RideID,
			Status:   model.StatusNumRideStarted,
		}).Return(model.RideData{}, expectedErr)

		res, err := usecaseMock.StartRideDriver(ctx, req)
		assert.Error(t, err, pkgError.NewInternalServerError(expectedErr, "error get driver data"))
		assert.Equal(t, model.RideData{}, res)
	})

	t.Run("failed - remove available driver returns error", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetDriverDataByMSISDN(ctx, driverMsisdn).Return(driverData, nil)
		ridesRepoMock.EXPECT().UpdateRideByDriver(ctx, model.UpdateRideByDriverRequest{
			DriverID: driverData.ID,
			RideID:   req.RideID,
			Status:   model.StatusNumRideStarted,
		}).Return(rideData, nil)
		locationRepoMock.EXPECT().RemoveAvailableDriver(ctx, driverMsisdn).Return(expectedErr)

		res, err := usecaseMock.StartRideDriver(ctx, req)
		assert.Error(t, err, pkgError.NewInternalServerError(expectedErr, "error removing available driver"))
		assert.Equal(t, model.RideData{}, res)
	})

	t.Run("failed - get rider msisdn by id returns error", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetDriverDataByMSISDN(ctx, driverMsisdn).Return(driverData, nil)
		ridesRepoMock.EXPECT().UpdateRideByDriver(ctx, model.UpdateRideByDriverRequest{
			DriverID: driverData.ID,
			RideID:   req.RideID,
			Status:   model.StatusNumRideStarted,
		}).Return(rideData, nil)
		locationRepoMock.EXPECT().RemoveAvailableDriver(ctx, driverMsisdn).Return(nil)

		ridesRepoMock.EXPECT().GetRiderMSISDNByID(ctx, rideData.RiderID).Return("", expectedErr)

		res, err := usecaseMock.StartRideDriver(ctx, req)
		assert.Error(t, err, pkgError.NewInternalServerError(expectedErr, "error get rider msisdn"))
		assert.Equal(t, model.RideData{}, res)
	})

	t.Run("failed - broadcast message returns error", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetDriverDataByMSISDN(ctx, driverMsisdn).Return(driverData, nil)
		ridesRepoMock.EXPECT().UpdateRideByDriver(ctx, model.UpdateRideByDriverRequest{
			DriverID: driverData.ID,
			RideID:   req.RideID,
			Status:   model.StatusNumRideStarted,
		}).Return(rideData, nil)
		locationRepoMock.EXPECT().RemoveAvailableDriver(ctx, driverMsisdn).Return(nil)

		ridesRepoMock.EXPECT().GetRiderMSISDNByID(ctx, rideData.RiderID).Return(riderMSISDN, nil)

		ridesPubsubMock.EXPECT().BroadcastMessage(ctx, constants.TopicRideStarted, model.RideStartedMessage{
			RideID:      rideData.RideID,
			RiderMSISDN: riderMSISDN,
		}).Return(expectedErr)

		res, err := usecaseMock.StartRideDriver(ctx, req)
		assert.Error(t, err, pkgError.NewInternalServerError(expectedErr, "error broadcasting matched ride to rider"))
		assert.Equal(t, model.RideData{}, res)
	})
}
