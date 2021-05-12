package newrelic

import (
	"context"

	"github.com/americanas-go/log"
	"github.com/go-redis/redis/v7"
	newrelic "github.com/jvitoroc/ignite/newrelic/go-agent.v3"
	"github.com/newrelic/go-agent/v3/integrations/nrredis-v7"
)

func Register(ctx context.Context, client *redis.Client) error {

	if !IsEnabled() || !newrelic.IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("integrating redis in newrelic")

	client.AddHook(nrredis.NewHook(client.Options()))

	logger.Debug("redis successfully integrated in newrelic")

	return nil
}
