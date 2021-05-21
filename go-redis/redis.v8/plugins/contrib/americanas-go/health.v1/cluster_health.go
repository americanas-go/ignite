package health

import (
	"context"

	"github.com/americanas-go/health"
	"github.com/americanas-go/log"
	"github.com/go-redis/redis/v8"
)

type ClusterHealth struct {
	options *Options
}

func NewClusterHealthWithOptions(options *Options) *ClusterHealth {
	return &ClusterHealth{options: options}
}

func NewClusterHealth() *ClusterHealth {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewClusterHealthWithOptions(o)
}

func (i *ClusterHealth) Register(ctx context.Context, client *redis.ClusterClient) error {

	logger := log.FromContext(ctx).WithTypeOf(*i)

	logger.Trace("integrating redis in health")

	checker := NewClusterClientChecker(client)
	hc := health.NewHealthChecker(i.options.Name, i.options.Description, checker, i.options.Required, i.options.Enabled)
	health.Add(hc)

	logger.Debug("redis successfully integrated in health")

	return nil
}
