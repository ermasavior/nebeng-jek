package handler_nats

import (
	"context"
	"encoding/json"
	nats_pkg "nebeng-jek/internal/pkg/nats"
	"nebeng-jek/internal/riders/model"
	"nebeng-jek/pkg/logger"

	"github.com/nats-io/nats.go"
)

func (h *natsHandler) SubscribeRideMatchedDriver(ctx context.Context) func(*nats.Msg) {
	return func(msg *nats.Msg) {
		var data model.RideMatchedDriverMessage
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			logger.Error(ctx, "invalid message data", map[string]interface{}{logger.ErrorKey: err})
			nats_pkg.AckMessage(ctx, msg)
			return
		}

		broadcastMsg := model.RiderMessage{
			Event: model.EventMatchedRide,
			Data:  msg.Data,
		}
		err = h.broadcastToRider(ctx, data.RiderID, broadcastMsg)
		if err != nil {
			nats_pkg.NakMessage(ctx, msg)
			return
		}
		nats_pkg.AckMessage(ctx, msg)
	}
}

func (h *natsHandler) SubscribeRideStarted(ctx context.Context) func(*nats.Msg) {
	return func(msg *nats.Msg) {
		var data model.RideStartedMessage
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			logger.Error(ctx, "invalid message data", map[string]interface{}{logger.ErrorKey: err})
			nats_pkg.AckMessage(ctx, msg)
			return
		}

		broadcastMsg := model.RiderMessage{
			Event: model.EventRideStarted,
			Data:  msg.Data,
		}
		err = h.broadcastToRider(ctx, data.RiderID, broadcastMsg)
		if err != nil {
			nats_pkg.NakMessage(ctx, msg)
			return
		}
		nats_pkg.AckMessage(ctx, msg)
	}
}

func (h *natsHandler) SubscribeRideEnded(ctx context.Context) func(*nats.Msg) {
	return func(msg *nats.Msg) {
		var data model.RideEndedMessage
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			logger.Error(ctx, "invalid message data", map[string]interface{}{logger.ErrorKey: err})
			nats_pkg.AckMessage(ctx, msg)
			return
		}

		broadcastMsg := model.RiderMessage{
			Event: model.EventRideEnded,
			Data:  msg.Data,
		}
		err = h.broadcastToRider(ctx, data.RiderID, broadcastMsg)
		if err != nil {
			nats_pkg.NakMessage(ctx, msg)
			return
		}
		nats_pkg.AckMessage(ctx, msg)
	}
}

func (h *natsHandler) SubscribeRidePaid(ctx context.Context) func(*nats.Msg) {
	return func(msg *nats.Msg) {
		var data model.RidePaidMessage
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			logger.Error(ctx, "invalid message data", map[string]interface{}{logger.ErrorKey: err})
			nats_pkg.AckMessage(ctx, msg)
			return
		}

		broadcastMsg := model.RiderMessage{
			Event: model.EventRidePaid,
			Data:  msg.Data,
		}
		err = h.broadcastToRider(ctx, data.RiderID, broadcastMsg)
		if err != nil {
			nats_pkg.NakMessage(ctx, msg)
			return
		}
		nats_pkg.AckMessage(ctx, msg)
	}
}
