package repository_postgres

import (
	"context"
	"database/sql"
	"nebeng-jek/internal/pkg/constants"
	"nebeng-jek/internal/rides/model"
	"nebeng-jek/internal/rides/repository"
	"nebeng-jek/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type ridesRepo struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) repository.RidesRepository {
	return &ridesRepo{
		db: db,
	}
}

func (r *ridesRepo) GetRiderDataByMSISDN(ctx context.Context, msisdn string) (model.RiderData, error) {
	var data model.RiderData
	err := r.db.GetContext(ctx, &data, queryGetRiderByMSISDN, msisdn)
	if err == sql.ErrNoRows {
		return model.RiderData{}, constants.ErrorDataNotFound
	}
	if err != nil {
		return model.RiderData{}, err
	}
	return data, nil
}

func (r *ridesRepo) GetRiderMSISDNByID(ctx context.Context, id int64) (string, error) {
	var msisdn string
	err := r.db.GetContext(ctx, &msisdn, queryGetRiderMSISDNByID, id)
	if err != nil {
		return "", err
	}
	return msisdn, nil
}

func (r *ridesRepo) GetDriverDataByMSISDN(ctx context.Context, msisdn string) (model.DriverData, error) {
	var data model.DriverData
	err := r.db.GetContext(ctx, &data, queryGetDriverByMSISDN, msisdn)
	if err == sql.ErrNoRows {
		return model.DriverData{}, constants.ErrorDataNotFound
	}
	if err != nil {
		return model.DriverData{}, err
	}
	data.VehicleType = model.MapVehicleType[data.VehicleTypeInt]
	return data, nil
}

func (r *ridesRepo) CreateNewRide(ctx context.Context, req model.CreateNewRideRequest) (int64, error) {
	var id int64
	values := []interface{}{
		req.RiderID, model.StatusNumRideWaitingForDriver,
		req.PickupLocation.Latitude, req.PickupLocation.Longitude,
		req.Destination.Latitude, req.Destination.Longitude,
	}
	err := r.db.QueryRowContext(ctx, queryInsertNewRide, values...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *ridesRepo) ConfirmRideDriver(ctx context.Context, req model.ConfirmRideDriverRequest) (model.RideData, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return model.RideData{}, err
	}

	var data model.RideData
	values := []interface{}{
		req.DriverID, req.RideID,
	}
	err = tx.QueryRowxContext(ctx, queryConfirmRideDriver, values...).StructScan(&data)
	if err == sql.ErrNoRows {
		if err := tx.Rollback(); err != nil {
			logger.Error(ctx, "error rollback tx", nil)
		}
		return model.RideData{}, constants.ErrorDataNotFound
	}
	if err != nil {
		if err := tx.Rollback(); err != nil {
			logger.Error(ctx, "error rollback tx", nil)
		}
		return model.RideData{}, err
	}

	err = tx.Commit()
	if err != nil {
		if err := tx.Rollback(); err != nil {
			logger.Error(ctx, "error rollback tx", nil)
		}
		return model.RideData{}, err
	}

	return data, nil
}
