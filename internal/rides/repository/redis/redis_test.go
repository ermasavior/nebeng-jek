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
	repositoryMock := NewRepository(redisMock)

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

func TestRepository_GetNearestAvailableDrivers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	redisMock := mockRedis.NewMockCollections(ctrl)
	repositoryMock := NewRepository(redisMock)

	var (
		location = model.Coordinate{
			Longitude: 11,
			Latitude:  11,
		}
	)

	ctx := context.Background()

	t.Run("success - should execute redis GeoRadius", func(t *testing.T) {
		res := &redis.GeoLocationCmd{}
		res.SetVal([]redis.GeoLocation{
			{Name: "0123"}, {Name: "0456"}, {Name: "0789"},
		})
		redisMock.EXPECT().GeoRadius(ctx, model.KeyAvailableDrivers,
			location.Longitude, location.Latitude, &redis.GeoRadiusQuery{
				Radius:   model.NearestRadius,
				Unit:     model.NearestRadiusUnit,
				WithDist: true,
			}).Return(res)

		actual, err := repositoryMock.GetNearestAvailableDrivers(ctx, location)
		assert.Nil(t, err)
		assert.Equal(t, []string{"0123", "0456", "0789"}, actual)
	})

	t.Run("failed - should return error when GeoRadius returns error", func(t *testing.T) {
		res := &redis.GeoLocationCmd{}
		res.SetErr(redis.ErrClosed)
		redisMock.EXPECT().GeoRadius(ctx, model.KeyAvailableDrivers,
			location.Longitude, location.Latitude, &redis.GeoRadiusQuery{
				Radius:   model.NearestRadius,
				Unit:     model.NearestRadiusUnit,
				WithDist: true,
			}).Return(res)

		_, err := repositoryMock.GetNearestAvailableDrivers(ctx, location)
		assert.EqualError(t, err, redis.ErrClosed.Error())
	})
}

func TestRepository_RemoveAvailableDriver(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	redisMock := mockRedis.NewMockCollections(ctrl)
	repositoryMock := ridesRepo{
		cache: redisMock,
	}

	var (
		msisdn = "0811111"
	)

	ctx := context.Background()
	ctx = pkgContext.SetMSISDNToContext(ctx, msisdn)

	t.Run("success - should execute redis ZREM", func(t *testing.T) {
		res := &redis.IntCmd{}
		redisMock.EXPECT().ZRem(ctx, model.KeyAvailableDrivers, msisdn).Return(res)

		err := repositoryMock.RemoveAvailableDriver(ctx, msisdn)
		assert.Nil(t, err)
	})

	t.Run("failed - should return error when ZREM returns error", func(t *testing.T) {
		res := &redis.IntCmd{}
		res.SetErr(redis.ErrClosed)
		redisMock.EXPECT().ZRem(ctx, model.KeyAvailableDrivers, msisdn).Return(res)

		err := repositoryMock.RemoveAvailableDriver(ctx, msisdn)
		assert.EqualError(t, err, redis.ErrClosed.Error())
	})
}
