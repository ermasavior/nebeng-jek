package handler

import (
	"context"
	"encoding/json"
	"sync"
	"testing"

	"nebeng-jek/internal/drivers/model"
	"nebeng-jek/internal/pkg/constants"
	mock_amqp "nebeng-jek/mock/pkg/amqp"
	mock_ws "nebeng-jek/mock/pkg/websocket"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/websocket"
	"github.com/rabbitmq/amqp091-go"
)

func TestSubscribeNewRides(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	amqpMock := mock_amqp.NewMockAMQPChannel(ctrl)
	h := &driversHandler{
		connStorage: &sync.Map{},
	}

	t.Run("consume message from AMQP", func(t *testing.T) {
		msgs := make(chan amqp091.Delivery)

		amqpMock.EXPECT().ExchangeDeclare(constants.RideRequestsExchange, constants.ExchangeTypeFanout, true, false, false, false, nil).
			Return(nil)
		amqpMock.EXPECT().QueueDeclare("", false, false, true, false, nil).
			Return(amqp091.Queue{}, nil)
		amqpMock.EXPECT().QueueBind(gomock.Any(), "", constants.RideRequestsExchange, gomock.Any(), nil).
			Return(nil)
		amqpMock.EXPECT().Consume(gomock.Any(), gomock.Any(), true, false, false, false, nil).
			Return((<-chan amqp091.Delivery)(msgs), nil)

		go h.SubscribeNewRides(context.Background(), amqpMock)

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

func TestBroadcastToActiveDrivers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()

	rideID := int64(666)
	riderData := model.RiderData{
		ID:     1111,
		Name:   "Mel",
		MSISDN: "082222",
	}
	pickupLocation := model.Coordinate{
		Latitude:  1,
		Longitude: 2,
	}
	destination := model.Coordinate{
		Latitude:  9,
		Longitude: 10,
	}
	availableDrivers := map[string]bool{
		"0812": true,
		"0813": true,
	}

	msg := model.NewRideRequestMessage{
		RideID:           rideID,
		Rider:            riderData,
		PickupLocation:   pickupLocation,
		Destination:      destination,
		AvailableDrivers: availableDrivers,
	}

	connStorage := &sync.Map{}
	handler := &driversHandler{connStorage: connStorage}

	mockConn1 := mock_ws.NewMockWebsocketInterface(ctrl)
	mockConn2 := mock_ws.NewMockWebsocketInterface(ctrl)

	connStorage.Store("0812", mockConn1)
	connStorage.Store("0813", mockConn2)

	t.Run("write message to websocket available drivers", func(t *testing.T) {
		broadcastMsg := model.DriverMessage{
			Event: model.EventNewRideRequest,
			Data: model.NewRideRequestBroadcast{
				RideID:         rideID,
				Rider:          riderData,
				PickupLocation: pickupLocation,
				Destination:    destination,
			},
		}

		msgBytes, _ := json.Marshal(broadcastMsg)

		mockConn1.EXPECT().WriteMessage(websocket.TextMessage, msgBytes).Return(nil)
		mockConn2.EXPECT().WriteMessage(websocket.TextMessage, msgBytes).Return(nil)

		handler.broadcastToActiveDrivers(ctx, msg)
	})
}
