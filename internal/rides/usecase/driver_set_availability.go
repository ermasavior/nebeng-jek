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

	if req.IsAvailable {
		err := u.locationRepo.AddAvailableDriver(ctx, driverID, req.CurrentLocation)
		if err != nil {
			logger.Error(ctx, "error adding available driver", map[string]interface{}{
				"driver_id": driverID,
				"error":     err,
			})
			return pkgError.NewInternalServerError("error adding available driver")
		}
		return nil
	}

	err = u.locationRepo.RemoveAvailableDriver(ctx, driverID)
	if err != nil {
		logger.Error(ctx, "error removing available driver", map[string]interface{}{
			"driver_id": driverID,
			"error":     err,
		})
		return pkgError.NewInternalServerError("error removing available driver")
	}

	return nil
}
