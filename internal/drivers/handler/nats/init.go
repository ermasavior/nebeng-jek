package handler_nats

import (
	"nebeng-jek/internal/drivers/usecase"
	"sync"
)

type natsHandler struct {
	connStorage *sync.Map
	usecase     usecase.DriverUsecase
}

func NewHandler(connStorage *sync.Map, uc usecase.DriverUsecase) *natsHandler {
	return &natsHandler{
		connStorage: connStorage,
		usecase:     uc,
	}
}
