package contrib

import (
	"context"

	"github.com/americanas-go/health"
	iredis "github.com/americanas-go/ignite/lib/go-redis.v7/redis"
	"github.com/americanas-go/log"
	"github.com/go-redis/redis/v7"
)

func Health(ctx context.Context, w *iredis.Wrapper) error {
	o := w.Options.Plugins.Health
	if !o.Enabled {
		return nil
	}
	logger := log.FromContext(ctx)

	logger.Trace("integrating redis in health")
	hc := health.NewHealthChecker(o.Name, o.Description, &checker{w.UniversalClient()}, o.Required, o.Enabled)
	health.Add(hc)

	logger.Debug("redis successfully integrated in health")

	return nil
}

type checker struct {
	client redis.UniversalClient
}

// Check checks if redis server is responding.
func (c *checker) Check(ctx context.Context) error {
	return c.client.Ping().Err()
}
