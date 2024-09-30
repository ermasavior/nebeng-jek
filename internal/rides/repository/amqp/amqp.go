package amqp

import (
	"context"
	"encoding/json"
	"nebeng-jek/internal/pkg/constants"
	"nebeng-jek/internal/rides/model"
	"nebeng-jek/internal/rides/repository"
	"nebeng-jek/pkg/amqp"
	"nebeng-jek/pkg/logger"

	"github.com/rabbitmq/amqp091-go"
)

type ridesRepo struct {
	chann amqp.AMQPChannel
}

func NewRepository(chann amqp.AMQPChannel) repository.RidesPubsubRepository {
	err := chann.ExchangeDeclare(
		constants.RideRequestsExchange,
		"fanout", // exchange type: fanout
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		logger.Fatal(context.Background(), "failed to declare an amqp exchange", map[string]interface{}{
			"error": err,
		})
	}

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
		constants.RideRequestsExchange, // exchange name
		"",                             // routing key (ignored for fanout)
		false,                          // mandatory
		false,                          // immediate
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
