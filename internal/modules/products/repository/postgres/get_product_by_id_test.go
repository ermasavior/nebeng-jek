package repository_postgres

import (
	"nebeng-jek/internal/modules/products/model"
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetProductByID(t *testing.T) {
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

	expectedResult := model.Product{
		ID:    "123",
		Name:  "Tea",
		Price: 25000,
	}

	t.Run("should return query result", func(t *testing.T) {
		sqlMock.ExpectQuery(queryGetProductByID).WithArgs(expectedResult.ID).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price"}).
				AddRow("123", "Tea", 25000),
			)

		product, err := repoMock.GetProductByID(context.Background(), expectedResult.ID)
		require.Nil(t, err)
		assert.Equal(t, expectedResult, product)
	})

	t.Run("should return error not found when row is not found", func(t *testing.T) {
		sqlMock.ExpectQuery(queryGetProductByID).
			WillReturnError(sql.ErrNoRows)

		res, err := repoMock.GetProductByID(context.Background(), "")

		assert.Empty(t, res)
		assert.EqualError(t, err, model.ErrorProductNotFound.Error())
	})

	t.Run("should return error when failed get query", func(t *testing.T) {
		rowErr := errors.New("error from db")
		sqlMock.ExpectQuery(queryGetProductByID).
			WillReturnError(rowErr)

		res, err := repoMock.GetProductByID(context.Background(), "")

		assert.Empty(t, res)
		assert.EqualError(t, err, "error while getting product by id")
	})
}
