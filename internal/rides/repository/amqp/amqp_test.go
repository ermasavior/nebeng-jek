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

func TestRepository_BroadcastMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockChan := mock_amqp.NewMockAMQPChannel(ctrl)
	mockChan.EXPECT().ExchangeDeclare(constants.NewRideRequestsExchange, constants.ExchangeTypeFanout, true, false, false, false, nil).Return(nil)
	mockChan.EXPECT().ExchangeDeclare(constants.DriverAcceptedRideExchange, constants.ExchangeTypeFanout, true, false, false, false, nil).Return(nil)
	mockChan.EXPECT().ExchangeDeclare(constants.RideReadyToPickupExchange, constants.ExchangeTypeFanout, true, false, false, false, nil).Return(nil)
	mockChan.EXPECT().ExchangeDeclare(constants.RideStartedExchange, constants.ExchangeTypeFanout, true, false, false, false, nil).Return(nil)
	mockChan.EXPECT().ExchangeDeclare(constants.RideEndedExchange, constants.ExchangeTypeFanout, true, false, false, false, nil).Return(nil)
	mockChan.EXPECT().Close()

	mockConn := mock_amqp.NewMockAMQPConnection(ctrl)
	mockConn.EXPECT().Channel().Return(mockChan, nil)

	r := NewRepository(mockConn)

	ctx := context.TODO()
	topic := constants.NewRideRequestsExchange
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
		mockChan.EXPECT().Publish(
			constants.NewRideRequestsExchange, "", false, false, amqp091.Publishing{
				ContentType: constants.TypeApplicationJSON,
				Body:        msgBytes,
			}).Return(nil)

		err := r.BroadcastMessage(ctx, topic, req)

		assert.NoError(t, err)
	})

	t.Run("failed - publish message returns error", func(t *testing.T) {
		expectedErr := errors.New("error")

		mockChan.EXPECT().Publish(
			constants.NewRideRequestsExchange, "", false, false, amqp091.Publishing{
				ContentType: constants.TypeApplicationJSON,
				Body:        msgBytes,
			}).Return(expectedErr)

		err := r.BroadcastMessage(ctx, topic, req)

		assert.Error(t, err, expectedErr)
	})
}
