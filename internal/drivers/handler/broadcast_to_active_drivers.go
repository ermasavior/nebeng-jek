package handler

import (
	"context"
	"encoding/json"
	"nebeng-jek/internal/drivers/model"
	"nebeng-jek/pkg/logger"
	pkg_ws "nebeng-jek/pkg/websocket"

	"github.com/gorilla/websocket"
)

func (h *driversHandler) broadcastToActiveDrivers(ctx context.Context, drivers map[string]bool, msg model.DriverMessage) {
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		logger.Error(ctx, "error unmarshalling message broadcast", map[string]interface{}{
			"error": err,
		})
		return
	}

	for driver := range drivers {
		conn, ok := h.connStorage.Load(driver)
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
				"error": err,
			})
			continue
		}
	}
}
