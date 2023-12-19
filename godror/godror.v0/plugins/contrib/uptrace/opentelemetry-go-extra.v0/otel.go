package otelsql

import (
	"context"
	"database/sql"
	"github.com/americanas-go/ignite/go.opentelemetry.io/otel.v1"
	"github.com/americanas-go/log"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
)

// Register registers a new otel plugin on sql DB.
func Register(ctx context.Context, db *sql.DB) error {
	o, err := NewOptions()
	if err != nil {
		return nil
	}
	h := NewOTelWithOptions(o)
	return h.Register(ctx, db)
}

// OTel represents otel plugin for go driver for oracle.
type OTel struct {
	options *Options
}

// NewOTelWithOptions returns a new otel with options.
func NewOTelWithOptions(options *Options) *OTel {
	return &OTel{options: options}
}

// NewOTelWithConfigPath returns a new otel with options from config path.
func NewOTelWithConfigPath(path string) (*OTel, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewOTelWithOptions(o), nil
}

// NewOTel returns a new otel plugin.
func NewOTel() *OTel {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewOTelWithOptions(o)
}

// Register registers this otel plugin on sql DB.
func (i *OTel) Register(ctx context.Context, db *sql.DB) error {
	if !i.options.Enabled || !otel.IsTracerEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("integrating oracle in otel")

	otelsql.ReportDBStatsMetrics(db)

	logger.Debug("otel successfully integrated in oracle")

	return nil
}
