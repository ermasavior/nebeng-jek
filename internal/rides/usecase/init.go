package usecase

import (
	"nebeng-jek/internal/rides/repository"
	"nebeng-jek/pkg/configs"
)

type ridesUsecase struct {
	locationRepo repository.RidesLocationRepository
	ridesRepo    repository.RidesRepository
	ridesPubSub  repository.RidesPubsubRepository
	paymentRepo  repository.PaymentRepository

	RidePricePerKm    float64
	RideFeePercentage int
}

func NewUsecase(cfg *configs.Config,
	locationRepo repository.RidesLocationRepository,
	ridesRepo repository.RidesRepository,
	ridesPubSub repository.RidesPubsubRepository,
	paymentRepo repository.PaymentRepository) RidesUsecase {

	return &ridesUsecase{
		locationRepo:      locationRepo,
		ridesRepo:         ridesRepo,
		ridesPubSub:       ridesPubSub,
		paymentRepo:       paymentRepo,
		RidePricePerKm:    cfg.RidePricePerKm,
		RideFeePercentage: cfg.RideFeePercentage,
	}
}
