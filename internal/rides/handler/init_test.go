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
	amqpMock.EXPECT().ExchangeDeclare(constants.RideStartedExchange, constants.ExchangeTypeFanout, true, false, false, false, nil).
		Return(nil)
	amqpMock.EXPECT().ExchangeDeclare(constants.RideEndedExchange, constants.ExchangeTypeFanout, true, false, false, false, nil).
		Return(nil)

	router := gin.New()
	RegisterHandler(&router.RouterGroup, nil, nil, amqpMock)

	expectedRoutes := map[string]gin.RouteInfo{
		"PUT:/v1/drivers/availability": {
			Method:  "PUT",
			Path:    "/v1/drivers/availability",
			Handler: "nebeng-jek/internal/rides/handler.(*ridesHandler).SetDriverAvailability-fm",
		},
		"POST:/v1/riders/rides": {
			Method:  "POST",
			Path:    "/v1/riders/rides",
			Handler: "nebeng-jek/internal/rides/handler.(*ridesHandler).CreateNewRide-fm",
		},
		"POST:/v1/riders/rides/confirm": {
			Method:  "POST",
			Path:    "/v1/riders/rides/confirm",
			Handler: "nebeng-jek/internal/rides/handler.(*ridesHandler).ConfirmRideRider-fm",
		},
		"POST:/v1/drivers/rides/confirm": {
			Method:  "POST",
			Path:    "/v1/drivers/rides/confirm",
			Handler: "nebeng-jek/internal/rides/handler.(*ridesHandler).ConfirmRideDriver-fm",
		},
		"POST:/v1/drivers/rides/start": {
			Method:  "POST",
			Path:    "/v1/drivers/rides/start",
			Handler: "nebeng-jek/internal/rides/handler.(*ridesHandler).StartRideDriver-fm",
		},
		"POST:/v1/drivers/rides/end": {
			Method:  "POST",
			Path:    "/v1/drivers/rides/end",
			Handler: "nebeng-jek/internal/rides/handler.(*ridesHandler).EndRideDriver-fm",
		},
		"POST:/v1/drivers/rides/confirm-payment": {
			Method:  "POST",
			Path:    "/v1/drivers/rides/confirm-payment",
			Handler: "nebeng-jek/internal/rides/handler.(*ridesHandler).ConfirmPaymentDriver-fm",
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
