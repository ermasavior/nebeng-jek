package repository_postgres

import (
	"nebeng-jek/internal/products/repository"

	"github.com/jmoiron/sqlx"
)

func NewProductRepository(db *sqlx.DB) repository.ProductRepository {
	return &postgresRepository{
		db: db,
	}
}
