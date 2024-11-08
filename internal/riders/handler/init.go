package handler

import (
	"context"
	"nebeng-jek/internal/pkg/constants"
	"nebeng-jek/internal/pkg/middleware"
	nats_pkg "nebeng-jek/internal/pkg/nats"
	handler_http "nebeng-jek/internal/riders/handler/http"
	handler_nats "nebeng-jek/internal/riders/handler/nats"
	"nebeng-jek/internal/riders/usecase"
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

func RegisterHandler(ctx context.Context, reg RegisterHandlerParam) {
	repo := nats_pkg.NewPubsubRepository(reg.NatsJS)
	uc := usecase.NewRiderUsecase(repo)
	wsUpgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
	connStorage := &sync.Map{}

	httpHandler := handler_http.NewHandler(connStorage, wsUpgrader, uc)

	mid := middleware.NewRidesMiddleware(reg.JWTGen)
	reg.Router.GET("/ws/riders", mid.RiderAuthMiddleware, httpHandler.RiderWebsocket)

	natsHandler := handler_nats.NewHandler(connStorage)

	go nats_pkg.SubscribeMessage(reg.NatsJS, constants.TopicRideMatchedDriver, natsHandler.SubscribeRideMatchedDriver(ctx))
	go nats_pkg.SubscribeMessage(reg.NatsJS, constants.TopicRideStarted, natsHandler.SubscribeRideStarted(ctx))
	go nats_pkg.SubscribeMessage(reg.NatsJS, constants.TopicRideEnded, natsHandler.SubscribeRideEnded(ctx))
	go nats_pkg.SubscribeMessage(reg.NatsJS, constants.TopicRidePaid, natsHandler.SubscribeRidePaid(ctx))
}
