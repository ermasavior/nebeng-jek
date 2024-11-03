package handler_nats

import (
	"context"
	"encoding/json"
	"nebeng-jek/internal/location/model"
	nats_pkg "nebeng-jek/internal/pkg/nats"
	"nebeng-jek/pkg/logger"

	"github.com/nats-io/nats.go"
)

func (h *natsHandler) SubscribeUserLiveLocation(ctx context.Context) func(*nats.Msg) {
	return func(msg *nats.Msg) {
		var reqBody model.TrackUserLocationRequest
		err := json.Unmarshal(msg.Data, &reqBody)
		if err != nil {
			logger.Error(ctx, "fail to unmarshal consumed message", map[string]interface{}{logger.ErrorKey: err})
			nats_pkg.AckMessage(ctx, msg)
			return
		}

		pkgErr := h.usecase.TrackUserLocation(ctx, reqBody)
		if pkgErr != nil {
			logger.Error(ctx, "error tracking user location", map[string]interface{}{
				"req": reqBody,
			})
			nats_pkg.NakMessage(ctx, msg)
			return
		}

		nats_pkg.AckMessage(ctx, msg)
	}
}
