package handler

import (
	"context"
	"fmt"
	"nebeng-jek/internal/pkg/constants"
	mock_nats "nebeng-jek/mock/pkg/messaging/nats"
	mock_redis "nebeng-jek/mock/pkg/redis"
	"nebeng-jek/pkg/configs"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRegisterHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	natsConn := mock_nats.NewMockJetStreamConnection(ctrl)
	natsConn.EXPECT().Subscribe(constants.TopicUserLiveLocation, gomock.Any(), gomock.Any()).AnyTimes()

	router := gin.New()

	reg := RegisterHandlerParam{
		Router:     &router.RouterGroup,
		Redis:      mock_redis.NewMockCollections(ctrl),
		DB:         nil, // no tests
		NatsJS:     natsConn,
		JWTGen:     nil, // no tests
		Cfg:        configs.NewMockConfig(),
		HttpClient: http.DefaultClient,
	}
	RegisterHandler(context.Background(), reg)

	expectedRoutes := map[string]gin.RouteInfo{
		"PUT:/drivers/availability": {
			Method:  "PUT",
			Path:    "/drivers/availability",
			Handler: "nebeng-jek/internal/rides/handler/http.(*httpHandler).DriverSetAvailability-fm",
		},
		"POST:/drivers/ride/confirm": {
			Method:  "POST",
			Path:    "/drivers/ride/confirm",
			Handler: "nebeng-jek/internal/rides/handler/http.(*httpHandler).DriverConfirmRide-fm",
		},
		"POST:/drivers/ride/start": {
			Method:  "POST",
			Path:    "/drivers/ride/start",
			Handler: "nebeng-jek/internal/rides/handler/http.(*httpHandler).DriverStartRide-fm",
		},
		"POST:/drivers/ride/end": {
			Method:  "POST",
			Path:    "/drivers/ride/end",
			Handler: "nebeng-jek/internal/rides/handler/http.(*httpHandler).DriverEndRide-fm",
		},
		"POST:/drivers/ride/confirm-payment": {
			Method:  "POST",
			Path:    "/drivers/ride/confirm-payment",
			Handler: "nebeng-jek/internal/rides/handler/http.(*httpHandler).DriverConfirmPayment-fm",
		},
		"POST:/riders/ride/create": {
			Method:  "POST",
			Path:    "/riders/ride/create",
			Handler: "nebeng-jek/internal/rides/handler/http.(*httpHandler).RiderCreateNewRide-fm",
		},
		"POST:/riders/ride/confirm": {
			Method:  "POST",
			Path:    "/riders/ride/confirm",
			Handler: "nebeng-jek/internal/rides/handler/http.(*httpHandler).RiderConfirmRide-fm",
		},
	}

	for _, r := range router.Routes() {
		key := fmt.Sprintf("%s:%s", r.Method, r.Path)
		expected, ok := expectedRoutes[key]
		if !ok {
			t.Errorf("route %s is not found", key)
			continue
		}

		assert.Equal(t, expected.Method, r.Method)
		assert.Equal(t, expected.Path, r.Path)
		assert.Equal(t, expected.Handler, r.Handler)
	}
}
