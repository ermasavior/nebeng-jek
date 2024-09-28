package handler

import (
	"context"
	constants "nebeng-jek/internal/pkg/constants/pubsub"
	"nebeng-jek/internal/rides/handler/middleware"
	repo_amqp "nebeng-jek/internal/rides/repository/amqp"
	repo_db "nebeng-jek/internal/rides/repository/postgres"
	repo_redis "nebeng-jek/internal/rides/repository/redis"
	"nebeng-jek/internal/rides/usecase"
	"nebeng-jek/pkg/amqp"
	"nebeng-jek/pkg/jwt"
	"nebeng-jek/pkg/logger"
	"nebeng-jek/pkg/redis"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/jmoiron/sqlx"
)

type ridesHandler struct {
	upgrader websocket.Upgrader
	usecase  usecase.RidesUsecase
	jwt      jwt.JWTInterface
}

func RegisterRidesHandler(router *gin.RouterGroup, redis redis.Collections, db *sqlx.DB, rideChannel amqp.AMQPChannel) {
	err := rideChannel.ExchangeDeclare(
		constants.RideRequestsFanout,
		"fanout", // exchange type: fanout
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		logger.Fatal(context.Background(), "failed to declare an amqp exchange", map[string]interface{}{
			"error": err,
		})
	}

	ridesPubSub := repo_amqp.NewRidesRepository(rideChannel)
	repoCache := repo_redis.NewRidesRepository(redis)
	repoDB := repo_db.NewRidesRepository(db)
	uc := usecase.NewRidesUsecase(repoCache, repoDB, ridesPubSub)

	h := &ridesHandler{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(r *http.Request) bool { return true },
		},
		jwt:     jwt.NewJWTGenerator(24*time.Hour, "PASSWORD"),
		usecase: uc,
	}

	mid := middleware.NewRidesMiddleware(h.jwt)

	// router.GET("/dummy/login", h.LoginHandler)

	router.Use(mid.LoginDriverMiddleware)

	// router.GET("/ws/drivers", h.DriverAllocation)
	router.PUT("/drivers/availability", h.SetDriverAvailability)

	router.POST("/riders/rides", h.CreateNewRide)
}
