package handler_http

import (
	pkg_context "nebeng-jek/internal/pkg/context"
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
	}
}
