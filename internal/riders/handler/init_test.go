package handler

import (
	"context"
	"testing"

	"nebeng-jek/internal/pkg/constants"
	mock_nats "nebeng-jek/mock/pkg/messaging/nats"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRegisterHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock subscribe rides
	natsConn := mock_nats.NewMockJetStreamConnection(ctrl)
	natsConn.EXPECT().Subscribe(constants.TopicRideMatchedDriver, gomock.Any(), gomock.Any()).AnyTimes()
	natsConn.EXPECT().Subscribe(constants.TopicRideReadyToPickup, gomock.Any(), gomock.Any()).AnyTimes()
	natsConn.EXPECT().Subscribe(constants.TopicRideStarted, gomock.Any(), gomock.Any()).AnyTimes()
	natsConn.EXPECT().Subscribe(constants.TopicRideEnded, gomock.Any(), gomock.Any()).AnyTimes()
	natsConn.EXPECT().Subscribe(constants.TopicRidePaid, gomock.Any(), gomock.Any()).AnyTimes()

	router := gin.New()
	reg := RegisterHandlerParam{
		Router: &router.RouterGroup,
		NatsJS: natsConn,
		JWTGen: nil, // no tests
	}
	RegisterHandler(context.Background(), reg)

	expectedRoutes := gin.RoutesInfo{
		{
			Method:  "GET",
			Path:    "/ws/riders",
			Handler: "nebeng-jek/internal/riders/handler/http.(*httpHandler).RiderWebsocket-fm",
		},
	}

	for i, r := range router.Routes() {
		assert.Equal(t, expectedRoutes[i].Method, r.Method)
		assert.Equal(t, expectedRoutes[i].Path, r.Path)
		assert.Equal(t, expectedRoutes[i].Handler, r.Handler)
	}
}
