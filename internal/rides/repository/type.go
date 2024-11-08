package repository

import (
	"context"
	pkgLocation "nebeng-jek/internal/pkg/location"
	"nebeng-jek/internal/rides/model"
)

type RidesLocationRepository interface {
	AddAvailableDriver(ctx context.Context, driverID int64, location pkgLocation.Coordinate) error
	RemoveAvailableDriver(ctx context.Context, driverID int64) error
	GetNearestAvailableDrivers(ctx context.Context, location pkgLocation.Coordinate) ([]int64, error)
	GetRidePath(ctx context.Context, req model.GetRidePathRequest) (model.GetRidePathResponse, error)
}

type RidesRepository interface {
	GetRiderDataByID(ctx context.Context, riderID int64) (model.RiderData, error)
	GetDriverDataByID(ctx context.Context, driverID int64) (model.DriverData, error)
	UpdateDriverStatus(ctx context.Context, req model.UpdateDriverStatusRequest) error
	GetRiderMSISDNByID(ctx context.Context, id int64) (string, error)
	GetDriverMSISDNByID(ctx context.Context, id int64) (string, error)
	GetRideData(ctx context.Context, id int64) (model.RideData, error)

	CreateNewRide(context.Context, model.CreateNewRideRequest) (int64, error)
	UpdateRideData(ctx context.Context, req model.UpdateRideDataRequest) error
	StoreRideCommission(ctx context.Context, req model.StoreRideCommissionRequest) error
}

type RidesPubsubRepository interface {
	BroadcastMessage(context.Context, string, interface{}) error
}

type PaymentRepository interface {
	DeductCredit(context.Context, model.DeductCreditRequest) error
	AddCredit(context.Context, model.AddCreditRequest) error
}
