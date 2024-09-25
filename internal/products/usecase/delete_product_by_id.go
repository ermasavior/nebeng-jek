package usecase

import (
	"context"
)

func (usecase *productUsecase) DeleteProductByID(ctx context.Context, id string) error {
	product, err := usecase.repository.GetProductByID(ctx, id)
	if err != nil {
		return err
	}

	err = usecase.repository.DeleteProductByID(ctx, product.ID)
	if err != nil {
		return err
	}

	return nil
}
