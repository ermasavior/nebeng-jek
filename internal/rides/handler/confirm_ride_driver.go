package handler

import (
	"nebeng-jek/internal/rides/model"
	httpUtils "nebeng-jek/pkg/http/utils"
	"nebeng-jek/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *ridesHandler) ConfirmRideDriver(c *gin.Context) {
	req := model.ConfirmRideDriverRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			httpUtils.NewFailedResponse(http.StatusBadRequest, "error reading request: "+err.Error()),
		)
		return
	}

	ctx := c.Request.Context()
	err := h.usecase.ConfirmRideDriver(ctx, req)
	if err != nil && err.Code == http.StatusNotFound {
		c.JSON(
			http.StatusNotFound,
			httpUtils.NewFailedResponse(http.StatusNotFound, err.Error()),
		)
		return
	}
	if err != nil {
		logger.Error(ctx, err.Error(), nil)
		c.JSON(
			http.StatusInternalServerError,
			httpUtils.NewFailedResponse(http.StatusInternalServerError, err.Error()),
		)
		return
	}

	c.JSON(http.StatusOK, httpUtils.NewSuccessResponse(nil))
}
