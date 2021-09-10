package datadog

import (
	"context"

	datadog "github.com/americanas-go/ignite/datadog/dd-trace-go.v1"
	"github.com/americanas-go/ignite/go.mongodb.org/mongo-driver.v1"
	"github.com/americanas-go/log"
	"go.mongodb.org/mongo-driver/mongo/options"
	mongotrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/go.mongodb.org/mongo-driver/mongo"
)

type DataDog struct {
	options *Options
}

func NewDatadogWithConfigPath(path string) (*DataDog, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewDatadogWithOptions(o), nil
}

func NewDatadogWithOptions(options *Options) *DataDog {
	return &DataDog{options: options}
}

func (d *DataDog) Register(ctx context.Context) (mongo.ClientOptionsPlugin, mongo.ClientPlugin) {
	if !d.options.Enabled || !datadog.IsTracerEnabled() {
		return nil, nil
	}

	return func(ctx context.Context, options *options.ClientOptions) error {
		logger := log.FromContext(ctx)

		logger.Trace("integrating datadog in mongo")

		options.SetMonitor(mongotrace.NewMonitor(d.options.Options...))

		logger.Debug("datadog successfully integrated in mongo")

		return nil
	}, nil
}

func Register(ctx context.Context) (mongo.ClientOptionsPlugin, mongo.ClientPlugin) {
	o, err := NewOptions()
	if err != nil {
		return nil, nil
	}
	datadog := NewDatadogWithOptions(o)
	return datadog.Register(ctx)
}
