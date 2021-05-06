package nats

import (
	"context"

	"github.com/nats-io/nats.go"
)

type msgHandler func(nats.MsgHandler) nats.MsgHandler

type Subscriber struct {
	conn        *nats.Conn
	options     *Options
	msgHandlers []msgHandler
}

func NewSubscriberWithOptions(ctx context.Context, options *Options, msgHandlers ...msgHandler) (*Subscriber, error) {
	conn, err := NewConnWithOptions(ctx, options)
	if err != nil {
		return nil, err
	}
	return &Subscriber{conn, options, msgHandlers}, nil
}

func NewSubscriber(ctx context.Context, msgHandlers ...msgHandler) (*Subscriber, error) {
	options, err := NewOptions()
	if err != nil {
		return nil, err
	}
	return NewSubscriberWithOptions(ctx, options, msgHandlers...)
}

func (p *Subscriber) Subscribe(subj string, queue string, cb nats.MsgHandler) (*nats.Subscription, error) {
	for _, msgHandler := range p.msgHandlers {
		cb = msgHandler(cb)
	}
	return p.conn.QueueSubscribe(subj, queue, cb)
}

func (p *Subscriber) Conn() *nats.Conn {
	return p.conn
}
