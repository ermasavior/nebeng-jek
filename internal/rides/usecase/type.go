package usecase

import (
	"context"
	"nebeng-jek/internal/rides/model"
	pkgError "nebeng-jek/pkg/error"
)

type RidesUsecase interface {
	SetDriverAvailability(context.Context, model.SetDriverAvailabilityRequest) *pkgError.AppError
	CreateNewRide(context.Context, model.CreateNewRideRequest) (int64, *pkgError.AppError)
	ConfirmRideDriver(context.Context, model.ConfirmRideDriverRequest) *pkgError.AppError
	ConfirmRideRider(context.Context, model.ConfirmRideRiderRequest) *pkgError.AppError
	StartRideDriver(context.Context, model.StartRideDriverRequest) (model.RideData, *pkgError.AppError)
	EndRideDriver(context.Context, model.EndRideDriverRequest) (model.RideData, *pkgError.AppError)
}
