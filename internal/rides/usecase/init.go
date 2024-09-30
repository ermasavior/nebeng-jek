package usecase

import (
	"nebeng-jek/internal/rides/repository"
)

type ridesUsecase struct {
	locationRepo repository.RidesLocationRepository
	ridesRepo    repository.RidesRepository
	ridesPubSub  repository.RidesPubsubRepository
}

func NewUsecase(
	locationRepo repository.RidesLocationRepository,
	ridesRepo repository.RidesRepository,
	ridesPubSub repository.RidesPubsubRepository) RidesUsecase {

	return &ridesUsecase{
		locationRepo: locationRepo,
		ridesRepo:    ridesRepo,
		ridesPubSub:  ridesPubSub,
	}
}
