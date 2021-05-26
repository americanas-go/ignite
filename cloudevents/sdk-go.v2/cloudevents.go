package cloudevents

import (
	"context"

	"github.com/americanas-go/log"
	v2 "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/client"
	"github.com/cloudevents/sdk-go/v2/protocol/http"
)

type Client struct {
	handler Handler
	client  client.Client
}

func NewHTTP(ctx context.Context, handler Handler, opts ...http.Option) *Client {
	logger := log.FromContext(ctx)
	c, err := client.NewHTTP(opts...)
	if err != nil {
		logger.Panic(err.Error())
	}
	return &Client{handler: handler, client: c}
}

// NewDefault has been replaced by NewHTTP
// Deprecated. To get the same as NewDefault provided, please use NewHTTP with
func NewDefaultClient(ctx context.Context, handler Handler, opts ...http.Option) *Client {
	logger := log.FromContext(ctx)
	c, err := v2.NewDefaultClient(opts...)
	if err != nil {
		logger.Panic(err.Error())
	}
	return &Client{handler: handler, client: c}
}

func (s *Client) Start(ctx context.Context) {
	logger := log.FromContext(ctx).WithTypeOf(*s)
	if err := s.client.StartReceiver(ctx, s.handler); err != nil {
		logger.Panic(err.Error())
	}
}
