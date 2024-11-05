package handler_nats

import (
	"context"
	"encoding/json"
	"errors"
	"nebeng-jek/internal/drivers/model"
	"nebeng-jek/pkg/logger"
	pkg_ws "nebeng-jek/pkg/websocket"

	"github.com/gorilla/websocket"
)

func (h *natsHandler) broadcastToDriver(ctx context.Context, driverID int64, msg model.DriverMessage) error {
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		logger.Error(ctx, "error unmarshalling message broadcast", map[string]interface{}{
			"error": err,
		})
		return nil
	}

	conn, ok := h.connStorage.Load(driverID)
	if !ok {
		return errors.New("driver id is not found")
	}

	wsConn, ok := conn.(pkg_ws.WebsocketInterface)
	if !ok {
		logger.Error(ctx, "error loading driver connection websocket", nil)
		return errors.New("error loading driver connection")
	}

	if err := wsConn.WriteMessage(websocket.TextMessage, msgBytes); err != nil {
		logger.Error(ctx, "error broadcasting to drivers via websocket", map[string]interface{}{
			"error":     err,
			"driver_id": driverID,
		})
		return err
	}

	return nil
}
