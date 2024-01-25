package contrib // import "github.com/americanas-go/ignite/labstack/echo.v4/plugins/contrib/opentelemetry/otelecho.v1"

import (
	"github.com/americanas-go/ignite"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
)

// Options represents the opentelemetry plugin for echo server options.
type Options struct {
	Enabled        bool
	TracingOptions []otelecho.Option
}

// NewOptions returns options from config file or environment vars.
func NewOptions() (*Options, error) {
	return ignite.NewOptionsWithPath[Options](root)
}

// NewOptionsWithPath unmarshals a given key path into options and returns it.
func NewOptionsWithPath(path string) (opts *Options, err error) {
	return ignite.NewOptionsWithPath[Options](root, path)
}
