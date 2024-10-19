package handler_http

import (
	"nebeng-jek/internal/drivers/usecase"
	"sync"

	"github.com/gorilla/websocket"
)

type httpHandler struct {
	connStorage *sync.Map
	upgrader    websocket.Upgrader
	usecase     usecase.DriverUsecase
}

func NewHandler(connStorage *sync.Map, upgrader websocket.Upgrader, uc usecase.DriverUsecase) *httpHandler {
	return &httpHandler{
		connStorage: connStorage,
		upgrader:    upgrader,
		usecase:     uc,
	}
}
