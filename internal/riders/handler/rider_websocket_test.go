package handler

import (
	"nebeng-jek/internal/riders/model"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestRiderAllocationConnection(t *testing.T) {
	path := "/ws"

	handler := ridersHandler{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(r *http.Request) bool { return true },
		},
	}

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
