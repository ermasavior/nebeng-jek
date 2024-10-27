package nats

import "github.com/nats-io/nats.go"

type NATSConnection interface {
	JetStream(opts ...nats.JSOpt) (nats.JetStreamContext, error)
	Close()
}

type JetStreamConnection interface {
	Publish(subj string, data []byte, opts ...nats.PubOpt) (*nats.PubAck, error)
	PublishAsync(subj string, data []byte, opts ...nats.PubOpt) (nats.PubAckFuture, error)
	Subscribe(subj string, cb nats.MsgHandler, opts ...nats.SubOpt) (*nats.Subscription, error)
	AddConsumer(stream string, cfg *nats.ConsumerConfig, opts ...nats.JSOpt) (*nats.ConsumerInfo, error)
}
