package datadog

import (
	"context"

	datadog "github.com/americanas-go/ignite/datadog/dd-trace-go.v1"
	"github.com/americanas-go/ignite/go.mongodb.org/mongo-driver.v1"
	"github.com/americanas-go/log"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func Register(ctx context.Context) (mongo.ClientOptionsPlugin, mongo.ClientPlugin) {

	if !IsEnabled() || !datadog.IsEnabled() {
		return nil, nil
	}

	return func(ctx context.Context, options *options.ClientOptions) error {
		logger := log.FromContext(ctx)

		logger.Trace("integrating mongo in datadog")

		options.SetMonitor(NewMonitor(
			tracer.ServiceName(datadog.Service()),
			tracer.AnalyticsRate(datadog.AnalyticsRate()),
		))

		logger.Debug("mongo successfully integrated in datadog")

		return nil
	}, nil
}
