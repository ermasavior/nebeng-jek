package amqp

import (
	"context"
	"encoding/json"
	"errors"
	"nebeng-jek/internal/pkg/constants"
	"nebeng-jek/internal/rides/model"
	mock_amqp "nebeng-jek/mock/pkg/amqp"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
)

func TestRepository_BroadcastRideToDrivers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAMQP := mock_amqp.NewMockAMQPChannel(ctrl)
	mockAMQP.EXPECT().ExchangeDeclare(constants.RideRequestsExchange, constants.ExchangeTypeFanout, true, false, false, false, nil).Return(nil)
	mockAMQP.EXPECT().ExchangeDeclare(constants.MatchedRideExchange, constants.ExchangeTypeFanout, true, false, false, false, nil).Return(nil)

	r := NewRepository(mockAMQP)

	ctx := context.TODO()
	req := model.RideRequestMessage{
		RideID: 111,
		Rider: model.RiderData{
			ID:     777,
			Name:   "Agus",
			MSISDN: "0812222",
		},
	}

	msgBytes, _ := json.Marshal(req)

	t.Run("success - publish message to ride requests pubsub channel", func(t *testing.T) {
		mockAMQP.EXPECT().Publish(
			constants.RideRequestsExchange, "", false, false, amqp091.Publishing{
				ContentType: constants.TypeApplicationJSON,
				Body:        msgBytes,
			}).Return(nil)

		err := r.BroadcastRideToDrivers(ctx, req)

		assert.NoError(t, err)
	})

	t.Run("failed - publish message returns error", func(t *testing.T) {
		expectedErr := errors.New("error")

		mockAMQP.EXPECT().Publish(
			constants.RideRequestsExchange, "", false, false, amqp091.Publishing{
				ContentType: constants.TypeApplicationJSON,
				Body:        msgBytes,
			}).Return(expectedErr)

		err := r.BroadcastRideToDrivers(ctx, req)

		assert.Error(t, err, expectedErr)
	})
}

func TestRepository_BroadcastMatchedRideToRider(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAMQP := mock_amqp.NewMockAMQPChannel(ctrl)
	mockAMQP.EXPECT().ExchangeDeclare(constants.RideRequestsExchange, constants.ExchangeTypeFanout, true, false, false, false, nil).Return(nil)
	mockAMQP.EXPECT().ExchangeDeclare(constants.MatchedRideExchange, constants.ExchangeTypeFanout, true, false, false, false, nil).Return(nil)

	r := NewRepository(mockAMQP)

	ctx := context.TODO()
	req := model.MatchedRideMessage{
		RideID: 111,
		Driver: model.DriverData{
			ID:           777,
			Name:         "Agus",
			MSISDN:       "0812222",
			VehicleType:  model.VehicleTypeCar,
			VehiclePlate: "B1212",
		},
	}

	msgBytes, _ := json.Marshal(req)

	t.Run("success - publish message to matched ride pubsub channel", func(t *testing.T) {
		mockAMQP.EXPECT().Publish(
			constants.MatchedRideExchange, "", false, false, amqp091.Publishing{
				ContentType: constants.TypeApplicationJSON,
				Body:        msgBytes,
			}).Return(nil)

		err := r.BroadcastMatchedRideToRider(ctx, req)

		assert.NoError(t, err)
	})

	t.Run("failed - publish message returns error", func(t *testing.T) {
		expectedErr := errors.New("error")

		mockAMQP.EXPECT().Publish(
			constants.MatchedRideExchange, "", false, false, amqp091.Publishing{
				ContentType: constants.TypeApplicationJSON,
				Body:        msgBytes,
			}).Return(expectedErr)

		err := r.BroadcastMatchedRideToRider(ctx, req)

		assert.Error(t, err, expectedErr)
	})
}
