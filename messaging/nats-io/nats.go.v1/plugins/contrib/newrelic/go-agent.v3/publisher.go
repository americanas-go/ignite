package newrelic

import (
	"context"

	newrelic "github.com/americanas-go/ignite/apm/newrelic/go-agent.v3"
	ginats "github.com/americanas-go/ignite/messaging/nats-io/nats.go.v1"
	"github.com/americanas-go/log"
	"github.com/nats-io/nats.go"
	"github.com/newrelic/go-agent/v3/integrations/nrnats"
	nr "github.com/newrelic/go-agent/v3/newrelic"
)

// NewPublisherRegister returns a new newrelic plugin with default options.
func NewPublisherRegister() ginats.PublisherMiddleware {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewPublisherRegisterWithOptions(o)
}

// NewPublisherRegisterWithConfigPath returns a new newrelic plugin with options from config path.
func NewPublisherRegisterWithConfigPath(path string) (ginats.PublisherMiddleware, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewPublisherRegisterWithOptions(o), nil
}

// NewPublisherRegisterWithOptions returns a new newrelic plugin with options.
func NewPublisherRegisterWithOptions(options *Options) ginats.PublisherMiddleware {
	return &Publisher{options: options}
}

// Publisher represents newrelic plugin for nats publisher.
type Publisher struct {
	options *Options
}

// Before creates and starts a newrelic.MessageProducerSegment
// (https://godoc.org/github.com/newrelic/go-agent#MessageProducerSegment)
// for NATS publishers.
func (p *Publisher) Before(ctx context.Context, conn *nats.Conn, msg *nats.Msg) (context.Context, error) {
	if !p.options.Enabled || !newrelic.IsEnabled() {
		return ctx, nil
	}

	txn := newrelic.FromContext(ctx)
	seg := nrnats.StartPublishSegment(txn, conn, msg.Subject)

	return context.WithValue(ctx, "seg", seg), nil
}

// After finishes the newrelic.MessageProducerSegment
// (https://godoc.org/github.com/newrelic/go-agent#MessageProducerSegment)
// on context.
func (p *Publisher) After(ctx context.Context) error {
	seg := ctx.Value("seg").(*nr.MessageProducerSegment)
	seg.End()
	return nil
}
