package usecase

import (
	"context"
	"fmt"
	"math"
	"nebeng-jek/internal/pkg/constants"
	pkgContext "nebeng-jek/internal/pkg/context"
	pkgLocation "nebeng-jek/internal/pkg/location"
	"nebeng-jek/internal/rides/model"
	pkgError "nebeng-jek/pkg/error"
	"nebeng-jek/pkg/haversine"
	"nebeng-jek/pkg/logger"
	"time"
)

func (u *ridesUsecase) DriverEndRide(ctx context.Context, req model.DriverEndRideRequest) (model.RideData, pkgError.AppError) {
	driverID := pkgContext.GetDriverIDFromContext(ctx)

	rideData, err := u.ridesRepo.GetRideData(ctx, req.RideID)
	if err == constants.ErrorDataNotFound {
		return model.RideData{}, pkgError.NewNotFoundError(pkgError.ErrResourceNotFoundMsg)
	}
	if err != nil {
		logger.Error(ctx, model.ErrMsgFailGetRideData, map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return model.RideData{}, pkgError.NewInternalServerError(model.ErrMsgFailGetRideData)
	}

	ridePath, err := u.locationRepo.GetRidePath(ctx, model.GetRidePathRequest{
		RideID:   rideData.RideID,
		RiderID:  rideData.RiderID,
		DriverID: driverID,
	})
	if err != nil {
		logger.Error(ctx, "error get distance traversed", map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return model.RideData{}, pkgError.NewInternalServerError("error get distance traversed")
	}

	if err := model.ValidateEndRide(rideData, driverID, ridePath); err != nil {
		return model.RideData{}, err
	}

	distance := calculateTotalDistance(ridePath.DriverPath)
	fare := u.calculateRideFare(distance)

	if distance == 0 {
		return model.RideData{}, pkgError.NewUnprocessableError(model.ErrMsgInvalidDistance)
	}

	now := time.Now()
	err = u.ridesRepo.UpdateRideData(ctx, model.UpdateRideDataRequest{
		RideID:   req.RideID,
		Status:   model.StatusNumRideEnded,
		Distance: &distance,
		Fare:     &fare,
		EndTime:  &now,
	})
	fmt.Println("cek", distance, ridePath)
	if err != nil {
		logger.Error(ctx, model.ErrMsgFailUpdateRideData, map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return model.RideData{}, pkgError.NewInternalServerError(model.ErrMsgFailUpdateRideData)
	}

	func() {
		err = u.ridesPubSub.BroadcastMessage(ctx, constants.TopicRideEnded, model.RideEndedMessage{
			RideID:   req.RideID,
			Distance: distance,
			Fare:     fare,
			RiderID:  rideData.RiderID,
		})
		if err != nil {
			logger.Error(ctx, "error broadcasting ride ended", map[string]interface{}{
				"driver_id": driverID,
				"error":     err,
			})
		}
	}()

	rideData.SetStatus(model.StatusNumRideEnded)
	rideData.SetDistance(distance)
	rideData.SetFare(fare)

	return rideData, nil
}

func calculateTotalDistance(path []pkgLocation.Coordinate) float64 {
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

func (u *ridesUsecase) calculateRideFare(distance float64) float64 {
	return math.Ceil(distance) * u.RidePricePerKm
}
