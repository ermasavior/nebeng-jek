package usecase

import (
	"context"

	"nebeng-jek/internal/products/model"
)

func (usecase *productUsecase) GetAllProducts(ctx context.Context) ([]model.Product, error) {
	products, err := usecase.repository.GetAllProducts(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}
