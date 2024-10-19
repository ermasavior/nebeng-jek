package handler_nats

import (
	"context"
	"encoding/json"
	"nebeng-jek/internal/drivers/model"
	"nebeng-jek/pkg/logger"

	"github.com/nats-io/nats.go"
)

func (h *natsHandler) SubscribeNewRideRequests(ctx context.Context) func(msg *nats.Msg) {
	return func(msg *nats.Msg) {
		var data model.NewRideRequestMessage
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			logger.Error(ctx, "fail to unmarshal consumed message", map[string]interface{}{"error": err})
			msg.Ack()
			return
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
		msg.Ack()
	}
}

func (h *natsHandler) SubscribeReadyToPickupRides(ctx context.Context) func(msg *nats.Msg) {
	return func(msg *nats.Msg) {
		var data model.RideReadyToPickupMessage
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			logger.Error(ctx, "fail to unmarshal consumed message", map[string]interface{}{logger.ErrorKey: err})
			msg.Ack()
			return
		}

		broadcastMsg := model.DriverMessage{
			Event: model.EventRideReadyToPickup,
			Data:  data,
		}
		h.broadcastToActiveDrivers(ctx, map[string]bool{data.DriverMSISDN: true}, broadcastMsg)
		msg.Ack()
	}
}
