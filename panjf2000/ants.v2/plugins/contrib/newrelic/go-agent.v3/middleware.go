package newrelic

import (
	"context"

	newrelic "github.com/americanas-go/ignite/newrelic/go-agent.v3"
	"github.com/americanas-go/ignite/panjf2000/ants.v2"
	"github.com/americanas-go/log"
	nr "github.com/newrelic/go-agent/v3/newrelic"
)

type middleware struct {
}

func (i *middleware) Before(ctx context.Context) context.Context {

	if IsEnabled() || !newrelic.IsEnabled() {
		return ctx
	}

	logger := log.FromContext(ctx).WithTypeOf(*i)

	logger.Trace("creating go routine for newrelic")

	txn := newrelic.FromContext(ctx).NewGoroutine()

	logger.Debug("goroutine for newrelic successfully created in context")

	return nr.NewContext(ctx, txn)
}

func (i *middleware) After(ctx context.Context) {

	if IsEnabled() || !newrelic.IsEnabled() {
		return
	}

}

func NewMiddleware() ants.Middleware {
	log.Trace("creating newrelic middleware for ants")
	return &middleware{}
}
