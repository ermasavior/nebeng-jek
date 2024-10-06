package handler

import (
	"context"
	"encoding/json"
	"nebeng-jek/internal/riders/model"
	mock_ws "nebeng-jek/mock/pkg/websocket"
	"sync"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/websocket"
)

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

	broadcastMsg := model.RiderMessage{
		Event: model.EventMatchedRide,
		Data: model.MatchedRideMessage{
			RideID:      rideID,
			Driver:      driverData,
			RiderMSISDN: msisdn,
		},
	}

	connStorage := &sync.Map{}
	handler := &ridersHandler{connStorage: connStorage}

	mockConn := mock_ws.NewMockWebsocketInterface(ctrl)
	connStorage.Store(msisdn, mockConn)

	t.Run("write message to websocket rider", func(t *testing.T) {
		msgBytes, _ := json.Marshal(broadcastMsg)
		mockConn.EXPECT().WriteMessage(websocket.TextMessage, msgBytes).Return(nil)

		handler.broadcastToRider(ctx, msisdn, broadcastMsg)
	})
}
