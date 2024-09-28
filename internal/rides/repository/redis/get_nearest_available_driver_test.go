package repository_redis

import (
	"context"
	"testing"

	"nebeng-jek/internal/rides/model"
	mockRedis "nebeng-jek/mock/pkg/redis"

	"github.com/go-redis/redis/v8"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRepository_GetNearestAvailableDrivers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	redisMock := mockRedis.NewMockCollections(ctrl)
	repositoryMock := NewRidesRepository(redisMock)

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
