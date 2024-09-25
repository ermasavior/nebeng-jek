package repository_postgres

import (
	"context"
	"time"

	"nebeng-jek/internal/products/model"
	errorPkg "nebeng-jek/pkg/error"
	"nebeng-jek/pkg/logger"

	"github.com/google/uuid"
)

func (r *postgresRepository) CreateProduct(ctx context.Context, req model.CreateProduct) (string, error) {
	id := uuid.NewString()

	_, err := r.db.ExecContext(ctx, queryCreateProduct, id, req.Name, req.Price, time.Now())
	if err != nil {
		logger.Error(ctx, err.Error(), nil)
		return "", errorPkg.NewInternalServerError(err, "error while creating product")
	}

	return id, nil
}
