package handler_nats

import "nebeng-jek/internal/rides/usecase"

type natsHandler struct {
	usecase usecase.RidesUsecase
}

func NewHandler(uc usecase.RidesUsecase) *natsHandler {
	return &natsHandler{
		usecase: uc,
	}
}
