package usecase

import (
	"context"
	"errors"
	"nebeng-jek/internal/drivers/model"
	"nebeng-jek/internal/pkg/constants"
	pkg_context "nebeng-jek/internal/pkg/context"
	"nebeng-jek/internal/pkg/location"
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

	ctx := pkg_context.SetDriverIDToContext(context.Background(), 1111)
	req := model.TrackUserLocationRequest{
		RideID:    1,
		Timestamp: time.Now().Unix(),
		Location: location.Coordinate{
			Latitude:  11,
			Longitude: 2,
		},
	}
	msg := location.TrackUserLocationMessage{
		RideID:    1,
		UserID:    1111,
		Timestamp: time.Now().Unix(),
		Location: location.Coordinate{
			Latitude:  11,
			Longitude: 2,
		},
		IsDriver: true,
	}

	t.Run("success - broadcast track user location", func(t *testing.T) {
		ridesPubsubMock.EXPECT().BroadcastMessage(ctx, constants.TopicUserLiveLocation, msg).Return(nil)
		err := usecaseMock.TrackUserLocation(ctx, req)
		assert.Nil(t, err)
	})

	t.Run("failed - repo return error", func(t *testing.T) {
		expectedErr := errors.New("error from repo")
		ridesPubsubMock.EXPECT().BroadcastMessage(ctx, constants.TopicUserLiveLocation, msg).Return(expectedErr)
		err := usecaseMock.TrackUserLocation(ctx, req)
		assert.Error(t, err)
	})
}
