package repository_postgres

import (
	"context"
	"errors"
	"testing"

	"nebeng-jek/internal/modules/products/model"
	"nebeng-jek/pkg/utils"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestRepository_CreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic("failed mocking sql")
	}
	defer func() {
		_ = db.Close()
	}()

	repoMock := postgresRepository{
		db: sqlx.NewDb(db, "sqlmock"),
	}

	ctx := context.Background()
	req := model.CreateProduct{
		Name:  "Laptop Lenovo XYZ",
		Price: 30000000,
	}

	expectedQuery := queryCreateProduct

	t.Run("should execute insert query", func(t *testing.T) {
		sqlMock.ExpectExec(expectedQuery).
			WithArgs(sqlmock.AnyArg(), req.Name, req.Price, utils.AnyTime{}).
			WillReturnResult(sqlmock.NewResult(0, 1))

		id, err := repoMock.CreateProduct(ctx, req)

		assert.NotEqual(t, "", id)
		assert.Nil(t, err)
	})

	t.Run("should return error when failed insert query", func(t *testing.T) {
		rowErr := errors.New("error from db")
		sqlMock.ExpectExec(expectedQuery).
			WillReturnError(rowErr)

		id, err := repoMock.CreateProduct(ctx, req)

		assert.Equal(t, "", id)
		assert.EqualError(t, err, "error while creating product")
	})
}
