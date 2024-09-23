package handler

import (
	"nebeng-jek/internal/modules/products/model"
	httpUtils "nebeng-jek/pkg/http/utils"
	"nebeng-jek/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p *productHandler) GetAllProducts(c *gin.Context) {
	ctx := c.Request.Context()
	products, err := p.usecase.GetAllProducts(ctx)
	if err != nil {
		logger.Error(ctx, err.Error(), nil)
		err = c.Error(err)
		c.JSON(
			http.StatusInternalServerError,
			httpUtils.NewFailedResponse(http.StatusInternalServerError, err.Error()),
		)
		return
	}

	var productsResponse []model.Product
	for _, p := range products {
		productsResponse = append(productsResponse, model.Product{
			ID:    p.ID,
			Name:  p.Name,
			Price: p.Price,
		})
	}

	res := httpUtils.NewSuccessResponse(productsResponse)
	c.JSON(http.StatusOK, res)
}
