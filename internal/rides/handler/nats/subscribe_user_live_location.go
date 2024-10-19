package handler_nats

import (
	"context"
	"encoding/json"
	"nebeng-jek/internal/rides/model"
	"nebeng-jek/pkg/logger"

	"github.com/nats-io/nats.go"
)

func (h *natsHandler) SubscribeUserLiveLocation(ctx context.Context) func(*nats.Msg) {
	return func(msg *nats.Msg) {
		var reqBody model.TrackUserLocationRequest
		err := json.Unmarshal(msg.Data, &reqBody)
		if err != nil {
			logger.Error(ctx, "fail to unmarshal consumed message", map[string]interface{}{logger.ErrorKey: err})
			msg.Ack()
			return
		}

		pkgErr := h.usecase.TrackUserLocation(ctx, reqBody)
		if pkgErr != nil {
			logger.Error(ctx, "error tracking user location", map[string]interface{}{
				"req": reqBody,
			})
			msg.Nak()
			return
		}

		msg.Ack()
	}
}
