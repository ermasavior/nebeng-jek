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

func (u *ridesUsecase) ConfirmPaymentDriver(ctx context.Context, req model.ConfirmPaymentDriverRequest) (model.RideData, *pkgError.AppError) {
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

	rideData, err := u.ridesRepo.GetRideData(ctx, req.RideID)
	if err == constants.ErrorDataNotFound {
		return model.RideData{}, pkgError.NewNotFound(err, "ride is not found")
	}
	if err != nil {
		logger.Error(ctx, "error get ride data", map[string]interface{}{
			"msisdn": msisdn,
			"error":  err,
		})
		return model.RideData{}, pkgError.NewInternalServerError(err, "error get ride data")
	}

	if driver.ID != rideData.DriverID {
		return model.RideData{}, pkgError.NewForbidden(err, "invalid ride id")
	}

	// TBD
	// if rideData.Status != model.StatusNumRideEnded {
	// 	return model.RideData{}, pkgError.NewBadRequest(nil, "invalid ride status, ride must be ended")
	// }

	if rideData.Fare == nil {
		return model.RideData{}, pkgError.NewForbidden(nil, "invalid fare, must not be empty")
	}

	if req.CustomPrice > *rideData.Fare {
		return model.RideData{}, pkgError.NewBadRequest(nil, "custom price must be lower than fare price")
	}

	var finalPrice = req.CustomPrice
	if req.CustomPrice > 0 {
		finalPrice = *rideData.Fare
	}

	riderMSISDN, err := u.ridesRepo.GetRiderMSISDNByID(ctx, rideData.RiderID)
	if err != nil {
		logger.Error(ctx, "error get rider msisdn", map[string]interface{}{
			"msisdn": msisdn,
			"error":  err,
		})
		return model.RideData{}, pkgError.NewInternalServerError(err, "error get rider msisdn")
	}

	err = u.handlePaymentTransaction(ctx, finalPrice, riderMSISDN, driver.MSISDN)
	if err != nil {
		logger.Error(ctx, "error handle payment", map[string]interface{}{
			"msisdn": msisdn,
			"error":  err,
		})
		return model.RideData{}, pkgError.NewInternalServerError(err, "error handle payment")
	}

	err = u.ridesRepo.UpdateRideData(ctx, model.UpdateRideDataRequest{
		RideID:     rideData.RideID,
		FinalPrice: finalPrice,
		Status:     model.StatusNumRidePaid,
	})
	if err != nil {
		logger.Error(ctx, "error update ride by driver", map[string]interface{}{
			"msisdn": msisdn,
			"error":  err,
		})
		return model.RideData{}, pkgError.NewInternalServerError(err, "error update ride by driver")
	}

	err = u.ridesPubSub.BroadcastMessage(ctx, constants.RidePaidExchange, model.RidePaidMessage{
		RideID:      rideData.RideID,
		Distance:    *rideData.Distance,
		FinalPrice:  finalPrice,
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
