package handler

import (
	"nebeng-jek/internal/products/model"
	httpUtils "nebeng-jek/pkg/http/utils"
	"nebeng-jek/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p *productHandler) UpdateProductByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(
			http.StatusBadRequest,
			httpUtils.NewFailedResponse(http.StatusBadRequest, "id is not found"),
		)
		return
	}

	req := model.UpdateProduct{
		ID: id,
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		err = c.Error(err)
		c.JSON(
			http.StatusBadRequest,
			httpUtils.NewFailedResponse(http.StatusBadRequest, err.Error()),
		)
		return
	}

	ctx := c.Request.Context()
	product, err := p.usecase.UpdateProductByID(ctx, req)

	if err == model.ErrorProductNotFound {
		c.JSON(
			http.StatusNotFound,
			httpUtils.NewFailedResponse(http.StatusNotFound, err.Error()),
		)
		return
	}

	if err != nil {
		logger.Error(ctx, err.Error(), nil)
		err = c.Error(err)
		c.JSON(
			http.StatusInternalServerError,
			httpUtils.NewFailedResponse(http.StatusInternalServerError, err.Error()),
		)
		return
	}

	res := httpUtils.NewSuccessResponse(model.UpdateProductResponse(product))
	c.JSON(http.StatusOK, res)
}
