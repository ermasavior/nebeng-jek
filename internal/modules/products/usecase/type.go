package usecase

import (
	"context"

	"nebeng-jek/internal/modules/products/model"
	"nebeng-jek/internal/modules/products/repository"
	"nebeng-jek/pkg/redis"
)

type ProductUsecase interface {
	CreateProduct(ctx context.Context, req model.CreateProduct) (string, error)
	GetProductByID(ctx context.Context, id string) (*model.Product, error)
	GetAllProducts(ctx context.Context) ([]model.Product, error)
	UpdateProductByID(ctx context.Context, req model.UpdateProduct) (model.Product, error)
	DeleteProductByID(ctx context.Context, id string) error
}

type productUsecase struct {
	repository repository.ProductRepository
	redis      redis.Collections
}
