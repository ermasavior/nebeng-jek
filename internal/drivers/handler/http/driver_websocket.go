package handler_http

import (
	"context"
	"encoding/json"
	"nebeng-jek/internal/drivers/model"
	pkg_context "nebeng-jek/internal/pkg/context"
	"nebeng-jek/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func (h *httpHandler) DriverAllocationWebsocket(c *gin.Context) {
	conn, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Error(c.Request.Context(), "error upgrade to websocket", map[string]interface{}{
			"error": err,
		})
		return
	}
	defer conn.Close()

	driverID := pkg_context.GetDriverIDFromContext(c.Request.Context())
	h.connStorage.Store(driverID, conn)
	defer h.connStorage.Delete(driverID)

	ctx := c.Request.Context()

	for {
		var msg model.DriverMessage
		err := conn.ReadJSON(&msg)
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logger.Debug(ctx, "websocket connection closed", map[string]interface{}{
					logger.ErrorKey: err, "driver_id": driverID,
				})
				break
			}

			logger.Error(ctx, "error reading message from driver", map[string]interface{}{
				logger.ErrorKey: err, "driver_id": driverID,
			})
			break
		}

		h.routeMessage(ctx, msg)
	}
}

func (h *httpHandler) routeMessage(ctx context.Context, msg model.DriverMessage) {
	if msg.Event == model.EventRealTimeLocation {
		var req model.TrackUserLocationRequest
		err := json.Unmarshal(msg.Data, &req)
		if err != nil {
			logger.Error(ctx, "error parsing track location request", map[string]interface{}{
				"error": err,
			})
			return
		}

		req.UserID = pkg_context.GetDriverIDFromContext(ctx)
		err = h.usecase.TrackUserLocation(ctx, req)
		if err != nil {
			logger.Error(ctx, "track user location", map[string]interface{}{
				"error": err,
			})
			return
		}
	}
}
