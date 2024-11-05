package handler

import (
	"context"
	handler_http "nebeng-jek/internal/location/handler/http"
	handler_nats "nebeng-jek/internal/location/handler/nats"
	"nebeng-jek/internal/location/usecase"
	"nebeng-jek/internal/pkg/constants"
	"nebeng-jek/internal/pkg/middleware"
	nats_pkg "nebeng-jek/internal/pkg/nats"
	"nebeng-jek/pkg/configs"
	"nebeng-jek/pkg/messaging/nats"
	"nebeng-jek/pkg/redis"

	"github.com/gin-gonic/gin"
)

type RegisterHandlerParam struct {
	Router *gin.RouterGroup
	Redis  redis.Collections
	NatsJS nats.JetStreamConnection
	Cfg    *configs.Config
}

func RegisterHandler(ctx context.Context, reg RegisterHandlerParam) {
	uc := usecase.NewLocationUsecase(reg.Redis)
	httpHandler := handler_http.NewHandler(uc)

	group := reg.Router.Group("/drivers")
	group.Use(middleware.InternalAuthorization(reg.Cfg.InternalAPIKey))
	{
		group.POST("/available", httpHandler.AddAvailableDriver)
		group.DELETE("/available/:driver_id", httpHandler.RemoveAvailableDriver)
		group.GET("/available/nearby", httpHandler.GetNearestAvailableDrivers)
		group.GET("/ride-path", httpHandler.GetRidePath)
	}

	natsHandler := handler_nats.NewHandler(uc)
	go nats_pkg.SubscribeMessage(reg.NatsJS, constants.TopicUserLiveLocation, natsHandler.SubscribeUserLiveLocation(ctx), "consumer_live_location")
}
