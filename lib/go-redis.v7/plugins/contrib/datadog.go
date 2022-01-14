package contrib

import (
	"context"

	datadog "github.com/americanas-go/ignite/datadog/dd-trace-go.v1"
	iredis "github.com/americanas-go/ignite/lib/go-redis.v7/redis"
	"github.com/americanas-go/log"
	redistrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/go-redis/redis.v7"
)

func Datadog(ctx context.Context, w *iredis.Wrapper) error {

	o := w.Options.Plugins.Datadog

	if !o.Enabled || !datadog.IsTracerEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("integrating redis in datadog")
	traceOpts := []redistrace.ClientOption{}
	if o.ServiceName != "" {
		traceOpts = append(traceOpts, redistrace.WithServiceName(o.ServiceName))
	}
	if o.AnalyticsRate != -1 {
		traceOpts = append(traceOpts, redistrace.WithAnalyticsRate(o.AnalyticsRate))
	}

	redistrace.WrapClient(w.UniversalClient(), traceOpts...)

	logger.Debug("redis successfully integrated in datadog")

	return nil
}
