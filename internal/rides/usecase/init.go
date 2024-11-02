package usecase

import (
	"nebeng-jek/internal/rides/repository"
)

type ridesUsecase struct {
	locationRepo repository.RidesLocationRepository
	ridesRepo    repository.RidesRepository
	ridesPubSub  repository.RidesPubsubRepository
	paymentRepo  repository.PaymentRepository
}

func NewUsecase(
	locationRepo repository.RidesLocationRepository,
	ridesRepo repository.RidesRepository,
	ridesPubSub repository.RidesPubsubRepository,
	paymentRepo repository.PaymentRepository) RidesUsecase {

	return &ridesUsecase{
		locationRepo: locationRepo,
		ridesRepo:    ridesRepo,
		ridesPubSub:  ridesPubSub,
		paymentRepo:  paymentRepo,
	}
}
