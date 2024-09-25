package handler

import (
	"net/http"

	"nebeng-jek/internal/products/model"
	httpUtils "nebeng-jek/pkg/http/utils"
	"nebeng-jek/pkg/logger"

	"github.com/gin-gonic/gin"
)

func (p *productHandler) DeleteProductByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(
			http.StatusBadRequest,
			httpUtils.NewFailedResponse(http.StatusBadRequest, "id is not found"),
		)
		return
	}

	ctx := c.Request.Context()
	err := p.usecase.DeleteProductByID(ctx, id)
	if err != nil {
		logger.Error(ctx, err.Error(), nil)
		err = c.Error(err)
		c.JSON(
			http.StatusInternalServerError,
			httpUtils.NewFailedResponse(http.StatusBadRequest, err.Error()),
		)
		return
	}

	res := httpUtils.NewSuccessResponse(model.ProductResponse{
		ID: id,
	})
	c.JSON(http.StatusOK, res)
}
