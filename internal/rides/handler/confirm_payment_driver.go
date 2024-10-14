package handler

import (
	"nebeng-jek/internal/rides/model"
	httpUtils "nebeng-jek/pkg/http/utils"
	"nebeng-jek/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *ridesHandler) ConfirmPaymentDriver(c *gin.Context) {
	req := model.ConfirmPaymentDriverRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			httpUtils.NewFailedResponse(http.StatusBadRequest, "error reading request: "+err.Error()),
		)
		return
	}

	ctx := c.Request.Context()
	data, err := h.usecase.ConfirmPaymentDriver(ctx, req)
	if err != nil {
		logger.Error(ctx, "error handler", map[string]interface{}{
			logger.ErrorKey: err.Error(),
		})
		c.JSON(
			err.Code,
			httpUtils.NewFailedResponse(err.Code, err.Message),
		)
		return
	}

	c.JSON(http.StatusOK, httpUtils.NewSuccessResponse(data))
}
