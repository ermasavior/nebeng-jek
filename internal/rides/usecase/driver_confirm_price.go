package usecase

import (
	"context"
	"nebeng-jek/internal/pkg/constants"
	pkgContext "nebeng-jek/internal/pkg/context"
	"nebeng-jek/internal/rides/model"
	"nebeng-jek/internal/rides/service/payment"
	pkgError "nebeng-jek/pkg/error"
	"nebeng-jek/pkg/logger"
)

func (u *ridesUsecase) DriverConfirmPrice(ctx context.Context, req model.DriverConfirmPriceRequest) (model.RideData, pkgError.AppError) {
	driverID := pkgContext.GetDriverIDFromContext(ctx)

	rideData, err := u.ridesRepo.GetRideData(ctx, req.RideID)
	if err == constants.ErrorDataNotFound {
		return model.RideData{}, pkgError.NewNotFoundError("ride is not found")
	}
	if err != nil {
		logger.Error(ctx, "error get ride data", map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return model.RideData{}, pkgError.NewInternalServerError("error get ride data")
	}

	if driverID != rideData.DriverID {
		return model.RideData{}, pkgError.NewForbiddenError("invalid ride id")
	}

	if rideData.Status != model.StatusRideEnded {
		return model.RideData{}, pkgError.NewBadRequestError("invalid ride status, must be RIDE_ENDED")
	}

	if rideData.Fare == nil {
		return model.RideData{}, pkgError.NewForbiddenError("invalid fare, must not be empty")
	}

	if req.CustomPrice > *rideData.Fare {
		return model.RideData{}, pkgError.NewBadRequestError("custom price must be lower than fare price")
	}

	driverMSISDN, err := u.ridesRepo.GetDriverMSISDNByID(ctx, rideData.DriverID)
	if err != nil {
		logger.Error(ctx, "error get driver msisdn", map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return model.RideData{}, pkgError.NewInternalServerError("error get driver msisdn")
	}

	riderMSISDN, err := u.ridesRepo.GetRiderMSISDNByID(ctx, rideData.RiderID)
	if err != nil {
		logger.Error(ctx, "error get rider driverID", map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return model.RideData{}, pkgError.NewInternalServerError("error get rider driverID")
	}

	var finalPrice = req.CustomPrice
	if req.CustomPrice > 0 && rideData.Fare != nil {
		finalPrice = *rideData.Fare
	}

	var distance = float64(0)
	if rideData.Distance != nil {
		distance = *rideData.Distance
	}

	err = u.handlePaymentTransaction(ctx, finalPrice, riderMSISDN, driverMSISDN)
	if err != nil {
		logger.Error(ctx, "error handle payment", map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return model.RideData{}, pkgError.NewInternalServerError("error handle payment")
	}

	err = u.ridesRepo.UpdateRideData(ctx, model.UpdateRideDataRequest{
		RideID:     rideData.RideID,
		FinalPrice: finalPrice,
		Status:     model.StatusNumRidePaid,
	})
	if err != nil {
		logger.Error(ctx, "error update ride by driver", map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return model.RideData{}, pkgError.NewInternalServerError("error update ride by driver")
	}

	err = u.ridesPubSub.BroadcastMessage(ctx, constants.TopicRidePaid, model.RidePaidMessage{
		RideID:     rideData.RideID,
		Distance:   distance,
		FinalPrice: finalPrice,
		RiderID:    rideData.RiderID,
	})
	if err != nil {
		logger.Error(ctx, "error broadcasting ride ended", map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return model.RideData{}, pkgError.NewInternalServerError("error broadcasting ride ended")
	}

	rideData.SetDistance(distance)
	rideData.SetFare(finalPrice)
	rideData.SetStatus(model.StatusNumRidePaid)

	return rideData, nil
}

func (u *ridesUsecase) handlePaymentTransaction(ctx context.Context, finalPrice float64, payerMSISDN, payeeMSISDN string) error {
	rideFee := finalPrice * model.RideFeeDiscount
	err := u.paymentService.DeductCredit(ctx, payment.DeductCreditRequest{
		MSISDN: payerMSISDN,
		Value:  finalPrice - rideFee,
	})
	if err != nil {
		logger.Error(ctx, "error deduct credit", map[string]interface{}{
			"payer": payerMSISDN,
			"payee": payeeMSISDN,
			"error": err,
		})
		return err
	}
	err = u.paymentService.AddCredit(ctx, payment.AddCreditRequest{
		MSISDN: payeeMSISDN,
		Value:  finalPrice - rideFee,
	})
	if err != nil {
		logger.Error(ctx, "error add credit", map[string]interface{}{
			"payer": payerMSISDN,
			"payee": payeeMSISDN,
			"error": err,
		})
		return err
	}
	err = u.paymentService.AddRevenue(ctx, payment.AddRevenueRequest{
		Value: rideFee,
	})
	if err != nil {
		logger.Error(ctx, "error add credit revenue", map[string]interface{}{
			"payer": payerMSISDN,
			"payee": payeeMSISDN,
			"error": err,
		})
		return err
	}

	return nil
}
