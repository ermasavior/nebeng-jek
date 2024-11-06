package handler_http

import (
	pkgErr "nebeng-jek/pkg/error"
	httpUtils "nebeng-jek/pkg/http/utils"
	"nebeng-jek/pkg/logger"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *httpHandler) GetRideData(c *gin.Context) {
	rideIDStr := c.Param("ride_id")
	rideID, _ := strconv.ParseInt(rideIDStr, 10, 64)
	if rideID == 0 {
		c.JSON(
			http.StatusBadRequest,
			httpUtils.NewFailedResponse(http.StatusBadRequest, "invalid ride id"),
		)
		return
	}

	ctx := c.Request.Context()
	data, err := h.usecase.GetRideData(ctx, rideID)
	if err != nil {
		logger.Error(ctx, "error from usecase", map[string]interface{}{
			logger.ErrorKey: err.Error(),
		})
		c.JSON(
			pkgErr.ToHttpError(err),
			httpUtils.NewFailedResponse(err.GetCode(), err.GetMessage()),
		)
		return
	}

	response := data.ToResponse()
	c.JSON(http.StatusOK, httpUtils.NewSuccessResponse(response))
}
