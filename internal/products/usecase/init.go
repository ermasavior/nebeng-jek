package usecase

import (
	repository "nebeng-jek/internal/products/repository"
	"nebeng-jek/pkg/redis"
)

func NewProductUsecase(repository repository.ProductRepository, redis redis.Collections) ProductUsecase {
	return &productUsecase{
		repository: repository,
		redis:      redis,
	}
}
