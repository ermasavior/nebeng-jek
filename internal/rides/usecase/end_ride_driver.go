package usecase

import (
	"context"
	"nebeng-jek/internal/pkg/constants"
	pkgContext "nebeng-jek/internal/pkg/context"
	"nebeng-jek/internal/pkg/haversine"
	"nebeng-jek/internal/rides/model"
	pkgError "nebeng-jek/pkg/error"
	"nebeng-jek/pkg/logger"
)

func (u *ridesUsecase) EndRideDriver(ctx context.Context, req model.EndRideDriverRequest) (model.RideData, *pkgError.AppError) {
	msisdn := pkgContext.GetMSISDNFromContext(ctx)

	driver, err := u.ridesRepo.GetDriverDataByMSISDN(ctx, msisdn)
	if err == constants.ErrorDataNotFound {
		return model.RideData{}, pkgError.NewNotFound(err, "driver is not found")
	}
	if err != nil {
		logger.Error(ctx, "error get driver data", map[string]interface{}{
			"msisdn": msisdn,
			"error":  err,
		})
		return model.RideData{}, pkgError.NewInternalServerError(err, "error get driver data")
	}

	ridePath, err := u.locationRepo.GetRidePath(ctx, req.RideID, msisdn)
	if err != nil {
		logger.Error(ctx, "error get distance traversed", map[string]interface{}{
			"msisdn": msisdn,
			"error":  err,
		})
		return model.RideData{}, pkgError.NewInternalServerError(err, "error get distance traversed")
	}

	distance := calculateTotalDistance(ridePath)

	rideData, err := u.ridesRepo.UpdateRideByDriver(ctx, model.UpdateRideByDriverRequest{
		DriverID: driver.ID,
		RideID:   req.RideID,
		Status:   model.StatusNumRideEnded,
		Distance: distance,
	})
	if err == constants.ErrorDataNotFound {
		return model.RideData{}, pkgError.NewNotFound(err, "ride data is not found or has been allocated to another driver")
	}
	if err != nil {
		logger.Error(ctx, "error update ride by driver", map[string]interface{}{
			"msisdn": msisdn,
			"error":  err,
		})
		return model.RideData{}, pkgError.NewInternalServerError(err, "error update ride by driver")
	}

	rideData.SetDistance(distance)
	rideData.CalculateRideFare(distance)

	riderMSISDN, err := u.ridesRepo.GetRiderMSISDNByID(ctx, rideData.RiderID)
	if err != nil {
		logger.Error(ctx, "error get rider msisdn", map[string]interface{}{
			"msisdn": msisdn,
			"error":  err,
		})
		return model.RideData{}, pkgError.NewInternalServerError(err, "error get rider msisdn")
	}

	err = u.ridesPubSub.BroadcastMessage(ctx, constants.TopicRideEnded, model.RideEndedMessage{
		RideID:      rideData.RideID,
		Distance:    *rideData.Distance,
		Fare:        *rideData.Fare,
		RiderMSISDN: riderMSISDN,
	})
	if err != nil {
		logger.Error(ctx, "error broadcasting ride ended", map[string]interface{}{
			"msisdn": msisdn,
			"error":  err,
		})
		return model.RideData{}, pkgError.NewInternalServerError(err, "error broadcasting ride ended")
	}

	return rideData, nil
}

func calculateTotalDistance(path []model.Coordinate) float64 {
	distance := float64(0)
	if len(path) >= 2 {
		i := 0
		for _, posB := range path[1:] {
			posA := path[i]
			distance += haversine.CalculateDistance(posA.Latitude, posA.Longitude, posB.Latitude, posB.Longitude)
			i += 1
		}
	}

	return distance
}
