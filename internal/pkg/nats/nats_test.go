package nats

import (
	"context"
	"encoding/json"
	"errors"
	"nebeng-jek/internal/pkg/constants"
	mock_nats "nebeng-jek/mock/pkg/messaging/nats"
	"nebeng-jek/pkg/utils"
	"testing"

	"github.com/golang/mock/gomock"
	nats_go "github.com/nats-io/nats.go"
	"github.com/stretchr/testify/assert"
)

func TestPubsubRepo_BroadcastMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	natsJS := mock_nats.NewMockJetStreamConnection(ctrl)
	r := NewPubsubRepository(natsJS)

	topic := constants.TopicRideStarted
	msg := map[string]string{
		"key": "value",
	}
	msgBytes, _ := json.Marshal(msg)

	t.Run("success - broadcast message to nats", func(t *testing.T) {
		natsJS.EXPECT().PublishAsync(topic, msgBytes).Return(nil, nil)

		err := r.BroadcastMessage(context.TODO(), topic, msg)
		assert.Nil(t, err)
	})
	t.Run("failed - invalid msg", func(t *testing.T) {
		err := r.BroadcastMessage(context.TODO(), topic, utils.FailingType{})
		assert.Error(t, err)
	})
	t.Run("failed - broadcast message failed", func(t *testing.T) {
		expectedErr := errors.New("error from nats")
		natsJS.EXPECT().PublishAsync(topic, msgBytes).Return(nil, expectedErr)

		err := r.BroadcastMessage(context.TODO(), topic, msg)
		assert.Error(t, err, expectedErr.Error())
	})
}

func TestSubscribeMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	natsJS := mock_nats.NewMockJetStreamConnection(ctrl)

	topic := constants.TopicRideStarted
	msgHandler := func(msg *nats_go.Msg) {}

	t.Run("failed - subscribe message failed", func(t *testing.T) {
		expectedErr := errors.New("error from nats")
		natsJS.EXPECT().Subscribe(topic, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, expectedErr)

		SubscribeMessage(natsJS, topic, msgHandler)
	})
}
