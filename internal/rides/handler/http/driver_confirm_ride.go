package handler_http

import (
	"nebeng-jek/internal/rides/model"
	pkgErr "nebeng-jek/pkg/error"
	httpUtils "nebeng-jek/pkg/http/utils"
	"nebeng-jek/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *httpHandler) DriverConfirmRide(c *gin.Context) {
	req := model.DriverConfirmRideRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			httpUtils.NewFailedResponse(http.StatusBadRequest, "error reading request: "+err.Error()),
		)
		return
	}

	ctx := c.Request.Context()
	err := h.usecase.DriverConfirmRide(ctx, req)
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

	c.JSON(http.StatusOK, httpUtils.NewSuccessResponse(nil))
}
