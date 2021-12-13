package health

import (
	"context"

	"github.com/americanas-go/health"
	"github.com/americanas-go/log"
	"github.com/go-redis/redis/v8"
)

// ClientHealth represents client health.
type ClientHealth struct {
	options *Options
}

// NewClientHealthWithOptions returns a health with the options provided.
func NewClientHealthWithOptions(options *Options) *ClientHealth {
	return &ClientHealth{options: options}
}

// NewClientHealthWithConfigPath returns a client health with options from config path.
func NewClientHealthWithConfigPath(path string) (*ClientHealth, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewClientHealthWithOptions(o), nil
}

// NewClientHealth returns a client health with default options.
func NewClientHealth() (*ClientHealth, error) {
	o, err := NewOptions()
	if err != nil {
		return nil, err
	}

	return NewClientHealthWithOptions(o), nil
}

// Register registers a new ClientChecker in the health package.
func (i *ClientHealth) Register(ctx context.Context, client *redis.Client) error {

	logger := log.FromContext(ctx).WithTypeOf(*i)

	logger.Trace("integrating redis in health")

	checker := NewClientChecker(client)
	hc := health.NewHealthChecker(i.options.Name, i.options.Description, checker, i.options.Required, i.options.Enabled)
	health.Add(hc)

	logger.Debug("redis successfully integrated in health")

	return nil
}

func ClientRegister(ctx context.Context, client *redis.Client) error {
	o, err := NewOptions()
	if err != nil {
		return err
	}
	health := NewClientHealthWithOptions(o)
	return health.Register(ctx, client)
}
