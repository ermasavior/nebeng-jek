package usecase

import (
	"context"
	pkgContext "nebeng-jek/internal/pkg/context"
	"nebeng-jek/internal/rides/model"
	pkgError "nebeng-jek/pkg/error"
	"nebeng-jek/pkg/logger"
)

func (u *ridesUsecase) SetDriverAvailability(ctx context.Context, req model.SetDriverAvailabilityRequest) error {
	msisdn := pkgContext.GetMSISDNFromContext(ctx)
	if req.IsAvailable {
		err := u.Repo.AddAvailableDriver(ctx, msisdn, req.CurrentLocation)
		if err != nil {
			logger.Error(ctx, "error adding available driver", map[string]interface{}{
				"msisdn": msisdn,
			})
			return pkgError.NewInternalServerError(err, "error adding available driver")
		}
		return nil
	}

	err := u.Repo.RemoveAvailableDriver(ctx, msisdn)
	if err != nil {
		logger.Error(ctx, "error removing available driver", map[string]interface{}{
			"msisdn": msisdn,
		})
		return pkgError.NewInternalServerError(err, "error removing available driver")
	}
	return nil
}
