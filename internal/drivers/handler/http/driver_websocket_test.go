package handler_http

import (
	"context"
	"nebeng-jek/internal/drivers/model"
	pkg_context "nebeng-jek/internal/pkg/context"
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

func TestDriverWebsocket(t *testing.T) {
	path := "/ws"

	wsUpgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
	handler := NewHandler(nil, wsUpgrader, nil)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET(path, handler.DriverAllocationWebsocket)

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

	msg := model.DriverMessage{
		Event: "test",
	}
	err = ws.WriteJSON(msg)
	assert.NoError(t, err)
}

func TestHttpHandler_routeMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUC := mock_usecase.NewMockDriverUsecase(ctrl)

	h := NewHandler(&sync.Map{}, websocket.Upgrader{}, mockUC)

	driverID := int64(2222)
	ctx := pkg_context.SetDriverIDToContext(context.TODO(), driverID)

	t.Run("route real time location", func(t *testing.T) {
		msg := model.DriverMessage{
			Event: model.EventRealTimeLocation,
			Data: map[string]interface{}{
				"ride_id":   111,
				"timestamp": 12345678,
				"location": map[string]float64{
					"longitude": 1.111,
					"latitude":  2.0001,
				},
			},
		}

		mockUC.EXPECT().TrackUserLocation(ctx, model.TrackUserLocationRequest{
			RideID:    111,
			UserID:    driverID,
			Timestamp: 12345678,
			Location: model.Coordinate{
				Longitude: 1.111, Latitude: 2.0001,
			},
		}).Return(nil)

		h.routeMessage(ctx, msg)
	})
}
