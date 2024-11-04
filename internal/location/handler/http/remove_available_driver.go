package handler_http

import (
	pkgErr "nebeng-jek/pkg/error"
	httpUtils "nebeng-jek/pkg/http/utils"
	"nebeng-jek/pkg/logger"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *httpHandler) RemoveAvailableDriver(c *gin.Context) {
	driverIDStr := c.Param("driver_id")
	driverID, _ := strconv.ParseInt(driverIDStr, 10, 64)
	if driverID == 0 {
		c.JSON(
			http.StatusBadRequest,
			httpUtils.NewFailedResponse(http.StatusBadRequest, "invalid driver_id"),
		)
		return
	}

	ctx := c.Request.Context()
	err := h.usecase.RemoveAvailableDriver(ctx, driverID)
	if err != nil {
		logger.Error(ctx, "error from usecase", map[string]interface{}{
			logger.ErrorKey: err.Error(),
		})
		c.JSON(
			http.StatusInternalServerError,
			httpUtils.NewFailedResponse(pkgErr.ErrInternalErrorCode, err.Error()),
		)
		return
	}

	c.JSON(http.StatusOK, httpUtils.NewSuccessResponse(nil))
}
