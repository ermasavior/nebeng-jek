package repository

import (
	"context"
	"nebeng-jek/internal/rides/model"
)

type RidesLocationRepository interface {
	AddAvailableDriver(context.Context, string, model.Coordinate) error
	RemoveAvailableDriver(context.Context, string) error
	GetNearestAvailableDrivers(context.Context, model.Coordinate) ([]string, error)
}

type RidesRepository interface {
	GetRiderIDByMSISDN(ctx context.Context, msisdn string) (int64, error)
	CreateNewRide(context.Context, model.CreateNewRideRequest) (int64, error)
}

type RidesPubsubRepository interface {
	BroadcastRideToDrivers(context.Context, model.RideRequestMessage) error
}
