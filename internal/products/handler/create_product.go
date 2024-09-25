package handler

import (
	"nebeng-jek/internal/products/model"
	httpUtils "nebeng-jek/pkg/http/utils"
	"nebeng-jek/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p *productHandler) CreateProduct(c *gin.Context) {
	var req model.CreateProduct
	if err := c.ShouldBindJSON(&req); err != nil {
		err = c.Error(err)
		c.JSON(
			http.StatusBadRequest,
			httpUtils.NewFailedResponse(http.StatusBadRequest, err.Error()),
		)
		return
	}

	ctx := c.Request.Context()
	id, err := p.usecase.CreateProduct(ctx, req)
	if err != nil {
		logger.Error(ctx, err.Error(), nil)
		err = c.Error(err)
		c.JSON(
			http.StatusInternalServerError,
			httpUtils.NewFailedResponse(http.StatusInternalServerError, err.Error()),
		)
		return
	}

	res := httpUtils.NewSuccessResponse(model.ProductResponse{
		ID: id,
	})
	c.JSON(http.StatusCreated, res)
}
