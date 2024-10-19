package usecase

import (
	"context"
	"errors"
	"nebeng-jek/internal/rides/model"
	mockRepo "nebeng-jek/mock/repository"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUsecase_TrackUserLocation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	locationRepoMock := mockRepo.NewMockRidesLocationRepository(ctrl)
	ridesRepoMock := mockRepo.NewMockRidesRepository(ctrl)
	ridesPubsubMock := mockRepo.NewMockRidesPubsubRepository(ctrl)
	usecaseMock := NewUsecase(locationRepoMock, ridesRepoMock, ridesPubsubMock, nil)

	req := model.TrackUserLocationRequest{
		RideID:    1,
		UserID:    1111,
		Timestamp: time.Now().Unix(),
		Location: model.Coordinate{
			Latitude:  11,
			Longitude: 2,
		},
	}

	t.Run("success - track user location", func(t *testing.T) {
		locationRepoMock.EXPECT().TrackUserLocation(gomock.Any(), req).Return(nil)
		err := usecaseMock.TrackUserLocation(context.Background(), req)
		assert.Nil(t, err)
	})

	t.Run("failed - repo return error", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		locationRepoMock.EXPECT().TrackUserLocation(gomock.Any(), req).Return(expectedErr)
		err := usecaseMock.TrackUserLocation(context.Background(), req)
		assert.EqualError(t, expectedErr, err.Raw.Error())
	})
}
