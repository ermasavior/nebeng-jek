package handler_nats

import (
	"context"
	"encoding/json"
	"errors"
	"nebeng-jek/internal/drivers/model"
	"nebeng-jek/internal/pkg/location"
	mock_ws "nebeng-jek/mock/pkg/websocket"
	"sync"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
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
	pickupLocation := location.Coordinate{
		Latitude:  1,
		Longitude: 2,
	}
	destination := location.Coordinate{
		Latitude:  9,
		Longitude: 10,
	}
	availableDriver := int64(1111)

	data, _ := json.Marshal(model.NewRideRequestBroadcast{
		RideID:         rideID,
		Rider:          riderData,
		PickupLocation: pickupLocation,
		Destination:    destination,
	})

	connStorage := &sync.Map{}
	mockConn1 := mock_ws.NewMockWebsocketInterface(ctrl)
	connStorage.Store(int64(1111), mockConn1)

	handler := NewHandler(connStorage, nil)

	t.Run("write message to websocket available drivers", func(t *testing.T) {
		broadcastMsg := model.DriverMessage{
			Event: model.EventNewRideRequest,
			Data:  data,
		}

		msgBytes, _ := json.Marshal(broadcastMsg)

		mockConn1.EXPECT().WriteMessage(websocket.TextMessage, msgBytes).Return(nil)

		err := handler.broadcastToDriver(ctx, availableDriver, broadcastMsg)
		assert.Nil(t, err)
	})

	t.Run("error - returns error", func(t *testing.T) {
		broadcastMsg := model.DriverMessage{
			Event: model.EventNewRideRequest,
			Data:  data,
		}

		msgBytes, _ := json.Marshal(broadcastMsg)

		mockConn1.EXPECT().WriteMessage(websocket.TextMessage, msgBytes).Return(errors.New("error"))

		err := handler.broadcastToDriver(ctx, availableDriver, broadcastMsg)
		assert.Error(t, err)
	})
}
