package datadog

import (
	"context"

	datadog "github.com/americanas-go/ignite/datadog/dd-trace-go.v1"
	"github.com/americanas-go/log"
	"github.com/aws/aws-sdk-go-v2/aws"
	awstrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/aws/aws-sdk-go-v2/aws"
)

type Datadog struct {
	options *Options
}

func NewDatadogWithConfigPath(path string, traceOptions ...awstrace.Option) (*Datadog, error) {
	o, err := NewOptionsWithPath(path, traceOptions...)
	if err != nil {
		return nil, err
	}
	return NewDatadogWithOptions(o), nil
}

func NewDatadogWithOptions(options *Options) *Datadog {
	return &Datadog{options: options}
}

func NewDatadog(traceOptions ...awstrace.Option) *Datadog {
	o, err := NewOptions(traceOptions...)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewDatadogWithOptions(o)
}

func Register(ctx context.Context, awsCfg *aws.Config) error {
	o, err := NewOptions()
	if err != nil {
		return err
	}

	plugin := NewDatadogWithOptions(o)
	return plugin.Register(ctx, awsCfg)
}

// Register registers AWS tracing for Datadog.
func (i *Datadog) Register(ctx context.Context, awsCfg *aws.Config) error {

	if !i.options.Enabled || !datadog.IsTracerEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling datadog middleware in aws")

	awstrace.AppendMiddleware(awsCfg, i.options.TraceOptions...)

	logger.Debug("datadog middleware successfully enabled in aws")

	return nil

}