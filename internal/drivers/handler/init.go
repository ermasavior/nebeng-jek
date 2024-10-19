package handler

import (
	"context"
	handler_http "nebeng-jek/internal/drivers/handler/http"
	handler_nats "nebeng-jek/internal/drivers/handler/nats"
	"nebeng-jek/internal/drivers/usecase"
	"nebeng-jek/internal/pkg/constants"
	"nebeng-jek/internal/pkg/middleware"
	nats_pkg "nebeng-jek/internal/pkg/nats"
	"nebeng-jek/pkg/jwt"
	"nebeng-jek/pkg/messaging/nats"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type RegisterHandlerParam struct {
	Router *gin.RouterGroup
	NatsJS nats.JetStreamConnection
	JWTGen jwt.JWTInterface
}

func RegisterHandler(reg RegisterHandlerParam) {
	repo := nats_pkg.NewPubsubRepository(reg.NatsJS)
	uc := usecase.NewDriverUsecase(repo)

	wsUpgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
	connStorage := &sync.Map{}

	httpHandler := handler_http.NewHandler(connStorage, wsUpgrader, uc)
	mid := middleware.NewRidesMiddleware(reg.JWTGen)
	reg.Router.GET("/ws/drivers", mid.AuthJWTMiddleware, httpHandler.DriverAllocationWebsocket)

	natsHandler := handler_nats.NewHandler(connStorage, uc)
	ctx := context.Background()
	go nats_pkg.SubscribeMessage(reg.NatsJS, constants.TopicRideNewRequest, natsHandler.SubscribeNewRideRequests(ctx))
	go nats_pkg.SubscribeMessage(reg.NatsJS, constants.TopicRideReadyToPickup, natsHandler.SubscribeReadyToPickupRides(ctx))
}
