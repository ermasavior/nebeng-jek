package handler

import (
	"context"
	"nebeng-jek/internal/pkg/constants"
	"nebeng-jek/internal/pkg/middleware"
	nats_pkg "nebeng-jek/internal/pkg/nats"
	handler_http "nebeng-jek/internal/rides/handler/http"
	handler_nats "nebeng-jek/internal/rides/handler/nats"
	repo_db "nebeng-jek/internal/rides/repository/postgres"
	repo_redis "nebeng-jek/internal/rides/repository/redis"
	"nebeng-jek/internal/rides/service/payment"
	"nebeng-jek/internal/rides/usecase"
	"nebeng-jek/pkg/jwt"
	"nebeng-jek/pkg/messaging/nats"
	"nebeng-jek/pkg/redis"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type RegisterHandlerParam struct {
	Router *gin.RouterGroup
	Redis  redis.Collections
	DB     *sqlx.DB
	NatsJS nats.JetStreamConnection
	JWTGen jwt.JWTInterface
}

func RegisterHandler(reg RegisterHandlerParam) {
	ridesPubSub := nats_pkg.NewPubsubRepository(reg.NatsJS)
	repoCache := repo_redis.NewRepository(reg.Redis)
	repoDB := repo_db.NewRepository(reg.DB)
	paymentSvc := payment.NewPaymentService()
	uc := usecase.NewUsecase(repoCache, repoDB, ridesPubSub, paymentSvc)

	httpHandler := handler_http.NewHandler(uc)
	mid := middleware.NewRidesMiddleware(reg.JWTGen)
	reg.Router.Use(mid.AuthJWTMiddleware)

	group := reg.Router.Group("/drivers")
	{
		group.PUT("/availability", httpHandler.SetDriverAvailability)
		group.POST("/rides/confirm", httpHandler.ConfirmRideDriver)
		group.POST("/rides/start", httpHandler.StartRideDriver)
		group.POST("/rides/end", httpHandler.EndRideDriver)
		group.POST("/rides/confirm-payment", httpHandler.ConfirmPaymentDriver)
	}

	group = reg.Router.Group("/riders")
	{
		group.POST("/rides", httpHandler.CreateNewRide)
		group.POST("/rides/confirm", httpHandler.ConfirmRideRider)
	}

	natsHandler := handler_nats.NewHandler(uc)
	ctx := context.Background()
	go nats_pkg.SubscribeMessage(reg.NatsJS, constants.TopicUserLiveLocation, natsHandler.SubscribeUserLiveLocation(ctx))
}
