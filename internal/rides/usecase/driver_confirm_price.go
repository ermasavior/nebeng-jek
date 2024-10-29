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
		return model.RideData{}, pkgError.NewNotFoundError(pkgError.ErrResourceNotFoundMsg)
	}
	if err != nil {
		logger.Error(ctx, model.ErrMsgFailGetRideData, map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return model.RideData{}, pkgError.NewInternalServerError(model.ErrMsgFailGetRideData)
	}

	if rideData.DriverID == nil || *rideData.DriverID != driverID {
		return model.RideData{}, pkgError.NewForbiddenError((pkgError.ErrForbiddenMsg))
	}

	if rideData.StatusNum != model.StatusNumRideEnded {
		return model.RideData{}, pkgError.NewBadRequestError(model.ErrMsgInvalidRideStatus)
	}

	if rideData.Fare == nil {
		return model.RideData{}, pkgError.NewForbiddenError("invalid fare, must not be empty")
	}

	if req.CustomPrice > *rideData.Fare {
		return model.RideData{}, pkgError.NewBadRequestError("custom price must be lower than fare price")
	}

	driverMSISDN, err := u.ridesRepo.GetDriverMSISDNByID(ctx, *rideData.DriverID)
	if err != nil {
		logger.Error(ctx, model.ErrMsgFailGetDriverMSISDN, map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return model.RideData{}, pkgError.NewInternalServerError(model.ErrMsgFailGetDriverMSISDN)
	}

	riderMSISDN, err := u.ridesRepo.GetRiderMSISDNByID(ctx, rideData.RiderID)
	if err != nil {
		logger.Error(ctx, model.ErrMsgFailGetRiderMSISDN, map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return model.RideData{}, pkgError.NewInternalServerError(model.ErrMsgFailGetRiderMSISDN)
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
		FinalPrice: &finalPrice,
		Status:     model.StatusNumRidePaid,
	})
	if err != nil {
		logger.Error(ctx, model.ErrMsgFailUpdateRideData, map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return model.RideData{}, pkgError.NewInternalServerError(model.ErrMsgFailUpdateRideData)
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
	rideData.SetFinalPrice(finalPrice)
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
