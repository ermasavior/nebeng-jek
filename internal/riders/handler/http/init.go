package handler_http

import (
	"sync"

	"github.com/gorilla/websocket"
)

type httpHandler struct {
	connStorage *sync.Map
	upgrader    websocket.Upgrader
}

func NewHandler(connStorage *sync.Map, upgrader websocket.Upgrader) *httpHandler {
	return &httpHandler{
		connStorage: connStorage,
		upgrader:    upgrader,
	}
}
