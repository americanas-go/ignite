package cloudevents

import (
	"context"

	"github.com/americanas-go/log"
	"github.com/cloudevents/sdk-go/v2/client"
	"github.com/cloudevents/sdk-go/v2/protocol/http"
)

// Client represents CloudEvents client.
type Client struct {
	handler Handler
	client  client.Client
}

// NewHTTP provides HTTP Protocol client.
func NewHTTP(ctx context.Context, handler Handler, opts ...http.Option) *Client {
	logger := log.FromContext(ctx)
	c, err := client.NewHTTP(opts...)
	if err != nil {
		logger.Panic(err.Error())
	}
	return &Client{handler: handler, client: c}
}

// Start sets up the given handler to handle Receive.
func (s *Client) Start(ctx context.Context) {
	logger := log.FromContext(ctx).WithTypeOf(*s)
	if err := s.client.StartReceiver(ctx, s.handler); err != nil {
		logger.Panic(err.Error())
	}
}
