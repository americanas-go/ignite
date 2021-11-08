package newrelic

import (
	"context"

	ginats "github.com/americanas-go/ignite/nats-io/nats.go.v1"
	newrelic "github.com/americanas-go/ignite/newrelic/go-agent.v3"
	"github.com/americanas-go/log"
	"github.com/nats-io/nats.go"
	"github.com/newrelic/go-agent/v3/integrations/nrnats"
	nr "github.com/newrelic/go-agent/v3/newrelic"
)

func NewPublisherRegister() ginats.PublisherMiddleware {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewPublisherRegisterWithOptions(o)
}

func NewPublisherRegisterWithConfigPath(path string) (ginats.PublisherMiddleware, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewPublisherRegisterWithOptions(o), nil
}

func NewPublisherRegisterWithOptions(options *Options) ginats.PublisherMiddleware {
	return &Publisher{options: options}
}

type Publisher struct {
	options *Options
}

func (p *Publisher) Before(ctx context.Context, conn *nats.Conn, msg *nats.Msg) (context.Context, error) {
	if !p.options.Enabled || !newrelic.IsEnabled() {
		return ctx, nil
	}

	txn := newrelic.FromContext(ctx)
	seg := nrnats.StartPublishSegment(txn, conn, msg.Subject)

	return context.WithValue(ctx, "seg", seg), nil
}

func (p *Publisher) After(ctx context.Context) error {
	seg := ctx.Value("seg").(*nr.MessageProducerSegment)
	seg.End()
	return nil
}
