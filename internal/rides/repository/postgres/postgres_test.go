package repository_postgres

import (
	"context"
	"errors"
	"nebeng-jek/internal/rides/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestRepository_GetRiderIDByMSISDN(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic("failed mocking sql")
	}
	defer func() {
		_ = db.Close()
	}()
	repoMock := NewRepository(sqlx.NewDb(db, "sqlmock"))

	ctx := context.Background()
	msisdn := "08111111"
	expectedID := int64(999)

	expectedQuery := queryGetRiderByMSISDN

	t.Run("should execute get query", func(t *testing.T) {
		sqlMock.ExpectQuery(expectedQuery).
			WithArgs(msisdn).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedID))

		id, err := repoMock.GetRiderIDByMSISDN(ctx, msisdn)

		assert.Equal(t, expectedID, id)
		assert.Nil(t, err)
	})

	t.Run("should return error when error from db", func(t *testing.T) {
		rowErr := errors.New("error from db")
		sqlMock.ExpectQuery(expectedQuery).
			WithArgs(msisdn).
			WillReturnError(rowErr)

		id, err := repoMock.GetRiderIDByMSISDN(ctx, msisdn)

		assert.Equal(t, int64(0), id)
		assert.NotNil(t, err)
	})
}

func TestRepository_CreateNewRide(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic("failed mocking sql")
	}
	defer func() {
		_ = db.Close()
	}()

	repoMock := NewRepository(sqlx.NewDb(db, "sqlmock"))

	ctx := context.Background()
	expectedID := int64(8888)
	req := model.CreateNewRideRequest{
		RiderID: 999,
		PickupLocation: model.Coordinate{
			Latitude:  1,
			Longitude: 2,
		},
		Destination: model.Coordinate{
			Latitude:  9,
			Longitude: 10,
		},
	}

	expectedQuery := queryInsertNewRide

	t.Run("should execute insert query", func(t *testing.T) {
		sqlMock.ExpectQuery(expectedQuery).
			WithArgs(
				req.RiderID, model.StatusNumRideWaitingForDriver,
				req.PickupLocation.Latitude, req.PickupLocation.Longitude,
				req.Destination.Latitude, req.Destination.Longitude).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedID))

		id, err := repoMock.CreateNewRide(ctx, req)

		assert.Equal(t, expectedID, id)
		assert.Nil(t, err)
	})

	t.Run("should return error when error from db", func(t *testing.T) {
		rowErr := errors.New("error from db")
		sqlMock.ExpectQuery(expectedQuery).
			WithArgs(
				req.RiderID, model.StatusNumRideWaitingForDriver,
				req.PickupLocation.Latitude, req.PickupLocation.Longitude,
				req.Destination.Latitude, req.Destination.Longitude).
			WillReturnError(rowErr)

		id, err := repoMock.CreateNewRide(ctx, req)

		assert.Equal(t, int64(0), id)
		assert.NotNil(t, err)
	})
}
