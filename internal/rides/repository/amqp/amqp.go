package amqp

import (
	"context"
	"encoding/json"
	constants "nebeng-jek/internal/pkg/constants/pubsub"
	"nebeng-jek/internal/rides/model"
	"nebeng-jek/internal/rides/repository"
	"nebeng-jek/pkg/amqp"

	"github.com/rabbitmq/amqp091-go"
)

type ridesRepo struct {
	chann amqp.AMQPChannel
}

func NewRidesRepository(chann amqp.AMQPChannel) repository.RidesPubsubRepository {
	return &ridesRepo{
		chann: chann,
	}
}

func (r *ridesRepo) BroadcastRideToDrivers(ctx context.Context, msg model.RideRequestMessage) error {
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = r.chann.Publish(
		constants.RideRequestsFanout, // exchange name
		"",                           // routing key (ignored for fanout)
		false,                        // mandatory
		false,                        // immediate
		amqp091.Publishing{
			ContentType: constants.TypeApplicationJSON, // Set content type to JSON
			Body:        msgBytes,                      // JSON message body
		},
	)

	if err != nil {
		return err
	}
	return nil
}
