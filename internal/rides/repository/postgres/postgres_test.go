package repository_postgres

import (
	"context"
	"database/sql"
	"errors"
	"nebeng-jek/internal/pkg/constants"
	"nebeng-jek/internal/rides/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestRepository_GetRiderDataByMSISDN(t *testing.T) {
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
	expectedData := model.RiderData{
		ID:     int64(999),
		Name:   "Melati",
		MSISDN: "0822222",
	}

	expectedQuery := queryGetRiderByMSISDN

	t.Run("should execute get query", func(t *testing.T) {
		sqlMock.ExpectQuery(expectedQuery).
			WithArgs(msisdn).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "phone_number"}).
				AddRow(expectedData.ID, expectedData.Name, expectedData.MSISDN),
			)

		actualData, err := repoMock.GetRiderDataByMSISDN(ctx, msisdn)

		assert.Equal(t, expectedData, actualData)
		assert.Nil(t, err)
	})

	t.Run("should return error not found when data is not found", func(t *testing.T) {
		sqlMock.ExpectQuery(expectedQuery).
			WithArgs(msisdn).
			WillReturnError(sql.ErrNoRows)

		id, err := repoMock.GetRiderDataByMSISDN(ctx, msisdn)

		assert.Equal(t, model.RiderData{}, id)
		assert.Error(t, err, constants.ErrorDataNotFound)
	})

	t.Run("should return error when error from db", func(t *testing.T) {
		rowErr := errors.New("error from db")
		sqlMock.ExpectQuery(expectedQuery).
			WithArgs(msisdn).
			WillReturnError(rowErr)

		id, err := repoMock.GetRiderDataByMSISDN(ctx, msisdn)

		assert.Equal(t, model.RiderData{}, id)
		assert.NotNil(t, err)
	})
}

func TestRepository_GetRiderMSISDNByID(t *testing.T) {
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
	id := int64(111)
	msisdn := "081111"

	expectedQuery := queryGetRiderMSISDNByID

	t.Run("should execute get query", func(t *testing.T) {
		sqlMock.ExpectQuery(expectedQuery).
			WithArgs(id).
			WillReturnRows(sqlmock.NewRows([]string{"phone_number"}).
				AddRow(msisdn),
			)

		res, err := repoMock.GetRiderMSISDNByID(ctx, id)

		assert.Equal(t, msisdn, res)
		assert.Nil(t, err)
	})

	t.Run("should return not found when data not found", func(t *testing.T) {
		sqlMock.ExpectQuery(expectedQuery).
			WithArgs(id).
			WillReturnError(sql.ErrNoRows)

		res, err := repoMock.GetRiderMSISDNByID(ctx, id)

		assert.Equal(t, "", res)
		assert.Error(t, err, constants.ErrorDataNotFound)
	})

	t.Run("should return error when error from db", func(t *testing.T) {
		rowErr := errors.New("error from db")
		sqlMock.ExpectQuery(expectedQuery).
			WithArgs(id).
			WillReturnError(rowErr)

		res, err := repoMock.GetRiderMSISDNByID(ctx, id)

		assert.Equal(t, "", res)
		assert.NotNil(t, err)
	})
}

func TestRepository_GetDriverDataByMSISDN(t *testing.T) {
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
	expectedData := model.DriverData{
		ID:             int64(999),
		Name:           "Agus",
		MSISDN:         "0811111",
		VehicleType:    "CAR",
		VehicleTypeInt: 1,
		VehiclePlate:   "B1111A",
	}

	expectedQuery := queryGetDriverByMSISDN

	t.Run("should execute get query", func(t *testing.T) {
		sqlMock.ExpectQuery(expectedQuery).
			WithArgs(msisdn).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "phone_number", "vehicle_type", "vehicle_plate"}).
				AddRow(expectedData.ID, expectedData.Name, expectedData.MSISDN, expectedData.VehicleTypeInt, expectedData.VehiclePlate),
			)

		actualData, err := repoMock.GetDriverDataByMSISDN(ctx, msisdn)

		assert.Equal(t, expectedData, actualData)
		assert.Nil(t, err)
	})

	t.Run("should return not found when data is not found", func(t *testing.T) {
		sqlMock.ExpectQuery(expectedQuery).
			WithArgs(msisdn).
			WillReturnError(sql.ErrNoRows)

		id, err := repoMock.GetDriverDataByMSISDN(ctx, msisdn)

		assert.Equal(t, model.DriverData{}, id)
		assert.Error(t, err, constants.ErrorDataNotFound)
	})

	t.Run("should return error when error from db", func(t *testing.T) {
		rowErr := errors.New("error from db")
		sqlMock.ExpectQuery(expectedQuery).
			WithArgs(msisdn).
			WillReturnError(rowErr)

		id, err := repoMock.GetDriverDataByMSISDN(ctx, msisdn)

		assert.Equal(t, model.DriverData{}, id)
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

func TestRepository_ConfirmRideDriver(t *testing.T) {
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
	expected := model.RideData{
		RideID:   111,
		RiderID:  666,
		DriverID: 222,
		PickupLocation: model.Coordinate{
			Latitude:  1,
			Longitude: 2,
		},
		Destination: model.Coordinate{
			Latitude:  10,
			Longitude: 20,
		},
	}
	req := model.ConfirmRideDriverRequest{
		DriverID: 111,
		RideID:   777,
	}

	expectedQuery := queryConfirmRideDriver

	t.Run("should execute update returning query", func(t *testing.T) {
		sqlMock.ExpectBegin()
		sqlMock.ExpectQuery(expectedQuery).
			WithArgs(req.DriverID, req.RideID).
			WillReturnRows(sqlmock.NewRows([]string{
				"id", "rider_id", "driver_id", "pickup_location.latitude", "pickup_location.longitude",
				"destination.latitude", "destination.longitude",
			}).AddRow(
				expected.RideID, expected.RiderID, expected.DriverID, expected.PickupLocation.Latitude,
				expected.PickupLocation.Longitude, expected.Destination.Latitude, expected.Destination.Longitude))
		sqlMock.ExpectCommit()

		actualRes, err := repoMock.ConfirmRideDriver(ctx, req)

		assert.Equal(t, expected, actualRes)
		assert.Nil(t, err)
	})

	t.Run("should return error when error from db", func(t *testing.T) {
		rowErr := errors.New("error from db")
		sqlMock.ExpectBegin()
		sqlMock.ExpectQuery(expectedQuery).
			WithArgs(req.DriverID, req.RideID).
			WillReturnError(rowErr)
		sqlMock.ExpectRollback()

		actual, err := repoMock.ConfirmRideDriver(ctx, req)

		assert.Equal(t, model.RideData{}, actual)
		assert.NotNil(t, err)
	})

	t.Run("should return error when error begin tx", func(t *testing.T) {
		rowErr := errors.New("error from db")
		sqlMock.ExpectBegin().WillReturnError(rowErr)

		actual, err := repoMock.ConfirmRideDriver(ctx, req)

		assert.Equal(t, model.RideData{}, actual)
		assert.NotNil(t, err)
	})

	t.Run("should return error when error commit tx", func(t *testing.T) {
		rowErr := errors.New("error from db")
		sqlMock.ExpectBegin()
		sqlMock.ExpectQuery(expectedQuery).
			WithArgs(req.DriverID, req.RideID)
		sqlMock.ExpectRollback().WillReturnError(rowErr)

		actual, err := repoMock.ConfirmRideDriver(ctx, req)

		assert.Equal(t, model.RideData{}, actual)
		assert.NotNil(t, err)
	})
}

func TestRepository_ConfirmRideRider(t *testing.T) {
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
	expected := model.RideData{
		RideID:   111,
		RiderID:  666,
		DriverID: 222,
		PickupLocation: model.Coordinate{
			Latitude:  1,
			Longitude: 2,
		},
		Destination: model.Coordinate{
			Latitude:  10,
			Longitude: 20,
		},
	}
	req := model.ConfirmRideRiderRequest{
		RiderID: 666,
		RideID:  777,
	}

	expectedQuery := queryConfirmRideRider

	t.Run("should execute update returning query", func(t *testing.T) {
		sqlMock.ExpectBegin()
		sqlMock.ExpectQuery(expectedQuery).
			WithArgs(model.StatusNumRideWaitingForPickup, req.RideID, req.RiderID).
			WillReturnRows(sqlmock.NewRows([]string{
				"id", "rider_id", "driver_id", "pickup_location.latitude", "pickup_location.longitude",
				"destination.latitude", "destination.longitude",
			}).AddRow(
				expected.RideID, expected.RiderID, expected.DriverID, expected.PickupLocation.Latitude,
				expected.PickupLocation.Longitude, expected.Destination.Latitude, expected.Destination.Longitude))
		sqlMock.ExpectCommit()

		actualRes, err := repoMock.ConfirmRideRider(ctx, req)

		assert.Equal(t, expected, actualRes)
		assert.Nil(t, err)
	})

	t.Run("should return error when error from db", func(t *testing.T) {
		rowErr := errors.New("error from db")
		sqlMock.ExpectBegin()
		sqlMock.ExpectQuery(expectedQuery).
			WithArgs(model.StatusNumRideWaitingForPickup, req.RideID, req.RiderID).
			WillReturnError(rowErr)
		sqlMock.ExpectRollback()

		actual, err := repoMock.ConfirmRideRider(ctx, req)

		assert.Equal(t, model.RideData{}, actual)
		assert.NotNil(t, err)
	})

	t.Run("should return error when error begin tx", func(t *testing.T) {
		rowErr := errors.New("error from db")
		sqlMock.ExpectBegin().WillReturnError(rowErr)

		actual, err := repoMock.ConfirmRideRider(ctx, req)

		assert.Equal(t, model.RideData{}, actual)
		assert.NotNil(t, err)
	})

	t.Run("should return error when error commit tx", func(t *testing.T) {
		rowErr := errors.New("error from db")
		sqlMock.ExpectBegin()
		sqlMock.ExpectQuery(expectedQuery).
			WithArgs(model.StatusNumRideWaitingForPickup, req.RideID, req.RiderID)
		sqlMock.ExpectRollback().WillReturnError(rowErr)

		actual, err := repoMock.ConfirmRideRider(ctx, req)

		assert.Equal(t, model.RideData{}, actual)
		assert.NotNil(t, err)
	})
}
