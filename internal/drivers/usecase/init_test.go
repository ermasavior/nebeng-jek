package usecase

import (
	"context"
	"errors"
	"nebeng-jek/internal/drivers/model"
	"nebeng-jek/internal/pkg/constants"
	mockRepo "nebeng-jek/mock/repository"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUsecase_TrackUserLocation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ridesPubsubMock := mockRepo.NewMockRidesPubsubRepository(ctrl)
	usecaseMock := NewDriverUsecase(ridesPubsubMock)

	req := model.TrackUserLocationRequest{
		RideID:    1,
		UserID:    1111,
		Timestamp: time.Now().Unix(),
		Location: model.Coordinate{
			Latitude:  11,
			Longitude: 2,
		},
	}

	t.Run("success - broadcast track user location", func(t *testing.T) {
		ridesPubsubMock.EXPECT().BroadcastMessage(gomock.Any(), constants.TopicUserLiveLocation, req).Return(nil)
		err := usecaseMock.TrackUserLocation(context.Background(), req)
		assert.Nil(t, err)
	})

	t.Run("failed - repo return error", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesPubsubMock.EXPECT().BroadcastMessage(gomock.Any(), constants.TopicUserLiveLocation, req).Return(expectedErr)
		err := usecaseMock.TrackUserLocation(context.Background(), req)
		assert.Error(t, err)
	})
}
