package amqp

import (
	"context"
	"encoding/json"
	"nebeng-jek/internal/pkg/constants"
	"nebeng-jek/internal/rides/repository"
	"nebeng-jek/pkg/amqp"
	"nebeng-jek/pkg/logger"

	"github.com/rabbitmq/amqp091-go"
)

type ridesRepo struct {
	chann amqp.AMQPChannel
}

func NewRepository(amqpConn amqp.AMQPConnection) repository.RidesPubsubRepository {
	chann, err := amqpConn.Channel()
	if err != nil {
		logger.Fatal(context.Background(), "error initializing amqp channel", map[string]interface{}{logger.ErrorKey: err})
	}
	defer chann.Close()

	err = chann.ExchangeDeclare(
		constants.NewRideRequestsExchange,
		constants.ExchangeTypeFanout, // exchange type: fanout
		true,                         // durable
		false,                        // auto-deleted
		false,                        // internal
		false,                        // no-wait
		nil,                          // arguments
	)
	if err != nil {
		logger.Fatal(context.Background(), "failed to declare ride request exchange", map[string]interface{}{
			"error": err,
		})
	}

	err = chann.ExchangeDeclare(constants.DriverAcceptedRideExchange, constants.ExchangeTypeFanout, true, false, false, false, nil)
	if err != nil {
		logger.Fatal(context.Background(), "failed to declare matched ride exchange", map[string]interface{}{
			"error": err,
		})
	}

	err = chann.ExchangeDeclare(constants.RideReadyToPickupExchange, constants.ExchangeTypeFanout, true, false, false, false, nil)
	if err != nil {
		logger.Fatal(context.Background(), "failed to declare matched ride exchange", map[string]interface{}{
			"error": err,
		})
	}

	err = chann.ExchangeDeclare(constants.RideStartedExchange, constants.ExchangeTypeFanout, true, false, false, false, nil)
	if err != nil {
		logger.Fatal(context.Background(), "failed to declare matched ride exchange", map[string]interface{}{
			"error": err,
		})
	}

	err = chann.ExchangeDeclare(constants.RideEndedExchange, constants.ExchangeTypeFanout, true, false, false, false, nil)
	if err != nil {
		logger.Fatal(context.Background(), "failed to declare matched ride exchange", map[string]interface{}{
			"error": err,
		})
	}

	return &ridesRepo{
		chann: chann,
	}
}

func (r *ridesRepo) BroadcastMessage(ctx context.Context, topic string, msg interface{}) error {
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = r.chann.Publish(
		topic, // exchange name
		"",    // routing key (ignored for fanout)
		false, // mandatory
		false, // immediate
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
