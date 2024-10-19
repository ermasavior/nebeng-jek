package handler

import (
	"testing"

	mock_nats "nebeng-jek/mock/pkg/messaging/nats"

	"nebeng-jek/internal/pkg/constants"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRegisterHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock subscribe rides
	natsConn := mock_nats.NewMockJetStreamConnection(ctrl)
	natsConn.EXPECT().Subscribe(constants.TopicRideNewRequest, gomock.Any(), gomock.Any()).AnyTimes()
	natsConn.EXPECT().Subscribe(constants.TopicRideReadyToPickup, gomock.Any(), gomock.Any()).AnyTimes()

	router := gin.New()
	reg := RegisterHandlerParam{
		Router: &router.RouterGroup,
		NatsJS: natsConn,
		JWTGen: nil, // no tests
	}
	RegisterHandler(reg)

	expectedRoutes := gin.RoutesInfo{
		{
			Method:  "GET",
			Path:    "/ws/drivers",
			Handler: "nebeng-jek/internal/drivers/handler/http.(*httpHandler).DriverAllocationWebsocket-fm",
		},
	}

	for i, r := range router.Routes() {
		assert.Equal(t, expectedRoutes[i].Method, r.Method)
		assert.Equal(t, expectedRoutes[i].Path, r.Path)
		assert.Equal(t, expectedRoutes[i].Handler, r.Handler)
	}
}
