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

	amqpConn := mock_amqp.NewMockAMQPConnection(ctrl)
	amqpChan := mock_amqp.NewMockAMQPChannel(ctrl)
	h := &driversHandler{
		connStorage: &sync.Map{},
	}

	t.Run("consume message from AMQP", func(t *testing.T) {
		msgs := make(chan amqp091.Delivery)

		amqpConn.EXPECT().Channel().Return(amqpChan, nil)
		amqpChan.EXPECT().ExchangeDeclare(constants.NewRideRequestsExchange, constants.ExchangeTypeFanout, true, false, false, false, nil).
			Return(nil)
		amqpChan.EXPECT().QueueDeclare("", false, false, true, false, nil).
			Return(amqp091.Queue{}, nil)
		amqpChan.EXPECT().QueueBind(gomock.Any(), "", constants.NewRideRequestsExchange, gomock.Any(), nil).
			Return(nil)
		amqpChan.EXPECT().Consume(gomock.Any(), gomock.Any(), true, false, false, false, nil).
			Return((<-chan amqp091.Delivery)(msgs), nil)
		amqpChan.EXPECT().Close().Return(nil)

		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			h.SubscribeNewRideRequests(context.Background(), amqpConn)
		}()

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
		wg.Wait()
	})
}

func TestSubscribeReadyToPickupRides(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	amqpConn := mock_amqp.NewMockAMQPConnection(ctrl)
	amqpChan := mock_amqp.NewMockAMQPChannel(ctrl)
	h := &driversHandler{
		connStorage: &sync.Map{},
	}

	t.Run("consume message from AMQP", func(t *testing.T) {
		msgs := make(chan amqp091.Delivery)

		amqpConn.EXPECT().Channel().Return(amqpChan, nil)
		amqpChan.EXPECT().ExchangeDeclare(constants.RideReadyToPickupExchange, constants.ExchangeTypeFanout, true, false, false, false, nil).
			Return(nil)
		amqpChan.EXPECT().QueueDeclare("", false, false, true, false, nil).
			Return(amqp091.Queue{}, nil)
		amqpChan.EXPECT().QueueBind(gomock.Any(), "", constants.RideReadyToPickupExchange, gomock.Any(), nil).
			Return(nil)
		amqpChan.EXPECT().Consume(gomock.Any(), gomock.Any(), true, false, false, false, nil).
			Return((<-chan amqp091.Delivery)(msgs), nil)
		amqpChan.EXPECT().Close().Return(nil)

		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			h.SubscribeReadyToPickupRides(context.Background(), amqpConn)
		}()

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
		wg.Wait()
	})
}
