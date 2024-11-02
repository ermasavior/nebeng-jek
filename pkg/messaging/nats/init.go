package nats

import (
	"context"
	"log"

	"github.com/nats-io/nats.go"
)

func NewNATSConnection(ctx context.Context, natsURL string) NATSConnection {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatal("error connecting nats ", err)
		return nil
	}

	return nc
}

func NewNATSJSConnection(ctx context.Context, nc NATSConnection) JetStreamConnection {
	js, err := nc.JetStream()
	if err != nil {
		log.Fatal("error connecting nats jetstream ", err)
		return nil
	}

	return js
}
