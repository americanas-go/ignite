package datadog

import (
	"context"

	"github.com/americanas-go/log"
	datadog "github.com/jvitoroc/ignite/datadog/dd-trace-go.v1"
	"github.com/jvitoroc/ignite/go.mongodb.org/mongo-driver.v1"
	"go.mongodb.org/mongo-driver/mongo/options"
	mongotrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/go.mongodb.org/mongo-driver/mongo"
)

func Register(ctx context.Context) (mongo.ClientOptionsPlugin, mongo.ClientPlugin) {

	if !IsEnabled() || !datadog.IsEnabled() {
		return nil, nil
	}

	return func(ctx context.Context, options *options.ClientOptions) error {
		logger := log.FromContext(ctx)

		logger.Trace("integrating mongo in datadog")

		options.SetMonitor(mongotrace.NewMonitor(mongotrace.WithServiceName(datadog.Service())))

		logger.Debug("mongo successfully integrated in datadog")

		return nil
	}, nil
}
