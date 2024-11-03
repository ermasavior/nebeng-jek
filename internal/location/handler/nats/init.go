package handler_nats

import "nebeng-jek/internal/location/usecase"

type natsHandler struct {
	usecase usecase.LocationUsecase
}

func NewHandler(uc usecase.LocationUsecase) *natsHandler {
	return &natsHandler{
		usecase: uc,
	}
}
