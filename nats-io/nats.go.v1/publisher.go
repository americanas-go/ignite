package nats

import (
	"context"

	"github.com/nats-io/nats.go"
)

// Publisher represents a nats publisher.
type Publisher struct {
	conn    *nats.Conn
	options *Options
}

// NewPublisherWithConfigPath returns a new nats publisher with options from config path.
func NewPublisherWithConfigPath(ctx context.Context, path string) (*Publisher, error) {
	options, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewPublisherWithOptions(ctx, options)
}

// NewPublisherWithOptions returns a new nats publisher with options.
func NewPublisherWithOptions(ctx context.Context, options *Options) (*Publisher, error) {
	conn, err := NewConnWithOptions(ctx, options)
	if err != nil {
		return nil, err
	}
	return &Publisher{conn, options}, nil
}

// NewPublisher returns a new nats publisher with default options.
func NewPublisher(ctx context.Context) (*Publisher, error) {
	options, err := NewOptions()
	if err != nil {
		return nil, err
	}
	return NewPublisherWithOptions(ctx, options)
}

// Publish publishes the message to nats.
func (p *Publisher) Publish(ctx context.Context, msg *nats.Msg) error {
	return p.conn.PublishMsg(msg)
}

// Conn returns this publisher connection.
func (p *Publisher) Conn() *nats.Conn {
	return p.conn
}
