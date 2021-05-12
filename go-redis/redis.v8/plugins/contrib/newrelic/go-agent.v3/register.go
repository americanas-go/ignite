package newrelic

import (
	"context"

	"github.com/americanas-go/log"
	"github.com/go-redis/redis/v8"
	newrelic "github.com/jvitoroc/ignite/newrelic/go-agent.v3"
)

func Register(ctx context.Context, client *redis.Client) error {

	if !IsEnabled() || !newrelic.IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("integrating redis in newrelic")

	client.AddHook(NewHook(client.Options()))

	logger.Debug("redis successfully integrated in newrelic")

	return nil
}
