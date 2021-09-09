package datadog

import (
	"context"
	"sync"

	"github.com/americanas-go/ignite/net/http/client"
	"github.com/americanas-go/log"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// StartTracer starts the tracer like StartTraceWithOptions but with default Options.
func StartTracer(ctx context.Context, startOptions ...tracer.StartOption) {

	o, err := NewOptions()
	if err != nil {
		panic(err)
	}

	StartTracerWithOptions(ctx, o, startOptions...)
}

var tracerOnce sync.Once

// StartTracerWithOptions starts the tracer with the given set of options. It will stop and replace any running tracer, meaning that calling it several times will result in a restart of the tracer by replacing the current instance with a new one.
func StartTracerWithOptions(ctx context.Context, options *Options, startOptions ...tracer.StartOption) {

	if !IsTracerEnabled() {
		return
	}

	tracerOnce.Do(func() {

		logger := log.FromContext(ctx)

		httpClient := client.NewClientWithOptions(ctx, &options.HttpClient)

		so := []tracer.StartOption{
			tracer.WithAgentAddr(options.Addr),
			tracer.WithEnv(options.Env),
			tracer.WithService(options.Service),
			tracer.WithServiceVersion(options.Version),
			tracer.WithLogger(NewLogger()),
			tracer.WithHTTPClient(httpClient),
			tracer.WithAnalytics(options.Analytics),
			tracer.WithAnalyticsRate(options.AnalyticsRate),
			tracer.WithLambdaMode(options.LambdaMode),
			tracer.WithDebugMode(options.DebugMode),
			tracer.WithDebugStack(options.DebugStack),
		}

		for k, v := range options.Tags {
			so = append(so, tracer.WithGlobalTag(k, v))
		}

		so = append(so, startOptions...)

		tracer.Start(so...)

		logger.Infof("started a datadog tracer: %s", options.Service)
	})

}
