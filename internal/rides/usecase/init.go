package usecase

import "nebeng-jek/internal/rides/repository"

type ridesUsecase struct {
	Repo repository.RidesRepository
}

func NewRidesUsecase(repo repository.RidesRepository) RidesUsecase {
	return &ridesUsecase{
		Repo: repo,
	}
}
