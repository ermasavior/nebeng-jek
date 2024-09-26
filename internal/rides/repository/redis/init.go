package repository_redis

import (
	"nebeng-jek/internal/rides/repository"
	"nebeng-jek/pkg/redis"
)

type ridesRepo struct {
	cache redis.Collections
}

func NewRidesRepository(cache redis.Collections) repository.RidesRepository {
	return &ridesRepo{
		cache: cache,
	}
}
