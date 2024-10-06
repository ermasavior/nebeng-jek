package handler

import (
	"context"
	"encoding/json"
	"sync"
	"testing"

	"nebeng-jek/internal/pkg/constants"
	"nebeng-jek/internal/riders/model"
	mock_amqp "nebeng-jek/mock/pkg/amqp"

	"github.com/golang/mock/gomock"
	"github.com/rabbitmq/amqp091-go"
)

func TestSubscribeDriverAcceptedRides(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	amqpMock := mock_amqp.NewMockAMQPChannel(ctrl)
	h := &ridersHandler{
		connStorage: &sync.Map{},
	}

	t.Run("consume message from AMQP", func(t *testing.T) {
		msgs := make(chan amqp091.Delivery)

		amqpMock.EXPECT().ExchangeDeclare(constants.DriverAcceptedRideExchange, constants.ExchangeTypeFanout, true, false, false, false, nil).
			Return(nil)
		amqpMock.EXPECT().QueueDeclare("", false, false, true, false, nil).
			Return(amqp091.Queue{}, nil)
		amqpMock.EXPECT().QueueBind(gomock.Any(), "", constants.DriverAcceptedRideExchange, gomock.Any(), nil).
			Return(nil)
		amqpMock.EXPECT().Consume(gomock.Any(), gomock.Any(), true, false, false, false, nil).
			Return((<-chan amqp091.Delivery)(msgs), nil)

		go h.SubscribeDriverAcceptedRides(context.Background(), amqpMock)

		// Simulate a message being received
		msgBody, _ := json.Marshal(model.MatchedRideMessage{
			RideID: 111,
			Driver: model.DriverData{
				ID:     666,
				Name:   "Mel",
				MSISDN: "0812222",
			},
			RiderMSISDN: "0812222",
		})
		msgs <- amqp091.Delivery{Body: msgBody}

		close(msgs)

	})
}

func TestSubscribeReadyToPickupRides(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	amqpMock := mock_amqp.NewMockAMQPChannel(ctrl)
	h := &ridersHandler{
		connStorage: &sync.Map{},
	}

	t.Run("consume message from AMQP", func(t *testing.T) {
		msgs := make(chan amqp091.Delivery)

		amqpMock.EXPECT().ExchangeDeclare(constants.RideReadyToPickupExchange, constants.ExchangeTypeFanout, true, false, false, false, nil).
			Return(nil)
		amqpMock.EXPECT().QueueDeclare("", false, false, true, false, nil).
			Return(amqp091.Queue{}, nil)
		amqpMock.EXPECT().QueueBind(gomock.Any(), "", constants.RideReadyToPickupExchange, gomock.Any(), nil).
			Return(nil)
		amqpMock.EXPECT().Consume(gomock.Any(), gomock.Any(), true, false, false, false, nil).
			Return((<-chan amqp091.Delivery)(msgs), nil)

		go h.SubscribeReadyToPickupRides(context.Background(), amqpMock)

		// Simulate a message being received
		msgBody, _ := json.Marshal(model.RideReadyToPickupMessage{
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
		msgs <- amqp091.Delivery{Body: msgBody}

		close(msgs)

	})
}
