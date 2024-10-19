package handler

import (
	"fmt"
	"nebeng-jek/internal/pkg/constants"
	mock_nats "nebeng-jek/mock/pkg/messaging/nats"
	mock_redis "nebeng-jek/mock/pkg/redis"
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
		Router: &router.RouterGroup,
		Redis:  mock_redis.NewMockCollections(ctrl),
		DB:     nil, // no tests
		NatsJS: natsConn,
		JWTGen: nil, // no tests
	}
	RegisterHandler(reg)

	expectedRoutes := map[string]gin.RouteInfo{
		"PUT:/drivers/availability": {
			Method:  "PUT",
			Path:    "/drivers/availability",
			Handler: "nebeng-jek/internal/rides/handler/http.(*httpHandler).SetDriverAvailability-fm",
		},
		"POST:/riders/rides": {
			Method:  "POST",
			Path:    "/riders/rides",
			Handler: "nebeng-jek/internal/rides/handler/http.(*httpHandler).CreateNewRide-fm",
		},
		"POST:/riders/rides/confirm": {
			Method:  "POST",
			Path:    "/riders/rides/confirm",
			Handler: "nebeng-jek/internal/rides/handler/http.(*httpHandler).ConfirmRideRider-fm",
		},
		"POST:/drivers/rides/confirm": {
			Method:  "POST",
			Path:    "/drivers/rides/confirm",
			Handler: "nebeng-jek/internal/rides/handler/http.(*httpHandler).ConfirmRideDriver-fm",
		},
		"POST:/drivers/rides/start": {
			Method:  "POST",
			Path:    "/drivers/rides/start",
			Handler: "nebeng-jek/internal/rides/handler/http.(*httpHandler).StartRideDriver-fm",
		},
		"POST:/drivers/rides/end": {
			Method:  "POST",
			Path:    "/drivers/rides/end",
			Handler: "nebeng-jek/internal/rides/handler/http.(*httpHandler).EndRideDriver-fm",
		},
		"POST:/drivers/rides/confirm-payment": {
			Method:  "POST",
			Path:    "/drivers/rides/confirm-payment",
			Handler: "nebeng-jek/internal/rides/handler/http.(*httpHandler).ConfirmPaymentDriver-fm",
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
