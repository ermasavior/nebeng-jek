package handler

import (
	"context"
	"encoding/json"
	"nebeng-jek/internal/drivers/model"
	"nebeng-jek/internal/pkg/constants"
	"nebeng-jek/pkg/amqp"
	"nebeng-jek/pkg/logger"
)

func (h *driversHandler) SubscribeNewRideRequests(ctx context.Context, ridesChannel amqp.AMQPChannel) {
	msgs, err := amqp.ConsumeMessageToExchange(ctx, constants.NewRideRequestsExchange, ridesChannel)
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

func (h *driversHandler) SubscribeReadyToPickupRides(ctx context.Context, ridesChannel amqp.AMQPChannel) {
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

		broadcastMsg := model.DriverMessage{
			Event: model.EventRideReadyToPickup,
			Data:  data,
		}
		h.broadcastToActiveDrivers(ctx, map[string]bool{data.DriverMSISDN: true}, broadcastMsg)
	}
}
