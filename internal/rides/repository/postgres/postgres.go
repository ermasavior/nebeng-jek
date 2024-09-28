package repository_postgres

import (
	"context"
	"nebeng-jek/internal/rides/model"
	"nebeng-jek/internal/rides/repository"

	"github.com/jmoiron/sqlx"
)

type ridesRepo struct {
	db *sqlx.DB
}

func NewRidesRepository(db *sqlx.DB) repository.RidesRepository {
	return &ridesRepo{
		db: db,
	}
}

func (r *ridesRepo) GetRiderIDByMSISDN(ctx context.Context, msisdn string) (int64, error) {
	var id int64
	err := r.db.GetContext(ctx, &id, queryGetRiderByMSISDN, msisdn)
	if err != nil {
		return 0, err
	}
	return id, nil
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
