package amqp

import (
	"context"
	"nebeng-jek/internal/pkg/constants"
	"nebeng-jek/pkg/logger"

	"github.com/rabbitmq/amqp091-go"
	amqp "github.com/rabbitmq/amqp091-go"
)

func InitAMQPConnection(url string) (AMQPConnection, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func ConsumeMessageToExchange(ctx context.Context, exchange string, ridesChannel AMQPChannel) (<-chan amqp091.Delivery, error) {
	err := ridesChannel.ExchangeDeclare(
		exchange,
		constants.ExchangeTypeFanout, // exchange type: fanout
		true,                         // durable
		false,                        // auto-deleted
		false,                        // internal
		false,                        // no-wait
		nil,                          // arguments
	)
	if err != nil {
		logger.Error(ctx, "failed to declare an amqp exchange", map[string]interface{}{
			"error":    err,
			"exchange": exchange,
		})
		return nil, err
	}

	q, err := ridesChannel.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		logger.Error(ctx, "failed to declare amqp queue", map[string]interface{}{
			"error":    err,
			"exchange": exchange,
		})
		return nil, err
	}

	err = ridesChannel.QueueBind(
		q.Name,   // queue name
		"",       // routing key
		exchange, // exchange
		false,
		nil,
	)
	if err != nil {
		logger.Error(ctx, "failed to bind amqp queue", map[string]interface{}{
			"error":    err,
			"exchange": exchange,
		})
		return nil, err
	}

	msgs, err := ridesChannel.Consume(
		q.Name,
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		logger.Error(ctx, "failed to consume amqp queue", map[string]interface{}{
			"error":    err,
			"exchange": exchange,
		})
		return nil, err
	}

	return msgs, nil
}
