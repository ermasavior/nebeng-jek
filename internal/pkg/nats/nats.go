package nats

import (
	"context"
	"encoding/json"
	"nebeng-jek/internal/drivers/repository"
	"nebeng-jek/pkg/logger"
	"nebeng-jek/pkg/messaging/nats"

	nats_go "github.com/nats-io/nats.go"
)

type pubsubRepo struct {
	nats nats.JetStreamConnection
}

func NewPubsubRepository(nats nats.JetStreamConnection) repository.RidesPubsubRepository {
	return &pubsubRepo{
		nats: nats,
	}
}

func (r *pubsubRepo) BroadcastMessage(ctx context.Context, topic string, msg interface{}) error {
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		logger.Error(ctx, "error marshalling message bytes", map[string]interface{}{
			"topic": topic,
			"msg":   msg,
		})
		return err
	}

	_, err = r.nats.Publish(topic, msgBytes)
	if err != nil {
		logger.Error(ctx, "error publish message", map[string]interface{}{
			"topic": topic,
			"msg":   msg,
		})
		return err
	}

	return nil
}

func SubscribeMessage(natsJS nats.JetStreamConnection, topic string, msgHandler nats_go.MsgHandler) {
	_, err := natsJS.Subscribe(topic, msgHandler, nats_go.MaxDeliver(5))
	if err != nil {
		logger.Error(context.Background(), "fail to subscribe messages", map[string]interface{}{"error": err})
		return
	}

	select {}
}
