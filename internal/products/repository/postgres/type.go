package repository_postgres

import "github.com/jmoiron/sqlx"

type postgresRepository struct {
	db *sqlx.DB
}
