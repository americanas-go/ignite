package health

import (
	"context"

	"github.com/americanas-go/health"
	"github.com/americanas-go/log"
	"github.com/go-redis/redis/v7"
)

// ClusterHealth represents cluster health.
type ClusterHealth struct {
	options *Options
}

// NewClusterHealthWithOptions returns a health with the options provided.
func NewClusterHealthWithOptions(options *Options) *ClusterHealth {
	return &ClusterHealth{options: options}
}

// NewClusterHealth returns a health with default options.
func NewClusterHealth() *ClusterHealth {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewClusterHealthWithOptions(o)
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
