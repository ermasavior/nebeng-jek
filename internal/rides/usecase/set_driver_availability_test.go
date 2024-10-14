package usecase

import (
	"context"
	"errors"
	"testing"

	pkgContext "nebeng-jek/internal/pkg/context"
	"nebeng-jek/internal/rides/model"
	mockRepo "nebeng-jek/mock/repository"
	pkgError "nebeng-jek/pkg/error"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUsecase_SetDriverAvailability(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mockRepo.NewMockRidesLocationRepository(ctrl)
	usecaseMock := NewUsecase(repoMock, nil, nil, nil)

	var (
		msisdn = "0811111"

		req = model.SetDriverAvailabilityRequest{
			CurrentLocation: model.Coordinate{
				Longitude: 11,
				Latitude:  11,
			},
		}
	)

	ctx := context.Background()
	ctx = pkgContext.SetMSISDNToContext(ctx, msisdn)

	t.Run("success - should add available driver", func(t *testing.T) {
		req.IsAvailable = true
		repoMock.EXPECT().AddAvailableDriver(ctx, msisdn, req.CurrentLocation).
			Return(nil)

		err := usecaseMock.SetDriverAvailability(ctx, req)
		assert.Nil(t, err)
	})

	t.Run("success - should remove available driver", func(t *testing.T) {
		req.IsAvailable = false
		repoMock.EXPECT().RemoveAvailableDriver(ctx, msisdn).
			Return(nil)

		err := usecaseMock.SetDriverAvailability(ctx, req)
		assert.Nil(t, err)
	})

	t.Run("failed - should return error when add available driver failed", func(t *testing.T) {
		req.IsAvailable = true
		expectedErr := errors.New("error from repo")

		repoMock.EXPECT().AddAvailableDriver(ctx, msisdn, req.CurrentLocation).
			Return(expectedErr)

		err := usecaseMock.SetDriverAvailability(ctx, req)
		assert.Equal(t, pkgError.NewInternalServerError(expectedErr, "error adding available driver"), err)
	})

	t.Run("failed - should return error when remove available driver failed", func(t *testing.T) {
		req.IsAvailable = false
		expectedErr := errors.New("error from repo")

		repoMock.EXPECT().RemoveAvailableDriver(ctx, msisdn).
			Return(expectedErr)

		err := usecaseMock.SetDriverAvailability(ctx, req)
		assert.Equal(t, pkgError.NewInternalServerError(expectedErr, "error removing available driver"), err)
	})
}
