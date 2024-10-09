package handler

import (
	"nebeng-jek/internal/pkg/middleware"
	repo_amqp "nebeng-jek/internal/rides/repository/amqp"
	repo_db "nebeng-jek/internal/rides/repository/postgres"
	repo_redis "nebeng-jek/internal/rides/repository/redis"
	"nebeng-jek/internal/rides/usecase"
	"nebeng-jek/pkg/amqp"
	"nebeng-jek/pkg/jwt"
	"nebeng-jek/pkg/redis"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type ridesHandler struct {
	usecase usecase.RidesUsecase
}

func RegisterHandler(router *gin.RouterGroup, redis redis.Collections, db *sqlx.DB, ridesChannel amqp.AMQPChannel) {
	ridesPubSub := repo_amqp.NewRepository(ridesChannel)
	repoCache := repo_redis.NewRepository(redis)
	repoDB := repo_db.NewRepository(db)
	uc := usecase.NewUsecase(repoCache, repoDB, ridesPubSub)

	h := &ridesHandler{
		usecase: uc,
	}

	j := jwt.NewJWTGenerator(24*time.Hour, "PASSWORD")
	mid := middleware.NewRidesMiddleware(j)

	router.Use(mid.AuthJWTMiddleware)

	group := router.Group("/v1/drivers")
	{
		group.PUT("/availability", h.SetDriverAvailability)
		group.POST("/rides/confirm", h.ConfirmRideDriver)
		group.POST("/rides/start", h.StartRideDriver)
		group.POST("/rides/end", h.EndRideDriver)
	}

	group = router.Group("/v1/riders")
	{
		group.POST("/rides", h.CreateNewRide)
		group.POST("/rides/confirm", h.ConfirmRideRider)
	}
}
