package handler

import (
	"context"
	"nebeng-jek/internal/pkg/middleware"
	"nebeng-jek/pkg/amqp"
	"nebeng-jek/pkg/jwt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ridersHandler struct {
	upgrader websocket.Upgrader
	jwt      jwt.JWTInterface

	connStorage *sync.Map
}

func RegisterHandler(router *gin.RouterGroup, ridesChannel amqp.AMQPChannel) {
	h := &ridersHandler{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(r *http.Request) bool { return true },
		},
		jwt:         jwt.NewJWTGenerator(24*time.Hour, "PASSWORD"),
		connStorage: &sync.Map{},
	}

	mid := middleware.NewRidesMiddleware(h.jwt)

	router.GET("/ws/riders", mid.AuthJWTMiddleware, h.RiderWebsocket)

	go h.SubscribeDriverAcceptedRides(context.Background(), ridesChannel)
	go h.SubscribeReadyToPickupRides(context.Background(), ridesChannel)
}
