package handler

import (
	"context"
	"nebeng-jek/internal/pkg/middleware"
	nats_pkg "nebeng-jek/internal/pkg/nats"
	handler_http "nebeng-jek/internal/rides/handler/http"
	location "nebeng-jek/internal/rides/repository/external_api/location"
	"nebeng-jek/internal/rides/repository/external_api/payment"
	repo_db "nebeng-jek/internal/rides/repository/postgres"
	"nebeng-jek/internal/rides/usecase"
	"nebeng-jek/pkg/configs"
	"nebeng-jek/pkg/jwt"
	"nebeng-jek/pkg/messaging/nats"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type RegisterHandlerParam struct {
	Router     *gin.RouterGroup
	DB         *sqlx.DB
	NatsJS     nats.JetStreamConnection
	JWTGen     jwt.JWTInterface
	Cfg        *configs.Config
	HttpClient *http.Client
}

func RegisterHandler(ctx context.Context, reg RegisterHandlerParam) {
	ridesPubSub := nats_pkg.NewPubsubRepository(reg.NatsJS)
	repoDB := repo_db.NewRepository(reg.DB)
	repoCache := location.NewLocationRepository(reg.Cfg, reg.HttpClient)
	paymentSvc := payment.NewPaymentRepository(reg.Cfg, reg.HttpClient)

	uc := usecase.NewUsecase(repoCache, repoDB, ridesPubSub, paymentSvc)

	httpHandler := handler_http.NewHandler(uc)
	mid := middleware.NewRidesMiddleware(reg.JWTGen)

	group := reg.Router.Group("/drivers")
	group.Use(mid.DriverAuthMiddleware)
	{
		group.PATCH("/availability", httpHandler.DriverSetAvailability)
		group.PATCH("/ride/confirm", httpHandler.DriverConfirmRide)
		group.PATCH("/ride/start", httpHandler.DriverStartRide)
		group.PATCH("/ride/end", httpHandler.DriverEndRide)
		group.PATCH("/ride/confirm-payment", httpHandler.DriverConfirmPayment)
	}

	group = reg.Router.Group("/riders")
	group.Use(mid.RiderAuthMiddleware)
	{
		group.POST("/ride/create", httpHandler.RiderCreateNewRide)
		group.GET("/ride/:ride_id", httpHandler.GetRideData)
		group.PATCH("/ride/confirm", httpHandler.RiderConfirmRide)
	}
}
