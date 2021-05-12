package newrelic

import (
	"context"

	ginats "github.com/jvitoroc/ignite/nats-io/nats.go.v1"
	newrelic "github.com/jvitoroc/ignite/newrelic/go-agent.v3"
	"github.com/nats-io/nats.go"
	"github.com/newrelic/go-agent/v3/integrations/nrnats"
	nr "github.com/newrelic/go-agent/v3/newrelic"
)

type PublisherRegister struct {
}

func (p *PublisherRegister) Before(ctx context.Context, conn *nats.Conn, msg *nats.Msg) (context.Context, error) {
	if !IsEnabled() || !newrelic.IsEnabled() {
		return ctx, nil
	}

	txn := newrelic.FromContext(ctx)
	seg := nrnats.StartPublishSegment(txn, conn, msg.Subject)

	return context.WithValue(ctx, "seg", seg), nil
}

func (p *PublisherRegister) After(ctx context.Context) error {
	seg := ctx.Value("seg").(*nr.MessageProducerSegment)
	seg.End()
	return nil
}

func NewPublisherRegister() ginats.PublisherMiddleware {
	return &PublisherRegister{}
}
