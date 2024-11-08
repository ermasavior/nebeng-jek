package handler_http

import (
	"nebeng-jek/internal/riders/usecase"
	"sync"

	"github.com/gorilla/websocket"
)

type httpHandler struct {
	connStorage *sync.Map
	upgrader    websocket.Upgrader
	usecase     usecase.RiderUsecase
}

func NewHandler(connStorage *sync.Map, upgrader websocket.Upgrader, usecase usecase.RiderUsecase) *httpHandler {
	return &httpHandler{
		connStorage: connStorage,
		upgrader:    upgrader,
		usecase:     usecase,
	}
}
