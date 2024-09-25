package repository_postgres

import (
	"context"
	"database/sql"
	model "nebeng-jek/internal/products/model"
	errorPkg "nebeng-jek/pkg/error"
	"nebeng-jek/pkg/logger"
)

func (r *postgresRepository) GetProductByID(ctx context.Context, id string) (model.Product, error) {
	var product model.Product

	err := r.db.GetContext(ctx, &product, queryGetProductByID, id)
	if err == sql.ErrNoRows {
		return model.Product{}, model.ErrorProductNotFound
	}
	if err != nil {
		logger.Error(ctx, err.Error(), nil)
		return model.Product{}, errorPkg.NewInternalServerError(err, "error while getting product by id")
	}

	return product, nil
}
