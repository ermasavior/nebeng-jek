package handler

import (
	"nebeng-jek/internal/drivers/model"
	pkg_context "nebeng-jek/internal/pkg/context"
	"nebeng-jek/pkg/logger"

	"github.com/gin-gonic/gin"
)

func (h *driversHandler) DriverAllocationWebsocket(c *gin.Context) {
	conn, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Error(c.Request.Context(), "error upgrade to websocket", map[string]interface{}{
			"error": err,
		})
		return
	}
	defer conn.Close()

	msisdn := pkg_context.GetMSISDNFromContext(c.Request.Context())
	h.connStorage.Store(msisdn, conn)
	defer func() {
		h.connStorage.Delete(msisdn)
		conn.Close()
	}()

	for {
		var msg model.DriverMessage
		err := conn.ReadJSON(&msg)
		if err != nil {
			logger.Error(c.Request.Context(), "error reading message from driver", map[string]interface{}{
				"error": err,
			})
			continue
		}
	}
}
