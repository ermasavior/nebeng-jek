package handler

import (
	"context"
	"encoding/json"
	"nebeng-jek/internal/drivers/model"
	"nebeng-jek/internal/pkg/constants"
	"nebeng-jek/pkg/amqp"
	"nebeng-jek/pkg/logger"
)

func (h *driversHandler) SubscribeNewRideRequests(ctx context.Context, amqpConn amqp.AMQPConnection) {
	channel, err := amqpConn.Channel()
	if err != nil {
		logger.Fatal(context.Background(), "error initializing amqp channel", map[string]interface{}{logger.ErrorKey: err})
	}
	defer channel.Close()

	msgs, err := amqp.ConsumeMessageToExchange(ctx, constants.NewRideRequestsExchange, channel)
	if err != nil {
		logger.Fatal(ctx, "error consuming message to exchange", nil)
	}
	for msg := range msgs {
		var data model.NewRideRequestMessage
		err := json.Unmarshal(msg.Body, &data)
		if err != nil {
			logger.Error(ctx, "fail to unmarshal consumed message", map[string]interface{}{"error": err})
		}

		broadcastMsg := model.DriverMessage{
			Event: model.EventNewRideRequest,
			Data: model.NewRideRequestBroadcast{
				RideID:         data.RideID,
				Rider:          data.Rider,
				PickupLocation: data.PickupLocation,
				Destination:    data.Destination,
			},
		}
		h.broadcastToActiveDrivers(ctx, data.AvailableDrivers, broadcastMsg)
	}
}

func (h *driversHandler) SubscribeReadyToPickupRides(ctx context.Context, amqpConn amqp.AMQPConnection) {
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

		broadcastMsg := model.DriverMessage{
			Event: model.EventRideReadyToPickup,
			Data:  data,
		}
		h.broadcastToActiveDrivers(ctx, map[string]bool{data.DriverMSISDN: true}, broadcastMsg)
	}
}

func (h *driversHandler) SubscribeRideStarted(ctx context.Context, amqpConn amqp.AMQPConnection) {
	ridesChannel, err := amqpConn.Channel()
	if err != nil {
		logger.Fatal(context.Background(), "error initializing amqp channel", map[string]interface{}{logger.ErrorKey: err})
	}
	defer ridesChannel.Close()

	msgs, err := amqp.ConsumeMessageToExchange(ctx, constants.RideStartedExchange, ridesChannel)
	if err != nil {
		logger.Fatal(ctx, "error consuming message to exchange", nil)
	}
	for msg := range msgs {
		var data model.RideReadyToPickupMessage
		err := json.Unmarshal(msg.Body, &data)
		if err != nil {
			logger.Error(ctx, "fail to unmarshal consumed message", map[string]interface{}{"error": err})
		}

		broadcastMsg := model.DriverMessage{
			Event: model.EventRideStarted,
			Data:  data,
		}
		h.broadcastToActiveDrivers(ctx, map[string]bool{data.DriverMSISDN: true}, broadcastMsg)
	}
}
