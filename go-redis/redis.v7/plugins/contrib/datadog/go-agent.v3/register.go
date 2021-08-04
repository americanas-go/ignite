package newrelic

import (
	"context"

	datadog "github.com/americanas-go/ignite/datadog/dd-trace-go.v1"
	"github.com/americanas-go/log"
	"github.com/go-redis/redis/v7"
	redistrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/go-redis/redis.v7"
)

func Register(ctx context.Context, client *redis.Client) error {

	if !IsEnabled() || !datadog.IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("integrating redis in datadog")

	redistrace.WrapClient(client,
		redistrace.WithServiceName(datadog.Service()),
		redistrace.WithAnalyticsRate(datadog.AnalyticsRate()),
	)

	logger.Debug("redis successfully integrated in datadog")

	return nil
}
