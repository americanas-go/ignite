package datadog

import (
	"context"
	"database/sql"

	datadog "github.com/americanas-go/ignite/datadog/dd-trace-go.v1"
	"github.com/americanas-go/log"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
)

func Register(ctx context.Context, db *sql.DB) error {
	o, err := NewOptions()
	if err != nil {
		return nil
	}
	h := NewDatadogWithOptions(o)
	return h.Register(ctx, db)
}

type Datadog struct {
	options *Options
}

func NewDatadogWithOptions(options *Options) *Datadog {
	return &Datadog{options: options}
}

func NewDatadogWithConfigPath(path string, traceOptions ...sqltrace.Option) (*Datadog, error) {
	o, err := NewOptionsWithPath(path, traceOptions...)
	if err != nil {
		return nil, err
	}
	return NewDatadogWithOptions(o), nil
}

func NewDatadog(traceOptions ...sqltrace.Option) *Datadog {
	o, err := NewOptions(traceOptions...)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewDatadogWithOptions(o)
}

func (i *Datadog) Register(ctx context.Context, db *sql.DB) error {
	if !i.options.Enabled || !datadog.IsTracerEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("integrating oracle in datadog")

	sqltrace.Register("godror", db.Driver(), i.options.TraceOptions...)

	logger.Debug("datadog successfully integrated in oracle")

	return nil

}
