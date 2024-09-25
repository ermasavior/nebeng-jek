package repository_postgres

import (
	"context"
	"database/sql"
	model "nebeng-jek/internal/products/model"
	"nebeng-jek/pkg/logger"
)

func (r *postgresRepository) UpdateProductByID(ctx context.Context, req model.UpdateProduct) (model.Product, error) {
	// transaction example
	tx, err := r.db.BeginTxx(ctx, &sql.TxOptions{
		Isolation: 0,
		ReadOnly:  false,
	})
	if err != nil {
		logger.Error(ctx, err.Error(), nil)
		return model.Product{}, err
	}

	defer func() {
		if err := tx.Commit(); err != nil && err != sql.ErrTxDone {
			logger.Error(ctx, err.Error(), nil)
		}
	}()

	var product model.Product

	err = tx.GetContext(ctx, &product, queryGetProductByID, req.ID)
	if err == sql.ErrNoRows {
		if err := tx.Rollback(); err != nil {
			logger.Error(ctx, err.Error(), nil)
		}
		return model.Product{}, model.ErrorProductNotFound
	}

	if err != nil {
		if err := tx.Rollback(); err != nil {
			logger.Error(ctx, err.Error(), nil)
		}
		logger.Error(ctx, err.Error(), nil)
		return model.Product{}, err
	}

	err = tx.QueryRowxContext(ctx, queryUpdateProductByID, req.ID, req.Name, req.Price).StructScan(&product)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			logger.Error(ctx, err.Error(), nil)
		}
		logger.Error(ctx, err.Error(), nil)
		return model.Product{}, err
	}

	return product, nil
}
