package repository

import (
	"context"

	"nebeng-jek/internal/modules/products/model"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, req model.CreateProduct) (id string, err error)
	GetProductByID(ctx context.Context, id string) (model.Product, error)
	GetAllProducts(ctx context.Context) ([]model.Product, error)
	UpdateProductByID(ctx context.Context, req model.UpdateProduct) (model.Product, error)
	DeleteProductByID(ctx context.Context, id string) error
}
