package usecase

import (
	"context"
	"nebeng-jek/internal/rides/model"
	pkgError "nebeng-jek/pkg/error"
	"nebeng-jek/pkg/logger"
)

func (u *ridesUsecase) TrackUserLocation(ctx context.Context, req model.TrackUserLocationRequest) pkgError.AppError {
	err := u.locationRepo.TrackUserLocation(ctx, req)
	if err != nil {
		logger.Error(ctx, "error track user location", map[string]interface{}{
			"req":           req,
			logger.ErrorKey: err,
		})
		return pkgError.NewInternalServerError("error track user location")
	}
	return nil
}
