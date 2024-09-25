package usecase

import (
	"context"

	"nebeng-jek/internal/products/model"
)

func (usecase *productUsecase) CreateProduct(ctx context.Context, req model.CreateProduct) (string, error) {
	id, err := usecase.repository.CreateProduct(ctx, req)
	if err != nil {
		return "", err
	}

	return id, nil
}
