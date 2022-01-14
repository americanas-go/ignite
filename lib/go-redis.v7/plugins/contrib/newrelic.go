package contrib

import (
	"context"

	"github.com/americanas-go/ignite/lib/go-redis.v7/redis"
	newrelic "github.com/americanas-go/ignite/newrelic/go-agent.v3"
	"github.com/americanas-go/log"
	"github.com/newrelic/go-agent/v3/integrations/nrredis-v7"
)

func Newrelic(ctx context.Context, w *redis.Wrapper) error {
	o := w.Options.Plugins.Newrelic
	if !o.Enabled || !newrelic.IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("integrating redis in newrelic")
	w.UniversalClient().AddHook(nrredis.NewHook(w.RedisOptions()))
	logger.Debug("redis successfully integrated in newrelic")

	return nil
}
