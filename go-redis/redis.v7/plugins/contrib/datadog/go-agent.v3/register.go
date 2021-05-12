package newrelic

import (
	"context"

	"github.com/americanas-go/log"
	"github.com/go-redis/redis/v7"
	datadog "github.com/jvitoroc/ignite/datadog/dd-trace-go.v1"
	redistrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/go-redis/redis.v7"
)

func Register(ctx context.Context, client *redis.Client) error {

	if !IsEnabled() || !datadog.IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("integrating redis in datadog")

	redistrace.WrapClient(client)

	logger.Debug("redis successfully integrated in datadog")

	return nil
}
