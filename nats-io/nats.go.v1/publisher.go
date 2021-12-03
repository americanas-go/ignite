package nats

import (
	"context"

	"github.com/nats-io/nats.go"
)

// PublisherMiddleware defines an interface for a nats publisher plugin.
type PublisherMiddleware interface {
	Before(context.Context, *nats.Conn, *nats.Msg) (context.Context, error)
	After(context.Context) error
}

// Publisher represents a nats publisher.
type Publisher struct {
	conn        *nats.Conn
	options     *Options
	middlewares []PublisherMiddleware
}

// NewPublisherWithConfigPath returns a new nats publisher with options from config path.
func NewPublisherWithConfigPath(ctx context.Context, path string, middlewares ...PublisherMiddleware) (*Publisher, error) {
	options, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewPublisherWithOptions(ctx, options, middlewares...)
}

// NewPublisherWithOptions returns a new nats publisher with options.
func NewPublisherWithOptions(ctx context.Context, options *Options, middlewares ...PublisherMiddleware) (*Publisher, error) {
	conn, err := NewConnWithOptions(ctx, options)
	if err != nil {
		return nil, err
	}
	return &Publisher{conn, options, middlewares}, nil
}

// NewPublisher returns a new nats publisher with default options.
func NewPublisher(ctx context.Context, middlewares ...PublisherMiddleware) (*Publisher, error) {
	options, err := NewOptions()
	if err != nil {
		return nil, err
	}
	return NewPublisherWithOptions(ctx, options, middlewares...)
}

// Publish publishes the message to nats.
func (p *Publisher) Publish(ctx context.Context, msg *nats.Msg) error {

	for _, middleware := range p.middlewares {
		ctxx, err := middleware.Before(ctx, p.conn, msg)
		if err != nil {
			return err
		}
		defer func() {
			middleware.After(ctxx)
		}()
	}

	return p.conn.PublishMsg(msg)
}

// Conn returns this publisher connection.
func (p *Publisher) Conn() *nats.Conn {
	return p.conn
}
