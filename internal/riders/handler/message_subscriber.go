package handler

import (
	"context"
	"encoding/json"
	"nebeng-jek/internal/pkg/constants"
	"nebeng-jek/internal/riders/model"
	"nebeng-jek/pkg/amqp"
	"nebeng-jek/pkg/logger"
)

func (h *ridersHandler) SubscribeDriverAcceptedRides(ctx context.Context, amqpConn amqp.AMQPConnection) {
	channel, err := amqpConn.Channel()
	if err != nil {
		logger.Fatal(context.Background(), "error initializing amqp channel", map[string]interface{}{logger.ErrorKey: err})
	}
	defer channel.Close()

	msgs, err := amqp.ConsumeMessageToExchange(ctx, constants.DriverAcceptedRideExchange, channel)
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

func (h *ridersHandler) SubscribeReadyToPickupRides(ctx context.Context, amqpConn amqp.AMQPConnection) {
	channel, err := amqpConn.Channel()
	if err != nil {
		logger.Fatal(context.Background(), "error initializing amqp channel", map[string]interface{}{logger.ErrorKey: err})
	}
	defer channel.Close()

	msgs, err := amqp.ConsumeMessageToExchange(ctx, constants.RideReadyToPickupExchange, channel)
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

func (h *ridersHandler) SubscribeRideStarted(ctx context.Context, amqpConn amqp.AMQPConnection) {
	channel, err := amqpConn.Channel()
	if err != nil {
		logger.Fatal(context.Background(), "error initializing amqp channel", map[string]interface{}{logger.ErrorKey: err})
	}
	defer channel.Close()

	msgs, err := amqp.ConsumeMessageToExchange(ctx, constants.RideStartedExchange, channel)
	if err != nil {
		logger.Fatal(ctx, "error consuming message to exchange", nil)
	}
	for msg := range msgs {
		var data model.RideStartedMessage
		err := json.Unmarshal(msg.Body, &data)
		if err != nil {
			logger.Error(ctx, "fail to unmarshal consumed message", map[string]interface{}{"error": err})
		}

		broadcastMsg := model.RiderMessage{
			Event: model.EventRideStarted,
			Data:  data,
		}
		h.broadcastToRider(ctx, data.RiderMSISDN, broadcastMsg)
	}
}

func (h *ridersHandler) SubscribeRideEnded(ctx context.Context, amqpConn amqp.AMQPConnection) {
	channel, err := amqpConn.Channel()
	if err != nil {
		logger.Fatal(context.Background(), "error initializing amqp channel", map[string]interface{}{logger.ErrorKey: err})
	}
	defer channel.Close()

	msgs, err := amqp.ConsumeMessageToExchange(ctx, constants.RideEndedExchange, channel)
	if err != nil {
		logger.Fatal(ctx, "error consuming message to exchange", nil)
	}
	for msg := range msgs {
		var data model.RideEndedMessage
		err := json.Unmarshal(msg.Body, &data)
		if err != nil {
			logger.Error(ctx, "fail to unmarshal consumed message", map[string]interface{}{"error": err})
		}

		broadcastMsg := model.RiderMessage{
			Event: model.EventRideEnded,
			Data:  data,
		}
		h.broadcastToRider(ctx, data.RiderMSISDN, broadcastMsg)
	}
}
