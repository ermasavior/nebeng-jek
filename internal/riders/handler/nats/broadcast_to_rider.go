package handler_nats

import (
	"context"
	"encoding/json"
	"nebeng-jek/internal/riders/model"
	"nebeng-jek/pkg/logger"
	pkg_ws "nebeng-jek/pkg/websocket"

	"github.com/gorilla/websocket"
)

func (h *natsHandler) broadcastToRider(ctx context.Context, riderID int64, msg model.RiderMessage) error {
	conn, ok := h.connStorage.Load(riderID)
	if !ok {
		return nil
	}

	wsConn, ok := conn.(pkg_ws.WebsocketInterface)
	if !ok {
		logger.Error(ctx, "error loading rider connection websocket", nil)
		return nil
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
