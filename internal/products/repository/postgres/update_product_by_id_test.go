package repository_postgres

import (
	"context"
	"database/sql"
	"errors"
	"nebeng-jek/internal/products/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateProductByID(t *testing.T) {
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

	req := model.UpdateProduct{
		ID:    "123",
		Name:  "Tea",
		Price: 25000,
	}
	oldRow := model.Product{
		ID:    "123",
		Name:  "Coffee",
		Price: 50000,
	}
	expectedRow := model.Product{
		ID:    "123",
		Name:  "Tea",
		Price: 25000,
	}

	t.Run("should execute update query", func(t *testing.T) {
		sqlMock.ExpectBegin()

		sqlMock.ExpectQuery(queryGetProductByID).
			WithArgs(oldRow.ID).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "name", "price"}).
					AddRow(oldRow.ID, oldRow.Name, oldRow.Price),
			)

		sqlMock.ExpectQuery(queryUpdateProductByID).
			WithArgs(expectedRow.ID, expectedRow.Name, expectedRow.Price).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "name", "price"}).
					AddRow(expectedRow.ID, expectedRow.Name, expectedRow.Price),
			)

		sqlMock.ExpectCommit()

		product, err := repoMock.UpdateProductByID(context.Background(), req)
		require.Nil(t, err)
		assert.Equal(t, expectedRow, product)
	})

	t.Run("should return error when failed update query", func(t *testing.T) {
		dbErr := errors.New("database error")
		sqlMock.ExpectBegin()
		sqlMock.ExpectQuery(queryGetProductByID).
			WithArgs(expectedRow.ID).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "name", "price"}).
					AddRow(oldRow.ID, oldRow.Name, oldRow.Price),
			)
		sqlMock.ExpectQuery(queryUpdateProductByID).WithArgs(expectedRow.ID, expectedRow.Name, expectedRow.Price).
			WillReturnError(dbErr)
		sqlMock.ExpectRollback()

		_, err := repoMock.UpdateProductByID(context.Background(), req)
		assert.EqualError(t, err, dbErr.Error())
	})

	t.Run("should return error when id is not found", func(t *testing.T) {
		noRowsErr := sql.ErrNoRows
		sqlMock.ExpectBegin()
		sqlMock.ExpectQuery(queryGetProductByID).
			WithArgs(expectedRow.ID).
			WillReturnError(noRowsErr)

		sqlMock.ExpectRollback()

		_, err := repoMock.UpdateProductByID(context.Background(), req)
		assert.EqualError(t, err, model.ErrorProductNotFound.Error())
	})

	t.Run("should return error from database", func(t *testing.T) {
		expectedErr := errors.New("error from db")
		sqlMock.ExpectBegin()
		sqlMock.ExpectQuery(queryGetProductByID).
			WithArgs(expectedRow.ID).
			WillReturnError(expectedErr)

		sqlMock.ExpectRollback()

		_, err := repoMock.UpdateProductByID(context.Background(), req)
		assert.EqualError(t, err, expectedErr.Error())
	})
}
