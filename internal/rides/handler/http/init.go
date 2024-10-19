package handler_http

import "nebeng-jek/internal/rides/usecase"

type httpHandler struct {
	usecase usecase.RidesUsecase
}

func NewHandler(uc usecase.RidesUsecase) *httpHandler {
	return &httpHandler{
		usecase: uc,
	}
}
