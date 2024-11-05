package usecase

import (
	"context"
	"nebeng-jek/internal/pkg/constants"
	pkgContext "nebeng-jek/internal/pkg/context"
	"nebeng-jek/internal/rides/model"
	pkgError "nebeng-jek/pkg/error"
	"nebeng-jek/pkg/logger"
)

func (u *ridesUsecase) DriverSetAvailability(ctx context.Context, req model.DriverSetAvailabilityRequest) pkgError.AppError {
	driverID := pkgContext.GetDriverIDFromContext(ctx)

	_, err := u.ridesRepo.GetDriverMSISDNByID(ctx, driverID)
	if err == constants.ErrorDataNotFound {
		return pkgError.NewUnauthorizedError("invalid driver id")
	}

	var status int

	if req.IsAvailable {
		status = model.StatusDriverAvailable
		err := u.locationRepo.AddAvailableDriver(ctx, driverID, req.CurrentLocation)
		if err != nil {
			logger.Error(ctx, "error adding available driver", map[string]interface{}{
				"driver_id": driverID,
				"error":     err,
			})
			return pkgError.NewInternalServerError("error adding available driver")
		}
	} else {
		status = model.StatusDriverOff
		err = u.locationRepo.RemoveAvailableDriver(ctx, driverID)
		if err != nil {
			logger.Error(ctx, model.ErrMsgFailRemoveAvailableDriver, map[string]interface{}{
				"driver_id": driverID,
				"error":     err,
			})
			return pkgError.NewInternalServerError(model.ErrMsgFailRemoveAvailableDriver)
		}
	}

	err = u.ridesRepo.UpdateDriverStatus(ctx, model.UpdateDriverStatusRequest{
		DriverID: driverID,
		Status:   status,
	})
	if err != nil {
		logger.Error(ctx, model.ErrMsgFailUpdateStatusDriver, map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return pkgError.NewInternalServerError(model.ErrMsgFailUpdateStatusDriver)
	}

	return nil
}
