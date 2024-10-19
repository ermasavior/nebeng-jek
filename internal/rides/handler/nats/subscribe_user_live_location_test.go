package handler_nats

import (
	"context"
	"encoding/json"
	"errors"
	"nebeng-jek/internal/pkg/constants"
	"nebeng-jek/internal/rides/model"
	mock_usecase "nebeng-jek/mock/usecase"
	errorPkg "nebeng-jek/pkg/error"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/nats-io/nats.go"
)

func TestSubscribeUserLiveLocation(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock_usecase.NewMockRidesUsecase(ctrl)
	h := NewHandler(mockUsecase)

	req := model.TrackUserLocationRequest{
		RideID:    111,
		MSISDN:    "081234",
		Timestamp: time.Now().Unix(),
		Location: model.Coordinate{
			Longitude: 11,
			Latitude:  -18,
		},
		IsDriver: true,
	}
	msgBytes, _ := json.Marshal(req)

	mockMsg := &nats.Msg{
		Subject: constants.TopicUserLiveLocation,
		Data:    msgBytes,
	}

	ctx := context.Background()

	t.Run("subscribed user live location and track it", func(t *testing.T) {
		mockUsecase.EXPECT().TrackUserLocation(ctx, req).Return(nil)
		handler := h.SubscribeUserLiveLocation(ctx)
		handler(mockMsg)
	})

	t.Run("error - failed", func(t *testing.T) {
		expectedError := errorPkg.NewInternalServerError(errors.New("error"), "error from usecase")
		mockUsecase.EXPECT().TrackUserLocation(ctx, req).Return(expectedError)
		handler := h.SubscribeUserLiveLocation(ctx)
		handler(mockMsg)
	})
}
