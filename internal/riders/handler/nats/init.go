package handler_nats

import (
	"sync"
)

type natsHandler struct {
	connStorage *sync.Map
}

func NewHandler(connStorage *sync.Map) *natsHandler {
	return &natsHandler{
		connStorage: connStorage,
	}
}
