package usecase

import (
	"context"
	pkgContext "nebeng-jek/internal/pkg/context"
	"nebeng-jek/internal/rides/model"
	pkgError "nebeng-jek/pkg/error"
	"nebeng-jek/pkg/logger"
)

func (u *ridesUsecase) CreateNewRide(ctx context.Context, req model.CreateNewRideRequest) (int64, error) {
	msisdn := pkgContext.GetMSISDNFromContext(ctx)

	riderID, err := u.ridesRepo.GetRiderIDByMSISDN(ctx, msisdn)
	if err != nil {
		logger.Error(ctx, "error get rider data", map[string]interface{}{
			"msisdn": msisdn,
			"error":  err,
		})
		return 0, pkgError.NewInternalServerError(err, "error get rider data")
	}

	req.RiderID = riderID

	rideID, err := u.ridesRepo.CreateNewRide(ctx, req)
	if err != nil {
		logger.Error(ctx, "error create new ride", map[string]interface{}{
			"msisdn": msisdn,
			"error":  err,
		})
		return 0, pkgError.NewInternalServerError(err, "error create new ride")
	}

	drivers, err := u.locationRepo.GetNearestAvailableDrivers(ctx, req.PickupLocation)
	if err != nil {
		logger.Error(ctx, "error get nearest available drivers", map[string]interface{}{
			"msisdn": msisdn,
			"error":  err,
		})
		return 0, pkgError.NewInternalServerError(err, "error get nearest available drivers")
	}
	if len(drivers) == 0 {
		return 0, pkgError.NewNotFound(nil, "no nearest driver available")
	}

	var mapDrivers = map[string]bool{}
	for _, d := range drivers {
		mapDrivers[d] = true
	}

	err = u.ridesPubSub.BroadcastRideToDrivers(ctx, model.RideRequestMessage{
		RideID:           rideID,
		RiderID:          riderID,
		PickupLocation:   req.PickupLocation,
		Destination:      req.Destination,
		AvailableDrivers: mapDrivers,
	})

	if err != nil {
		logger.Error(ctx, "error broadcasting ride to drivers", map[string]interface{}{
			"msisdn": msisdn,
			"error":  err,
		})
		return 0, pkgError.NewInternalServerError(err, "error broadcasting ride to drivers")
	}

	return rideID, nil
}
