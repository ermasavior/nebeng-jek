package nats

import (
	"context"
	"log"
	"nebeng-jek/pkg/logger"

	"github.com/nats-io/nats.go"
)

func NewNATSConnection(natsURL string) NATSConnection {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		logger.Fatal(context.Background(), err.Error(), nil)
	}

	return nc
}

func NewNATSJSConnection(nc NATSConnection) JetStreamConnection {
	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	return js
}
