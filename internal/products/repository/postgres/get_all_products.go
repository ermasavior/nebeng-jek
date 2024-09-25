package repository_postgres

import (
	"context"
	model "nebeng-jek/internal/products/model"
	errorPkg "nebeng-jek/pkg/error"
	"nebeng-jek/pkg/logger"
)

func (r *postgresRepository) GetAllProducts(ctx context.Context) ([]model.Product, error) {
	var products []model.Product

	err := r.db.SelectContext(ctx, &products, queryGetAllProduct)
	if err != nil {
		logger.Error(ctx, err.Error(), nil)
		return nil, errorPkg.NewInternalServerError(err, "error while getting all products")
	}

	return products, nil
}
