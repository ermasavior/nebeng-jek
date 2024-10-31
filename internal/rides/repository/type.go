package repository

import (
	"context"
	"nebeng-jek/internal/rides/model"
)

type RidesLocationRepository interface {
	AddAvailableDriver(ctx context.Context, driverID int64, location model.Coordinate) error
	RemoveAvailableDriver(ctx context.Context, driverID int64) error
	GetNearestAvailableDrivers(ctx context.Context, location model.Coordinate) ([]int64, error)
	GetRidePath(ctx context.Context, rideID int64, driverID int64) ([]model.Coordinate, error)
	TrackUserLocation(ctx context.Context, req model.TrackUserLocationRequest) error
}

type RidesRepository interface {
	GetRiderDataByID(ctx context.Context, riderID int64) (model.RiderData, error)
	GetDriverDataByID(ctx context.Context, driverID int64) (model.DriverData, error)
	GetRiderMSISDNByID(ctx context.Context, id int64) (string, error)
	GetDriverMSISDNByID(ctx context.Context, id int64) (string, error)
	GetRideData(ctx context.Context, id int64) (model.RideData, error)

	CreateNewRide(context.Context, model.CreateNewRideRequest) (int64, error)
	UpdateRideData(ctx context.Context, req model.UpdateRideDataRequest) error
}

type RidesPubsubRepository interface {
	BroadcastMessage(context.Context, string, interface{}) error
}
