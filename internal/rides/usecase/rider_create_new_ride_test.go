package usecase

import (
	"context"
	"errors"
	"testing"

	"nebeng-jek/internal/pkg/constants"
	pkgContext "nebeng-jek/internal/pkg/context"
	pkgLocation "nebeng-jek/internal/pkg/location"
	"nebeng-jek/internal/rides/model"
	mockRepo "nebeng-jek/mock/repository"
	pkgError "nebeng-jek/pkg/error"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUsecase_RiderCreateNewRide(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	locationRepoMock := mockRepo.NewMockRidesLocationRepository(ctrl)
	ridesRepoMock := mockRepo.NewMockRidesRepository(ctrl)
	ridesPubsubMock := mockRepo.NewMockRidesPubsubRepository(ctrl)
	usecaseMock := NewUsecase(locationRepoMock, ridesRepoMock, ridesPubsubMock, nil)

	var (
		riderID   = int64(9999)
		rideID    = int64(111)
		riderData = model.RiderData{
			ID:     riderID,
			Name:   "Melati",
			MSISDN: "0812222",
		}
		driverList = []int64{1111, 2222}

		req = model.CreateNewRideRequest{
			PickupLocation: pkgLocation.Coordinate{
				Longitude: 10,
				Latitude:  10,
			},
			Destination: pkgLocation.Coordinate{
				Longitude: 11,
				Latitude:  11,
			},
		}
	)

	ctx := context.Background()
	ctx = pkgContext.SetRiderIDToContext(ctx, riderID)

	t.Run("success - should create new ride and publish message broadcast", func(t *testing.T) {
		ridesRepoMock.EXPECT().GetRiderDataByID(ctx, riderID).Return(riderData, nil)

		locationRepoMock.EXPECT().GetNearestAvailableDrivers(ctx, req.PickupLocation).
			Return(driverList, nil)
		ridesRepoMock.EXPECT().CreateNewRide(ctx, model.CreateNewRideRequest{
			RiderID:        riderData.ID,
			PickupLocation: req.PickupLocation,
			Destination:    req.Destination,
		}).Return(rideID, nil)

		ridesPubsubMock.EXPECT().BroadcastMessage(ctx, constants.TopicRideNewRequest, model.NewRideRequestMessage{
			RideID:            rideID,
			Rider:             riderData,
			PickupLocation:    req.PickupLocation,
			Destination:       req.Destination,
			AvailableDriverID: driverList[0],
		}).Return(nil).AnyTimes()
		ridesPubsubMock.EXPECT().BroadcastMessage(ctx, constants.TopicRideNewRequest, model.NewRideRequestMessage{
			RideID:            rideID,
			Rider:             riderData,
			PickupLocation:    req.PickupLocation,
			Destination:       req.Destination,
			AvailableDriverID: driverList[1],
		}).Return(nil).AnyTimes()

		actual, err := usecaseMock.RiderCreateNewRide(ctx, req)
		assert.Nil(t, err)
		assert.Equal(t, rideID, actual)
	})

	t.Run("failed - should return error when get rider id failed", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetRiderDataByID(ctx, riderID).Return(model.RiderData{}, expectedErr)

		_, err := usecaseMock.RiderCreateNewRide(ctx, req)
		assert.Equal(t, err.GetCode(), pkgError.ErrInternalErrorCode)
	})

	t.Run("failed - driver data is not found", func(t *testing.T) {
		ridesRepoMock.EXPECT().GetRiderDataByID(ctx, riderID).Return(model.RiderData{}, constants.ErrorDataNotFound)

		_, err := usecaseMock.RiderCreateNewRide(ctx, req)
		assert.Equal(t, err.GetCode(), pkgError.ErrUnauthorizedCode)
	})

	t.Run("failed - should return error when get nearest available driver", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetRiderDataByID(ctx, riderID).Return(riderData, nil)
		locationRepoMock.EXPECT().GetNearestAvailableDrivers(ctx, req.PickupLocation).
			Return(nil, expectedErr)

		_, err := usecaseMock.RiderCreateNewRide(ctx, req)
		assert.Equal(t, err.GetCode(), pkgError.ErrInternalErrorCode)
	})

	t.Run("failed - not found - should return error not found when there is no available driver", func(t *testing.T) {
		ridesRepoMock.EXPECT().GetRiderDataByID(ctx, riderID).Return(riderData, nil)
		locationRepoMock.EXPECT().GetNearestAvailableDrivers(ctx, req.PickupLocation).
			Return([]int64{}, nil)

		_, err := usecaseMock.RiderCreateNewRide(ctx, req)
		assert.Equal(t, err.GetCode(), pkgError.ErrResourceNotFoundCode)
	})

	t.Run("failed - should return error when create new ride is failed", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetRiderDataByID(ctx, riderID).Return(riderData, nil)
		locationRepoMock.EXPECT().GetNearestAvailableDrivers(ctx, req.PickupLocation).
			Return(driverList, nil)
		ridesRepoMock.EXPECT().CreateNewRide(ctx, model.CreateNewRideRequest{
			RiderID:        riderData.ID,
			PickupLocation: req.PickupLocation,
			Destination:    req.Destination,
		}).Return(int64(0), expectedErr)

		_, err := usecaseMock.RiderCreateNewRide(ctx, req)
		assert.Equal(t, err.GetCode(), pkgError.ErrInternalErrorCode)
	})

	t.Run("success - ignore when fail broadcasting ride to drivers", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesRepoMock.EXPECT().GetRiderDataByID(ctx, riderID).Return(riderData, nil)
		locationRepoMock.EXPECT().GetNearestAvailableDrivers(ctx, req.PickupLocation).
			Return(driverList, nil)
		ridesRepoMock.EXPECT().CreateNewRide(ctx, model.CreateNewRideRequest{
			RiderID:        riderData.ID,
			PickupLocation: req.PickupLocation,
			Destination:    req.Destination,
		}).Return(rideID, nil)

		ridesPubsubMock.EXPECT().BroadcastMessage(ctx, constants.TopicRideNewRequest, model.NewRideRequestMessage{
			RideID:            rideID,
			Rider:             riderData,
			PickupLocation:    req.PickupLocation,
			Destination:       req.Destination,
			AvailableDriverID: driverList[0],
		}).Return(expectedErr).AnyTimes()
		ridesPubsubMock.EXPECT().BroadcastMessage(ctx, constants.TopicRideNewRequest, model.NewRideRequestMessage{
			RideID:            rideID,
			Rider:             riderData,
			PickupLocation:    req.PickupLocation,
			Destination:       req.Destination,
			AvailableDriverID: driverList[1],
		}).Return(expectedErr).AnyTimes()

		_, err := usecaseMock.RiderCreateNewRide(ctx, req)
		assert.Nil(t, err)
	})
}
