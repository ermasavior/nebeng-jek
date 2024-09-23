package handler

import (
	repository "nebeng-jek/internal/modules/products/repository/postgres"
	usecase "nebeng-jek/internal/modules/products/usecase"
	"nebeng-jek/pkg/redis"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func RegisterProductHandler(router *gin.RouterGroup, db *sqlx.DB, redis redis.Collections) {
	pRepository := repository.NewProductRepository(db)
	pUsecase := usecase.NewProductUsecase(pRepository, redis)
	pHandler := &productHandler{
		usecase: pUsecase,
	}

	router.GET("/:id", pHandler.GetProductByID)
	router.GET("", pHandler.GetAllProducts)
	router.POST("", pHandler.CreateProduct)
	router.PUT("/:id", pHandler.UpdateProductByID)
	router.DELETE("/:id", pHandler.DeleteProductByID)
}
