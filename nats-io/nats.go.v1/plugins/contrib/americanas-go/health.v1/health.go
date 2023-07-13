package health

import (
	"context"

	"github.com/americanas-go/health"
	"github.com/americanas-go/log"
	"github.com/nats-io/nats.go"
)

// Register registers a new health checker plugin for nats connection.
func Register(ctx context.Context) (func(context.Context, *nats.Conn) error, func(context.Context, []nats.Option) (error, []nats.Option)) {
	return func(ctx context.Context, conn *nats.Conn) error {
		o, err := NewOptions()
		if err != nil {
			return err
		}
		h := NewHealthWithOptions(o)
		return h.Register(ctx, conn)
	}, nil
}

// Health represents health checker plugin for nats connection.
type Health struct {
	options *Options
}

// NewHealthWithOptions returns a new health checker plugin with options.
func NewHealthWithOptions(options *Options) *Health {
	return &Health{options: options}
}

// NewHealthWithConfigPath returns a new health checker plugin with options from config path.
func NewHealthWithConfigPath(path string) (*Health, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewHealthWithOptions(o), nil
}

// NewHealth returns a new health checker plugin with default options.
func NewHealth() *Health {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewHealthWithOptions(o)
}

// Register registers this health checker plugin for nats connection.
func (i *Health) Register(ctx context.Context, conn *nats.Conn) error {

	logger := log.FromContext(ctx).WithTypeOf(*i)

	logger.Trace("integrating nats in health")

	checker := NewChecker(conn)
	hc := health.NewHealthChecker(i.options.Name, i.options.Description, checker, i.options.Required, i.options.Enabled)
	health.Add(hc)

	logger.Debug("nats successfully integrated in health")

	return nil
}
