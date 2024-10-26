package usecase

import (
	"context"
	"nebeng-jek/internal/rides/model"
	pkgError "nebeng-jek/pkg/error"
)

type RidesUsecase interface {
	DriverSetAvailability(context.Context, model.DriverSetAvailabilityRequest) pkgError.AppError
	RiderCreateNewRide(context.Context, model.RiderCreateNewRideRequest) (int64, pkgError.AppError)
	DriverConfirmRide(context.Context, model.DriverConfirmRideRequest) pkgError.AppError
	RiderConfirmRide(context.Context, model.RiderConfirmRideRequest) pkgError.AppError
	DriverStartRide(context.Context, model.DriverStartRideRequest) (model.RideData, pkgError.AppError)
	DriverEndRide(context.Context, model.DriverEndRideRequest) (model.RideData, pkgError.AppError)
	DriverConfirmPrice(context.Context, model.DriverConfirmPriceRequest) (model.RideData, pkgError.AppError)
	TrackUserLocation(ctx context.Context, req model.TrackUserLocationRequest) pkgError.AppError
}
