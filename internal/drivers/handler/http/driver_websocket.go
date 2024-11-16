package handler_http

import (
	"context"
	"encoding/json"
	"io"
	"nebeng-jek/internal/drivers/model"
	pkg_context "nebeng-jek/internal/pkg/context"
	"nebeng-jek/internal/pkg/location"
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
			if _, ok := err.(*json.SyntaxError); ok {
				logger.Info(ctx, "invalid json message from driver", map[string]interface{}{
					logger.ErrorKey: err, "driver_id": driverID,
				})
				continue
			} else if err == io.ErrUnexpectedEOF {
				continue
			} else if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				break
			} else if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logger.Error(ctx, "error unexpected closed connection", map[string]interface{}{
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
