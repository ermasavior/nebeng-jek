package handler_http

import (
	"nebeng-jek/internal/location/model"
	pkgErr "nebeng-jek/pkg/error"
	httpUtils "nebeng-jek/pkg/http/utils"
	"nebeng-jek/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *httpHandler) GetRidePath(c *gin.Context) {
	req := model.GetRidePathRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			httpUtils.NewFailedResponse(http.StatusBadRequest, "error reading request: "+err.Error()),
		)
		return
	}

	ctx := c.Request.Context()
	ridePath, err := h.usecase.GetRidePath(ctx, req.RideID, req.DriverID)
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

	res := model.GetRidePathResponse{
		Path: ridePath,
	}
	c.JSON(http.StatusOK, httpUtils.NewSuccessResponse(res))
}
