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
