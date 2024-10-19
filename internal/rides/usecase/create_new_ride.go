package usecase

import (
	"context"
	"nebeng-jek/internal/pkg/constants"
	pkgContext "nebeng-jek/internal/pkg/context"
	"nebeng-jek/internal/rides/model"
	pkgError "nebeng-jek/pkg/error"
	"nebeng-jek/pkg/logger"
)

func (u *ridesUsecase) CreateNewRide(ctx context.Context, req model.CreateNewRideRequest) (int64, *pkgError.AppError) {
	msisdn := pkgContext.GetMSISDNFromContext(ctx)

	rider, err := u.ridesRepo.GetRiderDataByMSISDN(ctx, msisdn)
	if err == constants.ErrorDataNotFound {
		return 0, pkgError.NewNotFound(err, err.Error())
	}
	if err != nil {
		logger.Error(ctx, "error get rider data", map[string]interface{}{
			"msisdn": msisdn,
			"error":  err,
		})
		return 0, pkgError.NewInternalServerError(err, "error get rider data")
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
		return 0, pkgError.NewNotFound(nil, "no nearest driver available, try again later")
	}

	req.RiderID = rider.ID
	rideID, err := u.ridesRepo.CreateNewRide(ctx, req)
	if err != nil {
		logger.Error(ctx, "error create new ride", map[string]interface{}{
			"msisdn": msisdn,
			"req":    req,
			"error":  err,
		})
		return 0, pkgError.NewInternalServerError(err, "error create new ride")
	}

	var mapDrivers = map[string]bool{}
	for _, d := range drivers {
		mapDrivers[d] = true
	}

	msg := model.RideRequestMessage{
		RideID:           rideID,
		Rider:            rider,
		PickupLocation:   req.PickupLocation,
		Destination:      req.Destination,
		AvailableDrivers: mapDrivers,
	}
	err = u.ridesPubSub.BroadcastMessage(ctx, constants.TopicRideNewRequest, msg)

	if err != nil {
		logger.Error(ctx, "error broadcasting ride to drivers", map[string]interface{}{
			"msisdn": msisdn,
			"msg":    msg,
			"error":  err,
		})
		return 0, pkgError.NewInternalServerError(err, "error broadcasting ride to drivers")
	}

	return rideID, nil
}
