package handler_nats

import (
	"context"
	"encoding/json"
	"errors"
	"nebeng-jek/internal/riders/model"
	"nebeng-jek/pkg/logger"
	pkg_ws "nebeng-jek/pkg/websocket"

	"github.com/gorilla/websocket"
)

func (h *natsHandler) broadcastToRider(ctx context.Context, riderID int64, msg model.RiderMessage) error {
	conn, ok := h.connStorage.Load(riderID)
	if !ok {
		logger.Error(ctx, "rider id not found", map[string]interface{}{
			"rider_id": riderID,
		})
		return errors.New("rider id is not found")
	}

	wsConn, ok := conn.(pkg_ws.WebsocketInterface)
	if !ok {
		logger.Error(ctx, "error loading rider connection websocket", map[string]interface{}{
			"rider_id": riderID,
		})
		return errors.New("rider connection is invalid")
	}

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		logger.Error(ctx, "error unmarshalling message broadcast", map[string]interface{}{
			"error": err,
		})
		return nil
	}

	if err := wsConn.WriteMessage(websocket.TextMessage, msgBytes); err != nil {
		logger.Error(ctx, "error broadcasting to riders via websocket", map[string]interface{}{
			"error":    err,
			"rider_id": riderID,
		})
		return err
	}

	return nil
}
