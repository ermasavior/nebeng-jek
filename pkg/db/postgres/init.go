package db

import (
	"fmt"
	"nebeng-jek/pkg/configs"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	"github.com/uptrace/opentelemetry-go-extra/otelsqlx"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type PostgresDsn struct {
	Host     string
	Port     int
	User     string
	Password string
	Db       string
	Env      string
}

func (p PostgresDsn) ToString() string {
	sslMode := "disable"
	if p.Env == configs.EnvProduction {
		sslMode = "require"
	}

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", p.Host, p.User, p.Password, p.Db, p.Port, sslMode)
}

func NewPostgresDB(pgDsn PostgresDsn) (*sqlx.DB, error) {
	db, err := otelsqlx.Open(
		"postgres",
		pgDsn.ToString(),
		otelsql.WithAttributes(semconv.DBSystemPostgreSQL),
	)
	if err != nil {
		return db, err
	}

	return db, nil
}
