package handler

import (
	"context"
	"encoding/json"
	"errors"
	"nebeng-jek/internal/drivers/model"
	mock_ws "nebeng-jek/mock/pkg/websocket"
	"sync"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/websocket"
)

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

		handler.broadcastToActiveDrivers(ctx, availableDrivers, broadcastMsg)
	})

	t.Run("error - skip write message to websocket available drivers", func(t *testing.T) {
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

		mockConn1.EXPECT().WriteMessage(websocket.TextMessage, msgBytes).Return(errors.New("error"))
		mockConn2.EXPECT().WriteMessage(websocket.TextMessage, msgBytes).Return(nil)

		handler.broadcastToActiveDrivers(ctx, availableDrivers, broadcastMsg)
	})
}
