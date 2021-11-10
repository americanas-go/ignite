package health

import (
	"context"

	"github.com/americanas-go/health"
	"github.com/americanas-go/log"
	"github.com/go-redis/redis/v8"
)

// ClusterHealth represents cluster health.
type ClusterHealth struct {
	options *Options
}

// NewClusterHealthWithOptions returns a health with the options provided.
func NewClusterHealthWithOptions(options *Options) *ClusterHealth {
	return &ClusterHealth{options: options}
}

// NewClusterHealthWithConfigPath returns a health with options from config path.
func NewClusterHealthWithConfigPath(path string) (*ClusterHealth, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewClusterHealthWithOptions(o), nil
}

// NewClusterHealth returns a health with default options.
func NewClusterHealth() (*ClusterHealth, error) {
	o, err := NewOptions()
	if err != nil {
		return nil, err
	}

	return NewClusterHealthWithOptions(o), nil
}

// Register registers a new ClusterClientChecker in the health package.
func (i *ClusterHealth) Register(ctx context.Context, client *redis.ClusterClient) error {

	logger := log.FromContext(ctx).WithTypeOf(*i)

	logger.Trace("integrating redis in health")

	checker := NewClusterClientChecker(client)
	hc := health.NewHealthChecker(i.options.Name, i.options.Description, checker, i.options.Required, i.options.Enabled)
	health.Add(hc)

	logger.Debug("redis successfully integrated in health")

	return nil
}

func ClusterRegister(ctx context.Context, client *redis.ClusterClient) error {
	o, err := NewOptions()
	if err != nil {
		return err
	}
	health := NewClusterHealthWithOptions(o)
	return health.Register(ctx, client)
}
