package handler

import (
	"nebeng-jek/internal/pkg/constants"
	mock_amqp "nebeng-jek/mock/pkg/amqp"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRegisterHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	amqpMock := mock_amqp.NewMockAMQPChannel(ctrl)
	amqpMock.EXPECT().ExchangeDeclare(constants.RideRequestsExchange, "fanout", true, false, false, false, nil).
		Return(nil)

	router := gin.New()
	RegisterHandler(&router.RouterGroup, nil, nil, amqpMock)

	expectedRoutes := gin.RoutesInfo{
		{
			Method:  "PUT",
			Path:    "/drivers/availability",
			Handler: "nebeng-jek/internal/rides/handler.(*ridesHandler).SetDriverAvailability-fm",
		},
		{
			Method:  "POST",
			Path:    "/riders/rides",
			Handler: "nebeng-jek/internal/rides/handler.(*ridesHandler).CreateNewRide-fm",
		},
	}

	for i, r := range router.Routes() {
		assert.Equal(t, expectedRoutes[i].Method, r.Method)
		assert.Equal(t, expectedRoutes[i].Path, r.Path)
		assert.Equal(t, expectedRoutes[i].Handler, r.Handler)
	}
}