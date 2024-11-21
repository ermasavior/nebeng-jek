package repository_postgres

import (
	"context"
	"database/sql"
	"errors"
	"nebeng-jek/internal/pkg/constants"
	pkgLocation "nebeng-jek/internal/pkg/location"
	"nebeng-jek/internal/rides/model"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestRepository_GetRiderDataByID(t *testing.T) {
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
	riderID := int64(9999)
	expectedData := model.RiderData{
		ID:     riderID,
		Name:   "Melati",
		MSISDN: "0822222",
	}

	expectedQuery := queryGetRiderDataByID

	t.Run("should execute get query", func(t *testing.T) {
		sqlMock.ExpectQuery(expectedQuery).
			WithArgs(riderID).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "phone_number"}).
				AddRow(expectedData.ID, expectedData.Name, expectedData.MSISDN),
			)

		actualData, err := repoMock.GetRiderDataByID(ctx, riderID)

		assert.Equal(t, expectedData, actualData)
		assert.Nil(t, err)
	})

	t.Run("should return error not found when data is not found", func(t *testing.T) {
		sqlMock.ExpectQuery(expectedQuery).
			WithArgs(riderID).
			WillReturnError(sql.ErrNoRows)

		id, err := repoMock.GetRiderDataByID(ctx, riderID)

		assert.Equal(t, model.RiderData{}, id)
		assert.Error(t, err, constants.ErrorDataNotFound)
	})

	t.Run("should return error when error from db", func(t *testing.T) {
		rowErr := errors.New("error from db")
		sqlMock.ExpectQuery(expectedQuery).
			WithArgs(riderID).
			WillReturnError(rowErr)

		id, err := repoMock.GetRiderDataByID(ctx, riderID)

		assert.Equal(t, model.RiderData{}, id)
		assert.NotNil(t, err)
	})
}

func TestRepository_GetDriverDataByID(t *testing.T) {
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
	driverID := int64(1111)
	expectedData := model.DriverData{
		ID:             driverID,
		Name:           "Agus",
		Status:         model.StatusDriverAvailable,
		MSISDN:         "0811111",
		VehicleType:    "CAR",
		VehicleTypeInt: 1,
		VehiclePlate:   "B1111A",
	}

	expectedQuery := queryGetDriverDataByID

	t.Run("should execute get query", func(t *testing.T) {
		sqlMock.ExpectQuery(expectedQuery).
			WithArgs(driverID).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "phone_number", "vehicle_type", "vehicle_plate", "status"}).
				AddRow(expectedData.ID, expectedData.Name, expectedData.MSISDN, expectedData.VehicleTypeInt, expectedData.VehiclePlate, expectedData.Status),
			)

		actualData, err := repoMock.GetDriverDataByID(ctx, driverID)

		assert.Equal(t, expectedData, actualData)
		assert.Nil(t, err)
	})

	t.Run("should return not found when data is not found", func(t *testing.T) {
		sqlMock.ExpectQuery(expectedQuery).
			WithArgs(driverID).
			WillReturnError(sql.ErrNoRows)

		id, err := repoMock.GetDriverDataByID(ctx, driverID)

		assert.Equal(t, model.DriverData{}, id)
		assert.Error(t, err, constants.ErrorDataNotFound)
	})

	t.Run("should return error when error from db", func(t *testing.T) {
		rowErr := errors.New("error from db")
		sqlMock.ExpectQuery(expectedQuery).
			WithArgs(driverID).
			WillReturnError(rowErr)

		id, err := repoMock.GetDriverDataByID(ctx, driverID)

		assert.Equal(t, model.DriverData{}, id)
		assert.NotNil(t, err)
	})
}

func TestRepository_UpdateDriverStatus(t *testing.T) {
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
	req := model.UpdateDriverStatusRequest{
		DriverID: 1111,
		Status:   model.StatusDriverAvailable,
	}
	expectedQuery := queryUpdateDriverStatus

	t.Run("should execute update query", func(t *testing.T) {
		sqlMock.ExpectExec(expectedQuery).
			WithArgs(req.Status, req.DriverID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repoMock.UpdateDriverStatus(ctx, req)

		assert.Nil(t, err)
	})

	t.Run("should return error when error from db", func(t *testing.T) {
		rowErr := errors.New("error from db")
		sqlMock.ExpectExec(expectedQuery).
			WithArgs(req.Status, req.DriverID).
			WillReturnError(rowErr)

		err := repoMock.UpdateDriverStatus(ctx, req)

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
	riderID := int64(1111)
	expectedMSISDN := "081111"

	expectedQuery := queryGetRiderMSISDNByID

	t.Run("should execute get query", func(t *testing.T) {
		sqlMock.ExpectQuery(expectedQuery).
			WithArgs(riderID).
			WillReturnRows(sqlmock.NewRows([]string{"phone_number"}).
				AddRow(expectedMSISDN),
			)

		msisdn, err := repoMock.GetRiderMSISDNByID(ctx, riderID)

		assert.Equal(t, expectedMSISDN, msisdn)
		assert.Nil(t, err)
	})

	t.Run("should return not found when data is not found", func(t *testing.T) {
		sqlMock.ExpectQuery(expectedQuery).
			WithArgs(riderID).
			WillReturnError(sql.ErrNoRows)

		msisdn, err := repoMock.GetRiderMSISDNByID(ctx, riderID)

		assert.Equal(t, "", msisdn)
		assert.Error(t, err, constants.ErrorDataNotFound)
	})

	t.Run("should return error when error from db", func(t *testing.T) {
		rowErr := errors.New("error from db")
		sqlMock.ExpectQuery(expectedQuery).
			WithArgs(riderID).
			WillReturnError(rowErr)

		msisdn, err := repoMock.GetRiderMSISDNByID(ctx, riderID)

		assert.Equal(t, "", msisdn)
		assert.NotNil(t, err)
	})
}

func TestRepository_GetDriverMSISDNByID(t *testing.T) {
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
	driverID := int64(1111)
	expectedMSISDN := "081111"

	expectedQuery := queryGetDriverMSISDNByID

	t.Run("should execute get query", func(t *testing.T) {
		sqlMock.ExpectQuery(expectedQuery).
			WithArgs(driverID).
			WillReturnRows(sqlmock.NewRows([]string{"phone_number"}).
				AddRow(expectedMSISDN),
			)

		msisdn, err := repoMock.GetDriverMSISDNByID(ctx, driverID)

		assert.Equal(t, expectedMSISDN, msisdn)
		assert.Nil(t, err)
	})

	t.Run("should return not found when data is not found", func(t *testing.T) {
		sqlMock.ExpectQuery(expectedQuery).
			WithArgs(driverID).
			WillReturnError(sql.ErrNoRows)

		msisdn, err := repoMock.GetDriverMSISDNByID(ctx, driverID)

		assert.Equal(t, "", msisdn)
		assert.Error(t, err, constants.ErrorDataNotFound)
	})

	t.Run("should return error when error from db", func(t *testing.T) {
		rowErr := errors.New("error from db")
		sqlMock.ExpectQuery(expectedQuery).
			WithArgs(driverID).
			WillReturnError(rowErr)

		msisdn, err := repoMock.GetDriverMSISDNByID(ctx, driverID)

		assert.Equal(t, "", msisdn)
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
		PickupLocation: pkgLocation.Coordinate{
			Latitude:  1,
			Longitude: 2,
		},
		Destination: pkgLocation.Coordinate{
			Latitude:  9,
			Longitude: 10,
		},
	}

	expectedQuery := queryInsertNewRide

	t.Run("should execute insert query", func(t *testing.T) {
		sqlMock.ExpectQuery(expectedQuery).
			WithArgs(
				req.RiderID, model.StatusNumRideNewRequest,
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
				req.RiderID, model.StatusNumRideNewRequest,
				req.PickupLocation.Latitude, req.PickupLocation.Longitude,
				req.Destination.Latitude, req.Destination.Longitude).
			WillReturnError(rowErr)

		id, err := repoMock.CreateNewRide(ctx, req)

		assert.Equal(t, int64(0), id)
		assert.NotNil(t, err)
	})
}

func TestRepository_GetRideData(t *testing.T) {
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
	rideID := int64(111)
	driverID := int64(222)
	expected := model.RideData{
		RideID:   111,
		RiderID:  666,
		DriverID: &driverID,
		PickupLocation: pkgLocation.Coordinate{
			Latitude:  1,
			Longitude: 2,
		},
		Destination: pkgLocation.Coordinate{
			Latitude:  10,
			Longitude: 20,
		},
	}

	expectedQuery := queryGetRideData

	t.Run("should execute get query", func(t *testing.T) {
		sqlMock.ExpectQuery(expectedQuery).
			WithArgs(rideID).
			WillReturnRows(sqlmock.NewRows([]string{
				"id", "rider_id", "driver_id", "pickup_location.latitude", "pickup_location.longitude",
				"destination.latitude", "destination.longitude",
			}).AddRow(
				expected.RideID, expected.RiderID, expected.DriverID, expected.PickupLocation.Latitude,
				expected.PickupLocation.Longitude, expected.Destination.Latitude, expected.Destination.Longitude))

		actualData, err := repoMock.GetRideData(ctx, rideID)

		assert.Equal(t, expected, actualData)
		assert.Nil(t, err)
	})

	t.Run("should return error not found when data is not found", func(t *testing.T) {
		sqlMock.ExpectQuery(expectedQuery).
			WithArgs(rideID).
			WillReturnError(sql.ErrNoRows)

		id, err := repoMock.GetRideData(ctx, rideID)

		assert.Equal(t, model.RideData{}, id)
		assert.Error(t, err, constants.ErrorDataNotFound)
	})

	t.Run("should return error when error from db", func(t *testing.T) {
		rowErr := errors.New("error from db")
		sqlMock.ExpectQuery(expectedQuery).
			WithArgs(rideID).
			WillReturnError(rowErr)

		id, err := repoMock.GetRideData(ctx, rideID)

		assert.Equal(t, model.RideData{}, id)
		assert.NotNil(t, err)
	})
}

func TestRepository_UpdateRideData(t *testing.T) {
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

	fare, distance, finalPrice := float64(10000), float64(10), float64(10000)
	ctx := context.Background()

	t.Run("should execute update returning query", func(t *testing.T) {
		req := model.UpdateRideDataRequest{
			DriverID: 222,
			RideID:   777,
			Status:   model.StatusNumRideNewRequest,
		}

		expectedQuery := `
			UPDATE rides
			SET status = $1, driver_id = $2, updated_at = NOW()
			WHERE id = $3
		`
		sqlMock.ExpectExec(expectedQuery).
			WithArgs(req.Status, req.DriverID, req.RideID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repoMock.UpdateRideData(ctx, req)
		assert.Nil(t, err)
	})

	t.Run("should execute update returning query - update fare", func(t *testing.T) {
		req := model.UpdateRideDataRequest{
			RideID:     777,
			Distance:   &distance,
			Fare:       &fare,
			FinalPrice: &finalPrice,
			Status:     model.StatusNumRideEnded,
		}

		expectedQuery := `
			UPDATE rides
			SET status = $1, distance = $2, fare = $3, final_price = $4, updated_at = NOW()
			WHERE id = $5
		`
		sqlMock.ExpectExec(expectedQuery).
			WithArgs(req.Status, distance, fare, finalPrice, req.RideID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repoMock.UpdateRideData(ctx, req)
		assert.Nil(t, err)
	})

	t.Run("should execute update returning query - update start & end time", func(t *testing.T) {
		currTime := time.Now()
		req := model.UpdateRideDataRequest{
			RideID:    777,
			StartTime: &currTime,
			EndTime:   &currTime,
		}

		expectedQuery := `
			UPDATE rides
			SET start_time = $1, end_time = $2, updated_at = NOW()
			WHERE id = $3
		`
		sqlMock.ExpectExec(expectedQuery).
			WithArgs(currTime, currTime, req.RideID).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repoMock.UpdateRideData(ctx, req)
		assert.Nil(t, err)
	})

	t.Run("should return error - failed executing query", func(t *testing.T) {
		req := model.UpdateRideDataRequest{
			RideID:     777,
			Distance:   &distance,
			Fare:       &fare,
			FinalPrice: &finalPrice,
			Status:     model.StatusNumRideEnded,
		}

		expectedQuery := `
			UPDATE rides
			SET status = $1, distance = $2, fare = $3, final_price = $4, updated_at = NOW()
			WHERE id = $5
		`
		sqlMock.ExpectExec(expectedQuery).
			WithArgs(req.Status, distance, fare, finalPrice, req.RideID).
			WillReturnError(errors.New("error from sql"))

		err := repoMock.UpdateRideData(ctx, req)
		assert.Error(t, err)
	})
}

func TestRepository_StoreRideCommission(t *testing.T) {
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
	req := model.StoreRideCommissionRequest{
		RideID:           1111,
		PlatformFee:      2000,
		DriverCommission: 20000,
	}

	expectedQuery := queryInsertRideCommission

	t.Run("should execute insert query", func(t *testing.T) {
		sqlMock.ExpectExec(expectedQuery).
			WithArgs(req.RideID, req.PlatformFee, req.DriverCommission).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repoMock.StoreRideCommission(ctx, req)

		assert.Nil(t, err)
	})

	t.Run("should return error when error from db", func(t *testing.T) {
		rowErr := errors.New("error from db")
		sqlMock.ExpectExec(expectedQuery).
			WithArgs(req.RideID, req.PlatformFee, req.DriverCommission).
			WillReturnError(rowErr)

		err := repoMock.StoreRideCommission(ctx, req)

		assert.NotNil(t, err)
	})
}
