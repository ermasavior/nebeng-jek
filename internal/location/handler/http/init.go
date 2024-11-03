package handler_http

import "nebeng-jek/internal/location/usecase"

type httpHandler struct {
	usecase usecase.LocationUsecase
}

func NewHandler(uc usecase.LocationUsecase) *httpHandler {
	return &httpHandler{
		usecase: uc,
	}
}
