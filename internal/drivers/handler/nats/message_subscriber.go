package handler_nats

import (
	"context"
	"encoding/json"
	"nebeng-jek/internal/drivers/model"
	nats_pkg "nebeng-jek/internal/pkg/nats"
	"nebeng-jek/pkg/logger"

	"github.com/nats-io/nats.go"
)

func (h *natsHandler) SubscribeNewRideRequests(ctx context.Context) func(msg *nats.Msg) {
	return func(msg *nats.Msg) {
		var data model.NewRideRequestMessage
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			logger.Error(ctx, "fail to unmarshal consumed message", map[string]interface{}{"error": err})
			nats_pkg.AckMessage(ctx, msg)
			return
		}

		dataMsg, _ := json.Marshal(model.NewRideRequestBroadcast{
			RideID:         data.RideID,
			Rider:          data.Rider,
			PickupLocation: data.PickupLocation,
			Destination:    data.Destination,
		})
		broadcastMsg := model.DriverMessage{
			Event: model.EventNewRideRequest,
			Data:  dataMsg,
		}
		h.broadcastToDrivers(ctx, data.AvailableDrivers, broadcastMsg)
		nats_pkg.AckMessage(ctx, msg)
	}
}

func (h *natsHandler) SubscribeReadyToPickupRides(ctx context.Context) func(msg *nats.Msg) {
	return func(msg *nats.Msg) {
		var data model.RideReadyToPickupMessage
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			logger.Error(ctx, "invalid message", map[string]interface{}{logger.ErrorKey: err})
			nats_pkg.AckMessage(ctx, msg)
			return
		}

		broadcastMsg := model.DriverMessage{
			Event: model.EventRideReadyToPickup,
			Data:  msg.Data,
		}
		h.broadcastToDrivers(ctx, map[int64]bool{data.DriverID: true}, broadcastMsg)
		nats_pkg.AckMessage(ctx, msg)
	}
}
