package handler

import (
	"testing"

	"nebeng-jek/internal/pkg/constants"
	mock_amqp "nebeng-jek/mock/pkg/amqp"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
)

func TestRegisterHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock subscribe new rides
	amqpMock := mock_amqp.NewMockAMQPChannel(ctrl)
	amqpMock.EXPECT().ExchangeDeclare(constants.RideRequestsExchange, constants.ExchangeTypeFanout, true, false, false, false, nil).
		Return(nil).AnyTimes()
	amqpMock.EXPECT().QueueDeclare(gomock.Any(), false, false, true, false, nil).
		Return(amqp091.Queue{}, nil).AnyTimes()
	amqpMock.EXPECT().QueueBind(gomock.Any(), gomock.Any(), constants.RideRequestsExchange, gomock.Any(), nil).
		Return(nil).AnyTimes()
	amqpMock.EXPECT().Consume(gomock.Any(), gomock.Any(), true, false, false, false, nil).
		Return(nil, nil).AnyTimes()

	router := gin.New()
	RegisterHandler(&router.RouterGroup, amqpMock)

	expectedRoutes := gin.RoutesInfo{
		{
			Method:  "GET",
			Path:    "/ws/drivers",
			Handler: "nebeng-jek/internal/drivers/handler.(*driversHandler).DriverAllocationWebsocket-fm",
		},
	}

	for i, r := range router.Routes() {
		assert.Equal(t, expectedRoutes[i].Method, r.Method)
		assert.Equal(t, expectedRoutes[i].Path, r.Path)
		assert.Equal(t, expectedRoutes[i].Handler, r.Handler)
	}
}
