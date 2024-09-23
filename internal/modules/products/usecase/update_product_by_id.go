package usecase

import (
	"context"

	"nebeng-jek/internal/modules/products/model"
)

func (usecase *productUsecase) UpdateProductByID(ctx context.Context, req model.UpdateProduct) (model.Product, error) {
	product, err := usecase.repository.UpdateProductByID(ctx, req)
	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}
