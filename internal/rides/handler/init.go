package handler

import (
	"nebeng-jek/internal/rides/handler/middleware"
	repo_redis "nebeng-jek/internal/rides/repository/redis"
	"nebeng-jek/internal/rides/usecase"
	"nebeng-jek/pkg/redis"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ridesHandler struct {
	upgrader websocket.Upgrader
	Usecase  usecase.RidesUsecase
}

func RegisterRidesHandler(router *gin.RouterGroup, redis redis.Collections) {
	pRepository := repo_redis.NewRidesRepository(redis)
	pUsecase := usecase.NewRidesUsecase(pRepository)

	h := &ridesHandler{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		Usecase: pUsecase,
	}

	router.GET("/dummy/login", LoginHandler)

	router.Use(middleware.LoginDriverMiddleware)

	router.GET("/ws/drivers", h.DriverAllocation)
	router.PUT("/drivers/availability", h.SetDriverAvailability)
}
