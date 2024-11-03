package usecase

import (
	"context"
	"nebeng-jek/internal/location/model"
)

type LocationUsecase interface {
	AddAvailableDriver(ctx context.Context, driverID int64, location model.Coordinate) error
	RemoveAvailableDriver(ctx context.Context, driverID int64) error
	GetNearestAvailableDrivers(ctx context.Context, location model.Coordinate) ([]int64, error)
	GetRidePath(ctx context.Context, rideID int64, driverID int64) ([]model.Coordinate, error)
	TrackUserLocation(ctx context.Context, req model.TrackUserLocationRequest) error
}
