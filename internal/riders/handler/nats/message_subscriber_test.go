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
			ID:     666,
			Name:   "Mel",
			MSISDN: "0812222",
		},
		RiderMSISDN: "0812222",
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

func TestSubscribeReadyToPickupRides(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	h := &natsHandler{
		connStorage: &sync.Map{},
	}

	msgBytes, _ := json.Marshal(model.RideReadyToPickupMessage{
		RideID: 111,
		PickupLocation: model.Coordinate{
			Longitude: 111,
			Latitude:  -22,
		},
		Destination: model.Coordinate{
			Longitude: 113,
			Latitude:  -21,
		},
		DriverMSISDN: "0821111",
		RiderMSISDN:  "0812222",
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

func TestSubscribeRideStarted(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	h := &natsHandler{
		connStorage: &sync.Map{},
	}

	msgBytes, _ := json.Marshal(model.RideStartedMessage{
		RideID:      111,
		RiderMSISDN: "0812222",
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
		RideID:      111,
		Distance:    10,
		Fare:        30000,
		RiderMSISDN: "0812222",
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
		RideID:      111,
		Distance:    10,
		FinalPrice:  20000,
		RiderMSISDN: "0812222",
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
