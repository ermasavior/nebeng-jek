package nats

import (
	"context"
	"nebeng-jek/pkg/logger"

	"github.com/nats-io/nats.go"
)

func NewNATSConnection(ctx context.Context, natsURL string) NATSConnection {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		logger.Error(ctx, "error connecting nats", map[string]interface{}{logger.ErrorKey: err})
		return nil
	}

	return nc
}

func NewNATSJSConnection(ctx context.Context, nc NATSConnection) JetStreamConnection {
	js, err := nc.JetStream()
	if err != nil {
		logger.Error(ctx, "error connecting nats jetstream", map[string]interface{}{logger.ErrorKey: err})
		return nil
	}

	return js
}
