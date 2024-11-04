package usecase

import (
	"context"
	"nebeng-jek/internal/location/model"
	pkgLocation "nebeng-jek/internal/pkg/location"
)

type LocationUsecase interface {
	AddAvailableDriver(ctx context.Context, driverID int64, location pkgLocation.Coordinate) error
	RemoveAvailableDriver(ctx context.Context, driverID int64) error
	GetNearestAvailableDrivers(ctx context.Context, location pkgLocation.Coordinate) ([]int64, error)
	GetRidePath(ctx context.Context, rideID int64, driverID int64) ([]pkgLocation.Coordinate, error)
	TrackUserLocation(ctx context.Context, req model.TrackUserLocationRequest) error
}
