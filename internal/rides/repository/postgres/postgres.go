package repository_postgres

import (
	"context"
	"database/sql"
	"fmt"
	"nebeng-jek/internal/pkg/constants"
	"nebeng-jek/internal/rides/model"
	"nebeng-jek/internal/rides/repository"
	"nebeng-jek/pkg/logger"
	"strings"

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
	if err == sql.ErrNoRows {
		return "", constants.ErrorDataNotFound
	}
	if err != nil {
		return "", err
	}
	return msisdn, nil
}

func (r *ridesRepo) GetDriverMSISDNByID(ctx context.Context, id int64) (string, error) {
	var msisdn string
	err := r.db.GetContext(ctx, &msisdn, queryGetDriverMSISDNByID, id)
	if err == sql.ErrNoRows {
		return "", constants.ErrorDataNotFound
	}
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
	data.SetVehicleType()
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

func (r *ridesRepo) GetRideData(ctx context.Context, rideID int64) (model.RideData, error) {
	var data model.RideData
	err := r.db.GetContext(ctx, &data, queryGetRideData, rideID)
	if err == sql.ErrNoRows {
		return model.RideData{}, constants.ErrorDataNotFound
	}
	if err != nil {
		return model.RideData{}, err
	}

	data.SetStatus()

	return data, nil
}

func (r *ridesRepo) UpdateRideData(ctx context.Context, req model.UpdateRideDataRequest) error {
	var (
		params     = []interface{}{}
		paramNum   = 0
		querySet   = []string{}
		queryWhere = "id = $%d"
	)
	if req.Distance != 0 {
		paramNum += 1
		params = append(params, req.Distance)
		querySet = append(querySet, fmt.Sprintf("distance = $%d", paramNum))
	}
	if req.Fare != 0 {
		paramNum += 1
		params = append(params, req.Fare)
		querySet = append(querySet, fmt.Sprintf("fare = $%d", paramNum))
	}
	if req.FinalPrice != 0 {
		paramNum += 1
		params = append(params, req.FinalPrice)
		querySet = append(querySet, fmt.Sprintf("final_price = $%d", paramNum))
	}
	if req.Status != 0 {
		paramNum += 1
		params = append(params, req.Status)
		querySet = append(querySet, fmt.Sprintf("status = $%d", paramNum))
	}
	if req.DriverID != 0 {
		paramNum += 1
		params = append(params, req.DriverID)
		querySet = append(querySet, fmt.Sprintf("driver_id = $%d", paramNum))
	}

	paramNum += 1
	queryWhere = fmt.Sprintf("id = $%d", paramNum)
	params = append(params, req.RideID)

	query := fmt.Sprintf(queryUpdateRideData, strings.Join(querySet, ", "), queryWhere)
	fmt.Println(query, params)
	_, err := r.db.ExecContext(ctx, query, params...)
	if err == sql.ErrNoRows {
		return constants.ErrorDataNotFound
	}
	if err != nil {
		return err
	}

	return nil
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

func (r *ridesRepo) ConfirmRideRider(ctx context.Context, req model.ConfirmRideRiderRequest) (model.RideData, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return model.RideData{}, err
	}

	var data model.RideData
	values := []interface{}{
		model.StatusNumRideWaitingForPickup, req.RideID, req.RiderID,
	}
	err = tx.QueryRowxContext(ctx, queryConfirmRideRider, values...).StructScan(&data)
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

func (r *ridesRepo) UpdateRideByDriver(ctx context.Context, req model.UpdateRideByDriverRequest) (model.RideData, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return model.RideData{}, err
	}

	var data model.RideData
	values := []interface{}{
		req.Status, req.Distance, req.Fare, req.FinalPrice, req.RideID, req.DriverID,
	}
	err = tx.QueryRowxContext(ctx, queryUpdateRideByDriver, values...).StructScan(&data)
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
