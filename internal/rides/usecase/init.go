package usecase

import (
	"nebeng-jek/internal/rides/repository"
	"nebeng-jek/internal/rides/service/payment"
)

type ridesUsecase struct {
	locationRepo   repository.RidesLocationRepository
	ridesRepo      repository.RidesRepository
	ridesPubSub    repository.RidesPubsubRepository
	paymentService payment.PaymentService
}

func NewUsecase(
	locationRepo repository.RidesLocationRepository,
	ridesRepo repository.RidesRepository,
	ridesPubSub repository.RidesPubsubRepository,
	paymentService payment.PaymentService) RidesUsecase {

	return &ridesUsecase{
		locationRepo:   locationRepo,
		ridesRepo:      ridesRepo,
		ridesPubSub:    ridesPubSub,
		paymentService: paymentService,
	}
}
