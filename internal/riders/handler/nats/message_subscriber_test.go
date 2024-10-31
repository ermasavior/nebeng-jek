package handler_nats

import (
	"context"
	"encoding/json"
	"sync"
	"testing"

	"nebeng-jek/internal/pkg/constants"
	"nebeng-jek/internal/riders/model"

	"github.com/golang/mock/gomock"
	"github.com/nats-io/nats.go"
)

func TestSubscribeRideMatchedDriver(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	h := &natsHandler{
		connStorage: &sync.Map{},
	}

	msgBytes, _ := json.Marshal(model.RideMatchedDriverMessage{
		RideID: 111,
		Driver: model.DriverData{
			Name:   "Mel",
			MSISDN: "0812222",
		},
		RiderID: 9999,
	})

	t.Run("consume message from NATS", func(t *testing.T) {
		mockMsg := &nats.Msg{
			Subject: constants.TopicRideMatchedDriver,
			Data:    msgBytes,
		}

		handler := h.SubscribeRideMatchedDriver(context.Background())
		handler(mockMsg)
	})
}

func TestSubscribeRideStarted(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	h := &natsHandler{
		connStorage: &sync.Map{},
	}

	msgBytes, _ := json.Marshal(model.RideStartedMessage{
		RideID:  111,
		RiderID: 9999,
	})

	t.Run("consume message from NATS", func(t *testing.T) {
		mockMsg := &nats.Msg{
			Subject: constants.TopicRideStarted,
			Data:    msgBytes,
		}

		handler := h.SubscribeRideStarted(context.Background())
		handler(mockMsg)
	})
}

func TestSubscribeRideEnded(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	h := &natsHandler{
		connStorage: &sync.Map{},
	}

	msgBytes, _ := json.Marshal(model.RideEndedMessage{
		RideID:   111,
		Distance: 10,
		Fare:     30000,
		RiderID:  9999,
	})

	t.Run("consume message from NATS", func(t *testing.T) {
		mockMsg := &nats.Msg{
			Subject: constants.TopicRideEnded,
			Data:    msgBytes,
		}

		handler := h.SubscribeRideEnded(context.Background())
		handler(mockMsg)
	})
}

func TestSubscribeRidePaid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	h := &natsHandler{
		connStorage: &sync.Map{},
	}

	msgBytes, _ := json.Marshal(model.RidePaidMessage{
		RideID:     111,
		Distance:   10,
		FinalPrice: 20000,
		RiderID:    9999,
	})

	t.Run("consume message from NATS", func(t *testing.T) {
		mockMsg := &nats.Msg{
			Subject: constants.TopicRidePaid,
			Data:    msgBytes,
		}

		handler := h.SubscribeRidePaid(context.Background())
		handler(mockMsg)
	})
}
