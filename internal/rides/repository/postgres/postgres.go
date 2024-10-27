package repository_postgres

import (
	"context"
	"database/sql"
	"fmt"
	"nebeng-jek/internal/pkg/constants"
	"nebeng-jek/internal/rides/model"
	"nebeng-jek/internal/rides/repository"
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

func (r *ridesRepo) GetRiderDataByID(ctx context.Context, riderID int64) (model.RiderData, error) {
	var data model.RiderData
	err := r.db.GetContext(ctx, &data, queryGetRiderDataByID, riderID)
	if err == sql.ErrNoRows {
		return model.RiderData{}, constants.ErrorDataNotFound
	}
	if err != nil {
		return model.RiderData{}, err
	}
	return data, nil
}

func (r *ridesRepo) GetDriverDataByID(ctx context.Context, driverID int64) (model.DriverData, error) {
	var data model.DriverData
	err := r.db.GetContext(ctx, &data, queryGetDriverDataByID, driverID)
	if err == sql.ErrNoRows {
		return model.DriverData{}, constants.ErrorDataNotFound
	}
	if err != nil {
		return model.DriverData{}, err
	}
	data.MapVehicleType()
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

func (r *ridesRepo) CreateNewRide(ctx context.Context, req model.CreateNewRideRequest) (int64, error) {
	var id int64
	values := []interface{}{
		req.RiderID, model.StatusNumRideNewRequest,
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

	return data, nil
}

func (r *ridesRepo) UpdateRideData(ctx context.Context, req model.UpdateRideDataRequest) error {
	var (
		params     = []interface{}{}
		querySet   = []string{}
		queryWhere = ""
		paramNum   = 0
	)

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
	if req.Distance != nil {
		paramNum += 1
		params = append(params, *req.Distance)
		querySet = append(querySet, fmt.Sprintf("distance = $%d", paramNum))
	}
	if req.Fare != nil {
		paramNum += 1
		params = append(params, *req.Fare)
		querySet = append(querySet, fmt.Sprintf("fare = $%d", paramNum))
	}
	if req.FinalPrice != nil {
		paramNum += 1
		params = append(params, *req.FinalPrice)
		querySet = append(querySet, fmt.Sprintf("final_price = $%d", paramNum))
	}

	paramNum += 1
	queryWhere = fmt.Sprintf("id = $%d", paramNum)
	params = append(params, req.RideID)

	query := fmt.Sprintf(queryUpdateRideData, strings.Join(querySet, ", "), queryWhere)
	_, err := r.db.ExecContext(ctx, query, params...)
	if err != nil {
		return err
	}

	return nil
}
