package repository_postgres

import (
	errorPkg "nebeng-jek/pkg/error"
	"nebeng-jek/pkg/logger"
	"context"
)

func (r *postgresRepository) DeleteProductByID(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, queryDeleteProductByID, id)
	if err != nil {
		logger.Error(ctx, err.Error(), nil)
		return errorPkg.NewInternalServerError(err, "error while deleting product by id")
	}

	return nil
}
