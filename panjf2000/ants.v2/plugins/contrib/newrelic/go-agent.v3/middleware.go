package newrelic

import (
	"context"

	newrelic "github.com/americanas-go/ignite/newrelic/go-agent.v3"
	"github.com/americanas-go/ignite/panjf2000/ants.v2"
	"github.com/americanas-go/log"
	nr "github.com/newrelic/go-agent/v3/newrelic"
)

func NewMiddlewareWithConfigPath(path string) (ants.Middleware, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewMiddlewareWithOptions(o), nil
}

func NewMiddlewareWithOptions(options *Options) ants.Middleware {
	return &middleware{options: options}
}

func NewMiddleware() ants.Middleware {
	log.Trace("creating newrelic middleware for ants")
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}
	return NewMiddlewareWithOptions(o)
}

type middleware struct {
	options *Options
}

func (i *middleware) Before(ctx context.Context) context.Context {

	if !i.options.Enabled || !newrelic.IsEnabled() {
		return ctx
	}

	logger := log.FromContext(ctx).WithTypeOf(*i)

	logger.Trace("creating go routine for newrelic")

	txn := newrelic.FromContext(ctx).NewGoroutine()

	logger.Debug("goroutine for newrelic successfully created in context")

	return nr.NewContext(ctx, txn)
}

func (i *middleware) After(ctx context.Context) {

	if !i.options.Enabled || !newrelic.IsEnabled() {
		return
	}

}
