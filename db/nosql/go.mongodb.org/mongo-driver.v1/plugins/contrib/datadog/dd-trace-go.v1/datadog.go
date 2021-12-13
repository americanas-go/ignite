package datadog

import (
	"context"

	datadog "github.com/americanas-go/ignite/apm/datadog/dd-trace-go.v1"
	"github.com/americanas-go/ignite/db/nosql/go.mongodb.org/mongo-driver.v1"
	"github.com/americanas-go/log"
	"go.mongodb.org/mongo-driver/mongo/options"
	mongotrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/go.mongodb.org/mongo-driver/mongo"
)

//Datadog represents a datadog plugin for mongo.
type Datadog struct {
	options *Options
}

// NewDatadogWithConfigPath returns a new datadog plugin with options from config path.
func NewDatadogWithConfigPath(path string, options ...mongotrace.Option) (*Datadog, error) {
	o, err := NewOptionsWithPath(path, options...)
	if err != nil {
		return nil, err
	}
	return NewDatadogWithOptions(o), nil
}

// NewDatadogWithOptions returns a new datadog plugin with options.
func NewDatadogWithOptions(options *Options) *Datadog {
	return &Datadog{options: options}
}

// NewDatadog returns a new datadog plugin with default options.
func NewDatadog(traceOptions ...mongotrace.Option) *Datadog {
	o, err := NewOptions(traceOptions...)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewDatadogWithOptions(o)
}

// Register registers this datadog plugin on a new mongo client.
func (d *Datadog) Register(ctx context.Context) (mongo.ClientOptionsPlugin, mongo.ClientPlugin) {
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

// Register registers a new datadog plugin on a new mongo client.
func Register(ctx context.Context) (mongo.ClientOptionsPlugin, mongo.ClientPlugin) {
	o, err := NewOptions()
	if err != nil {
		return nil, nil
	}
	datadog := NewDatadogWithOptions(o)
	return datadog.Register(ctx)
}
