package nats

import (
	"context"

	"github.com/nats-io/nats.go"
)

type msgHandler func(nats.MsgHandler) nats.MsgHandler

// Subscriber represents a nats subscriber.
type Subscriber struct {
	conn        *nats.Conn
	options     *Options
	msgHandlers []msgHandler
}

// NewSubscriberWithConfigPath returns a new nats subscriber with options from config path.
func NewSubscriberWithConfigPath(ctx context.Context, path string, msgHandlers ...msgHandler) (*Subscriber, error) {
	options, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewSubscriberWithOptions(ctx, options, msgHandlers...)
}

// NewSubscriberWithOptions returns a new nats subscriber with options.
func NewSubscriberWithOptions(ctx context.Context, options *Options, msgHandlers ...msgHandler) (*Subscriber, error) {
	conn, err := NewConnWithOptions(ctx, options)
	if err != nil {
		return nil, err
	}
	return &Subscriber{conn, options, msgHandlers}, nil
}

// NewSubscriber returns a new nats subscriber with default options.
func NewSubscriber(ctx context.Context, msgHandlers ...msgHandler) (*Subscriber, error) {
	options, err := NewOptions()
	if err != nil {
		return nil, err
	}
	return NewSubscriberWithOptions(ctx, options, msgHandlers...)
}

// Subscribe subscribes a handler to nats queue subject.
func (p *Subscriber) Subscribe(subj string, queue string, cb nats.MsgHandler) (*nats.Subscription, error) {
	for _, msgHandler := range p.msgHandlers {
		cb = msgHandler(cb)
	}
	return p.conn.QueueSubscribe(subj, queue, cb)
}

func (p *Subscriber) Conn() *nats.Conn {
	return p.conn
}
