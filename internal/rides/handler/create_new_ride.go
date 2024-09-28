package handler

import (
	"nebeng-jek/internal/rides/model"
	httpUtils "nebeng-jek/pkg/http/utils"
	"nebeng-jek/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *ridesHandler) CreateNewRide(c *gin.Context) {
	req := model.CreateNewRideRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		err = c.Error(err)
		c.JSON(
			http.StatusBadRequest,
			httpUtils.NewFailedResponse(http.StatusBadRequest, "error reading request: "+err.Error()),
		)
		return
	}

	ctx := c.Request.Context()
	id, err := h.usecase.CreateNewRide(ctx, req)
	if err != nil {
		logger.Error(ctx, err.Error(), nil)
		err = c.Error(err)
		c.JSON(
			http.StatusInternalServerError,
			httpUtils.NewFailedResponse(http.StatusInternalServerError, err.Error()),
		)
		return
	}

	c.JSON(http.StatusOK, httpUtils.NewSuccessResponse(id))
}
