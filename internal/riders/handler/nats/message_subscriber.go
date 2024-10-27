package handler_nats

import (
	"context"
	"encoding/json"
	"nebeng-jek/internal/riders/model"
	"nebeng-jek/pkg/logger"

	"github.com/nats-io/nats.go"
)

func (h *natsHandler) SubscribeRideMatchedDriver(ctx context.Context) func(*nats.Msg) {
	return func(msg *nats.Msg) {
		var data model.RideMatchedDriverMessage
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			logger.Error(ctx, "fail to unmarshal consumed message", map[string]interface{}{logger.ErrorKey: err})
			msg.Ack()
			return
		}

		broadcastMsg := model.RiderMessage{
			Event: model.EventMatchedRide,
			Data:  data,
		}
		err = h.broadcastToRider(ctx, data.RiderID, broadcastMsg)
		if err != nil {
			msg.Nak()
			return
		}
		msg.Ack()
	}
}

func (h *natsHandler) SubscribeRideStarted(ctx context.Context) func(*nats.Msg) {
	return func(msg *nats.Msg) {
		var data model.RideStartedMessage
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			logger.Error(ctx, "fail to unmarshal consumed message", map[string]interface{}{logger.ErrorKey: err})
			msg.Ack()
			return
		}

		broadcastMsg := model.RiderMessage{
			Event: model.EventRideStarted,
			Data:  data,
		}
		err = h.broadcastToRider(ctx, data.RiderID, broadcastMsg)
		if err != nil {
			msg.Nak()
			return
		}
		msg.Ack()
	}
}

func (h *natsHandler) SubscribeRideEnded(ctx context.Context) func(*nats.Msg) {
	return func(msg *nats.Msg) {
		var data model.RideEndedMessage
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			logger.Error(ctx, "fail to unmarshal consumed message", map[string]interface{}{logger.ErrorKey: err})
			msg.Ack()
			return
		}

		broadcastMsg := model.RiderMessage{
			Event: model.EventRideEnded,
			Data:  data,
		}
		err = h.broadcastToRider(ctx, data.RiderID, broadcastMsg)
		if err != nil {
			msg.Nak()
			return
		}
		msg.Ack()
	}
}

func (h *natsHandler) SubscribeRidePaid(ctx context.Context) func(*nats.Msg) {
	return func(msg *nats.Msg) {
		var data model.RidePaidMessage
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			logger.Error(ctx, "fail to unmarshal consumed message", map[string]interface{}{logger.ErrorKey: err})
			msg.Ack()
			return
		}

		broadcastMsg := model.RiderMessage{
			Event: model.EventRidePaid,
			Data:  data,
		}
		err = h.broadcastToRider(ctx, data.RiderID, broadcastMsg)
		if err != nil {
			msg.Nak()
			return
		}
		msg.Ack()
	}
}
