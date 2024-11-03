package handler

import (
	"context"
	"fmt"
	"nebeng-jek/internal/pkg/constants"
	mock_nats "nebeng-jek/mock/pkg/messaging/nats"
	mock_redis "nebeng-jek/mock/pkg/redis"
	"nebeng-jek/pkg/configs"
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
		NatsJS: natsConn,
		Cfg:    configs.NewMockConfig(),
	}
	RegisterHandler(context.Background(), reg)

	expectedRoutes := map[string]gin.RouteInfo{
		"POST:/drivers/available": {
			Method:  "POST",
			Path:    "/drivers/available",
			Handler: "nebeng-jek/internal/location/handler/http.(*httpHandler).AddAvailableDriver-fm",
		},
		"DELETE:/drivers/available/:driver_id": {
			Method:  "DELETE",
			Path:    "/drivers/available/:driver_id",
			Handler: "nebeng-jek/internal/location/handler/http.(*httpHandler).RemoveAvailableDriver-fm",
		},
		"GET:/drivers/available/nearby": {
			Method:  "GET",
			Path:    "/drivers/available/nearby",
			Handler: "nebeng-jek/internal/location/handler/http.(*httpHandler).GetNearestAvailableDrivers-fm",
		},
		"GET:/drivers/ride-path": {
			Method:  "GET",
			Path:    "/drivers/ride-path",
			Handler: "nebeng-jek/internal/location/handler/http.(*httpHandler).GetRidePath-fm",
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
