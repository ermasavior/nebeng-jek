package handler

import (
	"nebeng-jek/internal/rides/handler/middleware"
	repo_redis "nebeng-jek/internal/rides/repository/redis"
	"nebeng-jek/internal/rides/usecase"
	"nebeng-jek/pkg/jwt"
	"nebeng-jek/pkg/redis"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ridesHandler struct {
	upgrader websocket.Upgrader
	Usecase  usecase.RidesUsecase
}

var j = jwt.NewJWTGenerator(24*time.Hour, "PASSWORD")

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

	mid := middleware.NewRidesMiddleware(j)

	router.GET("/dummy/login", h.LoginHandler)

	router.Use(mid.LoginDriverMiddleware)

	// router.GET("/ws/drivers", h.DriverAllocation)
	router.PUT("/drivers/availability", h.SetDriverAvailability)
}
