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
	GetRiderDataByMSISDN(ctx context.Context, msisdn string) (model.RiderData, error)
	GetDriverDataByMSISDN(ctx context.Context, msisdn string) (model.DriverData, error)
	GetRiderMSISDNByID(ctx context.Context, id int64) (string, error)
	GetDriverMSISDNByID(ctx context.Context, id int64) (string, error)
	CreateNewRide(context.Context, model.CreateNewRideRequest) (int64, error)
	ConfirmRideDriver(ctx context.Context, req model.ConfirmRideDriverRequest) (model.RideData, error)
	ConfirmRideRider(ctx context.Context, req model.ConfirmRideRiderRequest) (model.RideData, error)
	UpdateRideByDriver(ctx context.Context, req model.UpdateRideByDriverRequest) (model.RideData, error)
}

type RidesPubsubRepository interface {
	BroadcastRideToDrivers(context.Context, model.RideRequestMessage) error
	BroadcastMatchedRideToRider(context.Context, model.MatchedRideMessage) error
	BroadcastRideReadyToPickup(context.Context, model.RideReadyToPickupMessage) error
	BroadcastRideStarted(context.Context, model.RideStartedMessage) error
}
