package handler_nats

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

func Test_broadcastToDrivers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()

	rideID := int64(666)
	riderData := model.RiderData{
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
	availableDrivers := map[int64]bool{
		1111: true,
		2222: true,
	}

	connStorage := &sync.Map{}
	mockConn1 := mock_ws.NewMockWebsocketInterface(ctrl)
	mockConn2 := mock_ws.NewMockWebsocketInterface(ctrl)
	connStorage.Store(int64(1111), mockConn1)
	connStorage.Store(int64(2222), mockConn2)

	handler := NewHandler(connStorage, nil)

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

		handler.broadcastToDrivers(ctx, availableDrivers, broadcastMsg)
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

		handler.broadcastToDrivers(ctx, availableDrivers, broadcastMsg)
	})
}
