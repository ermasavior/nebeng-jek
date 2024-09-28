package amqp

import amqp "github.com/rabbitmq/amqp091-go"

func InitAMQPConnection(url string) (AMQPConnection, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
