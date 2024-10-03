package handler

import (
	"context"
	"encoding/json"
	"sync"
	"testing"

	"nebeng-jek/internal/pkg/constants"
	"nebeng-jek/internal/riders/model"
	mock_amqp "nebeng-jek/mock/pkg/amqp"
	mock_ws "nebeng-jek/mock/pkg/websocket"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/websocket"
	"github.com/rabbitmq/amqp091-go"
)

func TestSubscribeMatchedRides(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	amqpMock := mock_amqp.NewMockAMQPChannel(ctrl)
	h := &ridersHandler{
		connStorage: &sync.Map{},
	}

	t.Run("consume message from AMQP", func(t *testing.T) {
		msgs := make(chan amqp091.Delivery)

		amqpMock.EXPECT().ExchangeDeclare(constants.MatchedRideExchange, constants.ExchangeTypeFanout, true, false, false, false, nil).
			Return(nil)
		amqpMock.EXPECT().QueueDeclare("", false, false, true, false, nil).
			Return(amqp091.Queue{}, nil)
		amqpMock.EXPECT().QueueBind(gomock.Any(), "", constants.MatchedRideExchange, gomock.Any(), nil).
			Return(nil)
		amqpMock.EXPECT().Consume(gomock.Any(), gomock.Any(), true, false, false, false, nil).
			Return((<-chan amqp091.Delivery)(msgs), nil)

		go h.SubscribeMatchedRides(context.Background(), amqpMock)

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

func TestBroadcastToRider(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()

	rideID := int64(666)
	msisdn := "0822222"
	driverData := model.DriverData{
		ID:           1111,
		Name:         "Agus",
		MSISDN:       "081111",
		VehicleType:  "car",
		VehiclePlate: "B1212",
	}
	pickupLocation := model.Coordinate{
		Latitude:  1,
		Longitude: 2,
	}
	destination := model.Coordinate{
		Latitude:  9,
		Longitude: 10,
	}

	msg := model.MatchedRideMessage{
		RideID:         rideID,
		Driver:         driverData,
		PickupLocation: pickupLocation,
		Destination:    destination,
		RiderMSISDN:    msisdn,
	}

	connStorage := &sync.Map{}
	handler := &ridersHandler{connStorage: connStorage}

	mockConn := mock_ws.NewMockWebsocketInterface(ctrl)
	connStorage.Store(msisdn, mockConn)

	t.Run("write message to websocket rider", func(t *testing.T) {
		broadcastMsg := model.RiderMessage{
			Event: model.EventMatchedRide,
			Data: model.MatchedRideBroadcast{
				RideID:         rideID,
				Driver:         driverData,
				PickupLocation: pickupLocation,
				Destination:    destination,
			},
		}

		msgBytes, _ := json.Marshal(broadcastMsg)

		mockConn.EXPECT().WriteMessage(websocket.TextMessage, msgBytes).Return(nil)

		handler.broadcastToRider(ctx, msg)
	})
}
