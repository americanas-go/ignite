package datadog

import (
	"context"
	"database/sql"
	"database/sql/driver"
	datadog "github.com/americanas-go/ignite/datadog/dd-trace-go.v1"
	"github.com/americanas-go/log"
	sqltrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/database/sql"
)

// Register registers a new datadog plugin on sql DB.
func Register(ctx context.Context, db *sql.DB, connector driver.Connector) (d *sql.DB, err error) {
	o, err := NewOptions()
	if err != nil {
		return nil, err
	}
	h := NewDatadogWithOptions(o)
	return h.Register(ctx, db, connector)
}

// Datadog represents datadog plugin for go driver for oracle.
type Datadog struct {
	options *Options
}

// NewDatadogWithOptions returns a new datadog with options.
func NewDatadogWithOptions(options *Options) *Datadog {
	return &Datadog{options: options}
}

// NewDatadogWithConfigPath returns a new datadog with options from config path.
func NewDatadogWithConfigPath(path string, traceOptions ...sqltrace.Option) (*Datadog, error) {
	o, err := NewOptionsWithPath(path, traceOptions...)
	if err != nil {
		return nil, err
	}
	return NewDatadogWithOptions(o), nil
}

// NewDatadog returns a new datadog plugin.
func NewDatadog(traceOptions ...sqltrace.Option) *Datadog {
	o, err := NewOptions(traceOptions...)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewDatadogWithOptions(o)
}

// Register registers this datadog plugin on sql DB.
func (i *Datadog) Register(ctx context.Context, db *sql.DB, connector driver.Connector) (d *sql.DB, err error) {
	if !i.options.Enabled || !datadog.IsTracerEnabled() {
		return nil, nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("integrating oracle in datadog")

	sqltrace.Register("godror", db.Driver(), i.options.TraceOptions...)

	logger.Debug("datadog successfully integrated in oracle")

	return db, nil

}
