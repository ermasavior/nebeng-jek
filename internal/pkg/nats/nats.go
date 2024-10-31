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
			"topic":         topic,
			"msg":           msg,
			logger.ErrorKey: err,
		})
		return err
	}

	_, err = r.nats.PublishAsync(topic, msgBytes)
	if err != nil {
		logger.Error(ctx, "error publish message", map[string]interface{}{
			"topic":         topic,
			"msg":           msg,
			logger.ErrorKey: err,
		})
		return err
	}

	return nil
}

func SubscribeMessage(natsJS nats.JetStreamConnection, topic string, msgHandler nats_go.MsgHandler, consumerName string) {
	_, err := natsJS.Subscribe(topic, msgHandler,
		nats_go.MaxDeliver(5), nats_go.AckExplicit(), nats_go.Durable(consumerName))
	if err != nil {
		logger.Error(context.Background(), "fail to subscribe messages", map[string]interface{}{
			"topic": topic,
			"error": err,
		})
		return
	}

	select {}
}

func AckMessage(ctx context.Context, msg *nats_go.Msg) {
	if err := msg.Ack(); err != nil {
		logger.Error(ctx, "error ack message", map[string]interface{}{"error": err})
	}
}

func NakMessage(ctx context.Context, msg *nats_go.Msg) {
	if err := msg.Nak(); err != nil {
		logger.Error(ctx, "error nak message", map[string]interface{}{"error": err})
	}
}
