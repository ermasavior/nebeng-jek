package handler

import (
	"fmt"
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
	amqpMock.EXPECT().ExchangeDeclare(constants.NewRideRequestsExchange, constants.ExchangeTypeFanout, true, false, false, false, nil).
		Return(nil)
	amqpMock.EXPECT().ExchangeDeclare(constants.DriverAcceptedRideExchange, constants.ExchangeTypeFanout, true, false, false, false, nil).
		Return(nil)
	amqpMock.EXPECT().ExchangeDeclare(constants.RideReadyToPickupExchange, constants.ExchangeTypeFanout, true, false, false, false, nil).
		Return(nil)

	router := gin.New()
	RegisterHandler(&router.RouterGroup, nil, nil, amqpMock)

	expectedRoutes := map[string]gin.RouteInfo{
		"PUT:/drivers/availability": {
			Method:  "PUT",
			Path:    "/drivers/availability",
			Handler: "nebeng-jek/internal/rides/handler.(*ridesHandler).SetDriverAvailability-fm",
		},
		"POST:/riders/rides": {
			Method:  "POST",
			Path:    "/riders/rides",
			Handler: "nebeng-jek/internal/rides/handler.(*ridesHandler).CreateNewRide-fm",
		},
		"POST:/riders/rides/confirm": {
			Method:  "POST",
			Path:    "/riders/rides/confirm",
			Handler: "nebeng-jek/internal/rides/handler.(*ridesHandler).ConfirmRideRider-fm",
		},
		"POST:/drivers/rides/confirm": {
			Method:  "POST",
			Path:    "/drivers/rides/confirm",
			Handler: "nebeng-jek/internal/rides/handler.(*ridesHandler).ConfirmRideDriver-fm",
		},
		"POST:/drivers/rides/start": {
			Method:  "POST",
			Path:    "/drivers/rides/start",
			Handler: "nebeng-jek/internal/rides/handler.(*ridesHandler).StartRideDriver-fm",
		},
	}

	for _, r := range router.Routes() {
		expected, ok := expectedRoutes[fmt.Sprintf("%s:%s", r.Method, r.Path)]
		if !ok {
			t.Errorf("route is not found")
			continue
		}

		assert.Equal(t, expected.Method, r.Method)
		assert.Equal(t, expected.Path, r.Path)
		assert.Equal(t, expected.Handler, r.Handler)
	}
}
