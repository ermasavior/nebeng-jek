package handler

import (
	"context"
	"encoding/json"
	"nebeng-jek/internal/pkg/constants"
	"nebeng-jek/internal/riders/model"
	"nebeng-jek/pkg/amqp"
	"nebeng-jek/pkg/logger"
)

func (h *ridersHandler) SubscribeDriverAcceptedRides(ctx context.Context, ridesChannel amqp.AMQPChannel) {
	msgs, err := amqp.ConsumeMessageToExchange(ctx, constants.DriverAcceptedRideExchange, ridesChannel)
	if err != nil {
		logger.Fatal(ctx, "error consuming message to exchange", nil)
	}
	for msg := range msgs {
		var data model.MatchedRideMessage
		err := json.Unmarshal(msg.Body, &data)
		if err != nil {
			logger.Error(ctx, "fail to unmarshal consumed message", map[string]interface{}{"error": err})
		}

		broadcastMsg := model.RiderMessage{
			Event: model.EventMatchedRide,
			Data:  data,
		}
		h.broadcastToRider(ctx, data.RiderMSISDN, broadcastMsg)
	}
}

func (h *ridersHandler) SubscribeReadyToPickupRides(ctx context.Context, ridesChannel amqp.AMQPChannel) {
	msgs, err := amqp.ConsumeMessageToExchange(ctx, constants.RideReadyToPickupExchange, ridesChannel)
	if err != nil {
		logger.Fatal(ctx, "error consuming message to exchange", nil)
	}
	for msg := range msgs {
		var data model.RideReadyToPickupMessage
		err := json.Unmarshal(msg.Body, &data)
		if err != nil {
			logger.Error(ctx, "fail to unmarshal consumed message", map[string]interface{}{"error": err})
		}

		broadcastMsg := model.RiderMessage{
			Event: model.EventRideReadyToPickup,
			Data:  data,
		}
		h.broadcastToRider(ctx, data.RiderMSISDN, broadcastMsg)
	}
}
