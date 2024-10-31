package handler_nats

import (
	"context"
	"encoding/json"
	"sync"
	"testing"

	"nebeng-jek/internal/drivers/model"
	"nebeng-jek/internal/pkg/constants"

	"github.com/golang/mock/gomock"
	"github.com/nats-io/nats.go"
)

func TestSubscribeNewRideRequests(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	h := NewHandler(&sync.Map{}, nil)

	msg := model.NewRideRequestMessage{
		RideID: 111,
		Rider: model.RiderData{
			Name:   "Mel",
			MSISDN: "0812222",
		},
		AvailableDrivers: map[int64]bool{
			1111: true,
		},
	}
	msgBytes, _ := json.Marshal(msg)

	t.Run("subscribe message from NATS", func(t *testing.T) {
		mockMsg := &nats.Msg{
			Subject: constants.TopicRideNewRequest,
			Data:    msgBytes,
		}

		handler := h.SubscribeNewRideRequests(context.Background())
		handler(mockMsg)
	})
}

func TestSubscribeReadyToPickupRides(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	h := NewHandler(&sync.Map{}, nil)

	msgBytes, _ := json.Marshal(model.RideReadyToPickupMessage{
		RideID:   111,
		DriverID: 1111,
		RiderID:  9999,
	})

	t.Run("consume message from NATS", func(t *testing.T) {
		mockMsg := &nats.Msg{
			Subject: constants.TopicRideReadyToPickup,
			Data:    msgBytes,
		}

		handler := h.SubscribeReadyToPickupRides(context.Background())
		handler(mockMsg)
	})
}
