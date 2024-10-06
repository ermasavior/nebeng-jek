package handler

import (
	"context"
	"encoding/json"
	"sync"
	"testing"

	"nebeng-jek/internal/drivers/model"
	"nebeng-jek/internal/pkg/constants"
	mock_amqp "nebeng-jek/mock/pkg/amqp"

	"github.com/golang/mock/gomock"
	"github.com/rabbitmq/amqp091-go"
)

func TestSubscribeNewRideRequests(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	amqpMock := mock_amqp.NewMockAMQPChannel(ctrl)
	h := &driversHandler{
		connStorage: &sync.Map{},
	}

	t.Run("consume message from AMQP", func(t *testing.T) {
		msgs := make(chan amqp091.Delivery)

		amqpMock.EXPECT().ExchangeDeclare(constants.NewRideRequestsExchange, constants.ExchangeTypeFanout, true, false, false, false, nil).
			Return(nil)
		amqpMock.EXPECT().QueueDeclare("", false, false, true, false, nil).
			Return(amqp091.Queue{}, nil)
		amqpMock.EXPECT().QueueBind(gomock.Any(), "", constants.NewRideRequestsExchange, gomock.Any(), nil).
			Return(nil)
		amqpMock.EXPECT().Consume(gomock.Any(), gomock.Any(), true, false, false, false, nil).
			Return((<-chan amqp091.Delivery)(msgs), nil)

		go h.SubscribeNewRideRequests(context.Background(), amqpMock)

		// Simulate a message being received
		msgBody, _ := json.Marshal(model.NewRideRequestMessage{
			RideID: 111,
			Rider: model.RiderData{
				ID:     666,
				Name:   "Mel",
				MSISDN: "0812222",
			},
			AvailableDrivers: map[string]bool{
				"081": true,
			},
		})
		msgs <- amqp091.Delivery{Body: msgBody}

		close(msgs)

	})
}

func TestSubscribeReadyToPickupRides(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	amqpMock := mock_amqp.NewMockAMQPChannel(ctrl)
	h := &driversHandler{
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
				Longitude: 11,
				Latitude:  -2,
			},
			Destination: model.Coordinate{
				Longitude: 10,
				Latitude:  -1,
			},
			DriverMSISDN: "081222",
			RiderMSISDN:  "082111",
		})
		msgs <- amqp091.Delivery{Body: msgBody}

		close(msgs)

	})
}
