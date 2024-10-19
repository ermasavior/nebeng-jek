package handler_nats

import (
	"context"
	"encoding/json"
	"nebeng-jek/internal/drivers/model"
	"nebeng-jek/pkg/logger"
	pkg_ws "nebeng-jek/pkg/websocket"

	"github.com/gorilla/websocket"
)

func (h *natsHandler) broadcastToDrivers(ctx context.Context, mapDriverID map[int64]bool, msg model.DriverMessage) {
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		logger.Error(ctx, "error unmarshalling message broadcast", map[string]interface{}{
			"error": err,
		})
		return
	}

	for id := range mapDriverID {
		conn, ok := h.connStorage.Load(id)
		if !ok {
			continue
		}

		wsConn, ok := conn.(pkg_ws.WebsocketInterface)
		if !ok {
			logger.Error(ctx, "error loading driver connection websocket", nil)
			continue
		}

		if err := wsConn.WriteMessage(websocket.TextMessage, msgBytes); err != nil {
			logger.Error(ctx, "error broadcasting to drivers via websocket", map[string]interface{}{
				"error":      err,
				"driver_ids": mapDriverID,
			})
			continue
		}
	}
}
