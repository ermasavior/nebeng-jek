package handler_http

import (
	"context"
	"encoding/json"
	"errors"
	pkg_context "nebeng-jek/internal/pkg/context"
	"nebeng-jek/internal/pkg/location"
	"nebeng-jek/internal/riders/model"
	mock_usecase "nebeng-jek/mock/usecase"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestRiderWebsocket(t *testing.T) {
	path := "/ws"

	wsUpgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
	connStorage := &sync.Map{}

	ctrl := gomock.NewController(t)
	mockUC := mock_usecase.NewMockRiderUsecase(ctrl)
	handler := NewHandler(connStorage, wsUpgrader, mockUC)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET(path, handler.RiderWebsocket)

	server := httptest.NewServer(router)
	defer server.Close()

	u, err := url.Parse(server.URL)
	assert.NoError(t, err)

	u.Scheme = "ws"
	u.Path = path

	// Establish a WebSocket connection
	ws, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	assert.NoError(t, err)
	defer ws.Close()

	msg := model.RiderMessage{
		Event: "test",
	}
	err = ws.WriteJSON(msg)
	assert.NoError(t, err)
}

func TestHttpHandler_routeMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUC := mock_usecase.NewMockRiderUsecase(ctrl)

	h := NewHandler(&sync.Map{}, websocket.Upgrader{}, mockUC)

	driverID := int64(2222)
	ctx := pkg_context.SetRiderIDToContext(context.TODO(), driverID)

	req, _ := json.Marshal(map[string]interface{}{
		"ride_id":   111,
		"timestamp": 12345678,
		"location": map[string]float64{
			"longitude": 1.111,
			"latitude":  2.0001,
		},
	})

	t.Run("route real time location", func(t *testing.T) {
		msg := model.RiderMessage{
			Event: location.EventRealTimeLocation,
			Data:  req,
		}

		mockUC.EXPECT().TrackUserLocation(ctx, model.TrackUserLocationRequest{
			RideID:    111,
			Timestamp: 12345678,
			Location: location.Coordinate{
				Longitude: 1.111, Latitude: 2.0001,
			},
		}).Return(nil)

		h.routeMessage(ctx, msg)
	})

	t.Run("route real time location", func(t *testing.T) {
		msg := model.RiderMessage{
			Event: location.EventRealTimeLocation,
			Data:  req,
		}

		mockUC.EXPECT().TrackUserLocation(ctx, model.TrackUserLocationRequest{
			RideID:    111,
			Timestamp: 12345678,
			Location: location.Coordinate{
				Longitude: 1.111, Latitude: 2.0001,
			},
		}).Return(nil)

		h.routeMessage(ctx, msg)
	})

	t.Run("ignore - invalid data", func(t *testing.T) {
		msg := model.RiderMessage{
			Event: location.EventRealTimeLocation,
			Data:  []byte("/"),
		}

		h.routeMessage(ctx, msg)
	})

	t.Run("ignore - usecase returns error", func(t *testing.T) {
		msg := model.RiderMessage{
			Event: location.EventRealTimeLocation,
			Data:  req,
		}

		mockUC.EXPECT().TrackUserLocation(ctx, model.TrackUserLocationRequest{
			RideID:    111,
			Timestamp: 12345678,
			Location: location.Coordinate{
				Longitude: 1.111, Latitude: 2.0001,
			},
		}).Return(errors.New("error"))

		h.routeMessage(ctx, msg)
	})

}
