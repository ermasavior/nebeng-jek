package repository_redis

import (
	"context"
	"testing"

	pkgContext "nebeng-jek/internal/pkg/context"
	"nebeng-jek/internal/rides/model"
	mockRedis "nebeng-jek/mock/pkg/redis"

	"github.com/go-redis/redis/v8"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRepository_AddAvailableDriver(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	redisMock := mockRedis.NewMockCollections(ctrl)
	repositoryMock := NewRidesRepository(redisMock)

	var (
		msisdn   = "0811111"
		location = model.Coordinate{
			Longitude: 11,
			Latitude:  11,
		}
	)

	ctx := context.Background()
	ctx = pkgContext.SetMSISDNToContext(ctx, msisdn)

	t.Run("success - should execute redis GEOADD", func(t *testing.T) {
		res := &redis.IntCmd{}
		redisMock.EXPECT().GeoAdd(ctx, model.KeyAvailableDrivers, &redis.GeoLocation{
			Name:      msisdn,
			Longitude: location.Longitude,
			Latitude:  location.Latitude,
		}).Return(res)

		err := repositoryMock.AddAvailableDriver(ctx, msisdn, location)
		assert.Nil(t, err)
	})

	t.Run("failed - should return error when GEOADD returns error", func(t *testing.T) {
		res := &redis.IntCmd{}
		res.SetErr(redis.ErrClosed)
		redisMock.EXPECT().GeoAdd(ctx, model.KeyAvailableDrivers, &redis.GeoLocation{
			Name:      msisdn,
			Longitude: location.Longitude,
			Latitude:  location.Latitude,
		}).Return(res)

		err := repositoryMock.AddAvailableDriver(ctx, msisdn, location)
		assert.EqualError(t, err, redis.ErrClosed.Error())
	})
}
