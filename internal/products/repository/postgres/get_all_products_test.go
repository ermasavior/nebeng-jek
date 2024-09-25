package repository_postgres

import (
	"context"
	"errors"
	"nebeng-jek/internal/products/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAllProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic("failed mocking sql")
	}
	defer func() {
		_ = db.Close()
	}()

	repoMock := NewProductRepository(sqlx.NewDb(db, "sqlmock"))

	expectedResults := []model.Product{
		{
			ID:    "123",
			Name:  "abc",
			Price: 8000,
		},
	}

	t.Run("should return query result", func(t *testing.T) {
		sqlMock.ExpectQuery(queryGetAllProduct).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price"}).
				AddRow("123", "abc", 8000),
			)

		products, err := repoMock.GetAllProducts(context.Background())
		require.Nil(t, err)
		assert.Equal(t, expectedResults, products)
	})

	t.Run("should return error when failed select query", func(t *testing.T) {
		rowErr := errors.New("error from db")
		sqlMock.ExpectQuery(queryGetAllProduct).
			WillReturnError(rowErr)

		res, err := repoMock.GetAllProducts(context.Background())

		assert.Equal(t, []model.Product(nil), res)
		assert.EqualError(t, err, "error while getting all products")
	})
}
