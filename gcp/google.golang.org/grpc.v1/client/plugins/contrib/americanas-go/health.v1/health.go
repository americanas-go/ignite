package health

import (
	"context"

	"github.com/americanas-go/health"
	"github.com/americanas-go/log"
	"google.golang.org/grpc"
)

// Register registers a new health checker plugin for grpc client.
func Register(ctx context.Context, conn *grpc.ClientConn) error {
	o, err := NewOptions()
	if err != nil {
		return nil
	}
	h := NewHealthWithOptions(o)
	return h.Register(ctx, conn)
}

// Health represents health checker plugin for grpc client.
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

// Register registers this health checker plugin for grpc client.
func (i *Health) Register(ctx context.Context, conn *grpc.ClientConn) error {

	logger := log.FromContext(ctx).WithTypeOf(*i)

	logger.Trace("integrating grpc client in health")

	checker := NewChecker(conn)
	hc := health.NewHealthChecker(i.options.Name, i.options.Description, checker, i.options.Required, i.options.Enabled)
	health.Add(hc)

	logger.Debug("grpc client successfully integrated in health")

	return nil
}
