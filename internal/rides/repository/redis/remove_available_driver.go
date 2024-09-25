package repository_redis

import (
	"context"
	"nebeng-jek/internal/rides/model"
)

func (r *ridesRepo) RemoveAvailableDriver(ctx context.Context, msisdn string) error {
	return r.cache.ZRem(ctx, model.KeyAvailableDrivers, model.SetDriverKey(msisdn)).Err()
}
