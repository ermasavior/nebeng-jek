package usecase

import (
	"context"
	"nebeng-jek/internal/rides/model"
	pkgError "nebeng-jek/pkg/error"
)

func (u *ridesUsecase) TrackUserLocation(ctx context.Context, req model.TrackUserLocationRequest) *pkgError.AppError {
	err := u.locationRepo.TrackUserLocation(ctx, req)
	if err != nil {
		return pkgError.NewInternalServerError(err, "error add user location track")
	}
	return nil
}
