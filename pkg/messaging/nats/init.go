package nats

import (
	"context"
	"nebeng-jek/pkg/logger"

	"github.com/nats-io/nats.go"
)

func NewNATSConnection(ctx context.Context, natsURL string) NATSConnection {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		logger.Fatal(ctx, err.Error(), nil)
	}

	return nc
}

func NewNATSJSConnection(ctx context.Context, nc NATSConnection) JetStreamConnection {
	js, err := nc.JetStream()
	if err != nil {
		logger.Fatal(ctx, err.Error(), nil)
	}

	return js
}
