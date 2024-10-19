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

func TestUsecase_CreateNewRide(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	locationRepoMock := mockRepo.NewMockRidesLocationRepository(ctrl)
	ridesRepoMock := mockRepo.NewMockRidesRepository(ctrl)
	ridesPubsubMock := mockRepo.NewMockRidesPubsubRepository(ctrl)
	usecaseMock := NewUsecase(locationRepoMock, ridesRepoMock, ridesPubsubMock, nil)

	var (
		msisdn    = "0811111"
		rideID    = int64(111)
		riderData = model.RiderData{
			ID:     1111,
			Name:   "Melati",
			MSISDN: "0812222",
		}
		driverList = []string{"021111", "021112"}
		driverMap  = map[string]bool{
			"021111": true,
			"021112": true,
		}

		req = model.CreateNewRideRequest{
			PickupLocation: model.Coordinate{
				Longitude: 10,
				Latitude:  10,
			},
			Destination: model.Coordinate{
				Longitude: 11,
				Latitude:  11,
			},
		}
	)

	ctx := context.Background()
	ctx = pkgContext.SetMSISDNToContext(ctx, msisdn)

	t.Run("success - should create new ride and publish message broadcast", func(t *testing.T) {
		ridesRepoMock.EXPECT().GetRiderDataByMSISDN(ctx, msisdn).Return(riderData, nil)

		locationRepoMock.EXPECT().GetNearestAvailableDrivers(ctx, req.PickupLocation).
			Return(driverList, nil)
		ridesRepoMock.EXPECT().CreateNewRide(ctx, model.CreateNewRideRequest{
			RiderID:        riderData.ID,
			PickupLocation: req.PickupLocation,
			Destination:    req.Destination,
		}).Return(rideID, nil)

		ridesPubsubMock.EXPECT().BroadcastMessage(ctx, constants.TopicRideNewRequest, model.RideRequestMessage{
			RideID:           rideID,
			Rider:            riderData,
			PickupLocation:   req.PickupLocation,
			Destination:      req.Destination,
			AvailableDrivers: driverMap,
		}).Return(nil)

		actual, err := usecaseMock.CreateNewRide(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, rideID, actual)
	})

	t.Run("failed - should return error when get rider id failed", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetRiderDataByMSISDN(ctx, msisdn).Return(model.RiderData{}, expectedErr)

		_, err := usecaseMock.CreateNewRide(ctx, req)
		assert.Equal(t, pkgError.NewInternalServerError(expectedErr, "error get rider data"), err)
	})

	t.Run("failed - should return error when get nearest available driver", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetRiderDataByMSISDN(ctx, msisdn).Return(riderData, nil)
		locationRepoMock.EXPECT().GetNearestAvailableDrivers(ctx, req.PickupLocation).
			Return(nil, expectedErr)

		_, err := usecaseMock.CreateNewRide(ctx, req)
		assert.Equal(t, pkgError.NewInternalServerError(expectedErr, "error get nearest available drivers"), err)
	})

	t.Run("failed - not found - should return error not found when there is no available driver", func(t *testing.T) {
		ridesRepoMock.EXPECT().GetRiderDataByMSISDN(ctx, msisdn).Return(riderData, nil)
		locationRepoMock.EXPECT().GetNearestAvailableDrivers(ctx, req.PickupLocation).
			Return([]string{}, nil)

		_, err := usecaseMock.CreateNewRide(ctx, req)
		assert.Equal(t, pkgError.NewNotFound(nil, "no nearest driver available, try again later"), err)
	})

	t.Run("failed - should return error when create new ride is failed", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetRiderDataByMSISDN(ctx, msisdn).Return(riderData, nil)
		locationRepoMock.EXPECT().GetNearestAvailableDrivers(ctx, req.PickupLocation).
			Return(driverList, nil)
		ridesRepoMock.EXPECT().CreateNewRide(ctx, model.CreateNewRideRequest{
			RiderID:        riderData.ID,
			PickupLocation: req.PickupLocation,
			Destination:    req.Destination,
		}).Return(int64(0), expectedErr)

		_, err := usecaseMock.CreateNewRide(ctx, req)
		assert.Equal(t, pkgError.NewInternalServerError(expectedErr, "error create new ride"), err)
	})

	t.Run("failed - should return error when fail broadcasting ride to drivers", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetRiderDataByMSISDN(ctx, msisdn).Return(riderData, nil)
		locationRepoMock.EXPECT().GetNearestAvailableDrivers(ctx, req.PickupLocation).
			Return(driverList, nil)
		ridesRepoMock.EXPECT().CreateNewRide(ctx, model.CreateNewRideRequest{
			RiderID:        riderData.ID,
			PickupLocation: req.PickupLocation,
			Destination:    req.Destination,
		}).Return(rideID, nil)

		ridesPubsubMock.EXPECT().BroadcastMessage(ctx, constants.TopicRideNewRequest, model.RideRequestMessage{
			RideID:           rideID,
			Rider:            riderData,
			PickupLocation:   req.PickupLocation,
			Destination:      req.Destination,
			AvailableDrivers: driverMap,
		}).Return(expectedErr)

		_, err := usecaseMock.CreateNewRide(ctx, req)
		assert.Equal(t, pkgError.NewInternalServerError(expectedErr, "error broadcasting ride to drivers"), err)
	})
}
