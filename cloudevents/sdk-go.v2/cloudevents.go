package cloudevents

import (
	"context"

	"github.com/americanas-go/log"
	v2 "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/client"
)

type Client struct {
	handler Handler
	client  client.Client
}

func NewDefaultClient(ctx context.Context, handler Handler) *Client {
	logger := log.FromContext(ctx)
	c, err := v2.NewDefaultClient()
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
