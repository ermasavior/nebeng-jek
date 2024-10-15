package handler

import (
	"context"
	"encoding/json"
	"nebeng-jek/internal/pkg/constants"
	"nebeng-jek/internal/rides/model"
	"nebeng-jek/pkg/amqp"
	"nebeng-jek/pkg/logger"
)

func (h *ridesHandler) SubscribeUserLocationTracking(ctx context.Context, amqpConn amqp.AMQPConnection) {
	channel, err := amqpConn.Channel()
	if err != nil {
		logger.Fatal(context.Background(), "error initializing amqp channel", map[string]interface{}{logger.ErrorKey: err})
	}
	defer channel.Close()

	msgs, err := amqp.ConsumeMessageToExchange(ctx, constants.UserLocationLiveTrackExchange, channel)
	if err != nil {
		logger.Fatal(ctx, "error consuming message to exchange", nil)
	}
	for msg := range msgs {
		var reqBody model.TrackUserLocationRequest
		err := json.Unmarshal(msg.Body, &reqBody)
		if err != nil {
			logger.Error(ctx, "fail to unmarshal consumed message", map[string]interface{}{"error": err})
		}

		pkgErr := h.usecase.TrackUserLocation(ctx, reqBody)
		if pkgErr != nil {
			logger.Error(ctx, "error tracking user location", map[string]interface{}{
				"req": reqBody,
			})
		}
	}
}
