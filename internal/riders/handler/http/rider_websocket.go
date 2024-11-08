package handler_http

import (
	"context"
	"encoding/json"
	pkg_context "nebeng-jek/internal/pkg/context"
	"nebeng-jek/internal/pkg/location"
	"nebeng-jek/internal/riders/model"
	"nebeng-jek/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func (h *httpHandler) RiderWebsocket(c *gin.Context) {
	conn, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Error(c.Request.Context(), "error upgrade to websocket", map[string]interface{}{
			"error": err,
		})
		return
	}
	defer conn.Close()

	riderID := pkg_context.GetRiderIDFromContext(c.Request.Context())
	h.connStorage.Store(riderID, conn)
	defer func() {
		h.connStorage.Delete(riderID)
		conn.Close()
	}()

	ctx := c.Request.Context()

	for {
		var msg model.RiderMessage
		err := conn.ReadJSON(&msg)
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logger.Debug(ctx, "websocket connection closed", map[string]interface{}{
					logger.ErrorKey: err, "rider_id": riderID,
				})
				break
			}

			logger.Error(ctx, "error reading message from rider", map[string]interface{}{
				logger.ErrorKey: err, "rider_id": riderID,
			})
			break
		}

		h.routeMessage(ctx, msg)
	}
}

func (h *httpHandler) routeMessage(ctx context.Context, msg model.RiderMessage) {
	if msg.Event == location.EventRealTimeLocation {
		var req model.TrackUserLocationRequest
		err := json.Unmarshal(msg.Data, &req)
		if err != nil {
			logger.Error(ctx, "error parsing track location request", map[string]interface{}{
				"error": err,
			})
			return
		}

		err = h.usecase.TrackUserLocation(ctx, req)
		if err != nil {
			logger.Error(ctx, "track user location", map[string]interface{}{
				"error": err,
			})
			return
		}
	}
}
