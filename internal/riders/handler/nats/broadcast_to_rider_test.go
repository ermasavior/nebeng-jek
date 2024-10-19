package handler_nats

import (
	"context"
	"encoding/json"
	"errors"
	"nebeng-jek/internal/riders/model"
	mock_ws "nebeng-jek/mock/pkg/websocket"
	"sync"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestBroadcastToRider(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.TODO()

	rideID := int64(666)
	riderID := int64(9999)
	driverData := model.DriverData{
		Name:         "Agus",
		MSISDN:       "081111",
		VehicleType:  "car",
		VehiclePlate: "B1212",
	}

	broadcastMsg := model.RiderMessage{
		Event: model.EventMatchedRide,
		Data: model.RideMatchedDriverMessage{
			RideID:  rideID,
			Driver:  driverData,
			RiderID: riderID,
		},
	}

	connStorage := &sync.Map{}
	handler := NewHandler(connStorage)

	mockConn := mock_ws.NewMockWebsocketInterface(ctrl)
	connStorage.Store(riderID, mockConn)

	t.Run("write message to websocket rider", func(t *testing.T) {
		msgBytes, _ := json.Marshal(broadcastMsg)
		mockConn.EXPECT().WriteMessage(websocket.TextMessage, msgBytes).Return(nil)

		err := handler.broadcastToRider(ctx, riderID, broadcastMsg)
		assert.NoError(t, err)
	})

	t.Run("write message to websocket rider", func(t *testing.T) {
		expectedErr := errors.New("expected error")

		msgBytes, _ := json.Marshal(broadcastMsg)
		mockConn.EXPECT().WriteMessage(websocket.TextMessage, msgBytes).Return(expectedErr)

		err := handler.broadcastToRider(ctx, riderID, broadcastMsg)
		assert.Error(t, err)
	})
}
