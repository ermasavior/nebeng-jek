package usecase

import (
	"context"
	"nebeng-jek/internal/pkg/constants"
	pkgContext "nebeng-jek/internal/pkg/context"
	"nebeng-jek/internal/rides/model"
	pkgError "nebeng-jek/pkg/error"
	"nebeng-jek/pkg/logger"
)

func (u *ridesUsecase) DriverConfirmPayment(ctx context.Context, req model.DriverConfirmPaymentRequest) (model.RideData, pkgError.AppError) {
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

	if err := model.ValidateConfirmPayment(rideData, driverID, req.CustomPrice); err != nil {
		return model.RideData{}, err
	}

	var finalPrice = *rideData.Fare
	if req.CustomPrice > 0 {
		finalPrice = req.CustomPrice
	}

	var distance = float64(0)
	if rideData.Distance != nil {
		distance = *rideData.Distance
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

	err = u.processPayment(ctx, req.RideID, finalPrice, riderMSISDN, driverMSISDN)
	if err != nil {
		logger.Error(ctx, model.ErrMsgFailProcessPayment, map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return model.RideData{}, pkgError.NewInternalServerError(model.ErrMsgFailProcessPayment)
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

	func() {
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
		}
	}()

	rideData.SetDistance(distance)
	rideData.SetFinalPrice(finalPrice)
	rideData.SetStatus(model.StatusNumRidePaid)

	return rideData, nil
}

func (u *ridesUsecase) processPayment(ctx context.Context, rideID int64, finalPrice float64, payerMSISDN, payeeMSISDN string) error {
	commission := finalPrice * model.RideFeeDiscount
	netPrice := finalPrice - commission

	err := u.paymentRepo.DeductCredit(ctx, model.DeductCreditRequest{
		MSISDN: payerMSISDN,
		Value:  netPrice,
	})
	if err != nil {
		logger.Error(ctx, "error deduct credit", map[string]interface{}{
			"payer": payerMSISDN, "payee": payeeMSISDN, "error": err,
		})
		return err
	}
	err = u.paymentRepo.AddCredit(ctx, model.AddCreditRequest{
		MSISDN: payeeMSISDN,
		Value:  netPrice,
	})
	if err != nil {
		logger.Error(ctx, "error add credit", map[string]interface{}{
			"payer": payerMSISDN, "payee": payeeMSISDN, "error": err,
		})
		return err
	}
	err = u.ridesRepo.StoreRideCommission(ctx, model.StoreRideCommissionRequest{
		RideID:     rideID,
		Commission: commission,
	})
	if err != nil {
		logger.Error(ctx, "error store ride commission", map[string]interface{}{
			"payer": payerMSISDN, "payee": payeeMSISDN, "error": err,
		})
		return err
	}

	return nil
}
