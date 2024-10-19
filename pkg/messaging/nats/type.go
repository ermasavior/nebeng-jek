package nats

import "github.com/nats-io/nats.go"

type NATSConnection interface {
	JetStream(opts ...nats.JSOpt) (nats.JetStreamContext, error)
	Close()
}

type JetStreamConnection interface {
	Publish(subj string, data []byte, opts ...nats.PubOpt) (*nats.PubAck, error)
	Subscribe(subj string, cb nats.MsgHandler, opts ...nats.SubOpt) (*nats.Subscription, error)
}
