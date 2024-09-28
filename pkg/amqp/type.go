package amqp

import amqp "github.com/rabbitmq/amqp091-go"

type AMQPConnection interface {
	Channel() (*amqp.Channel, error)
}

type AMQPChannel interface {
	ExchangeDeclare(name, kind string, durable, autoDelete, internal, noWait bool, args amqp.Table) error
	Publish(exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error
}
