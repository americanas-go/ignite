package datadog

import (
	"context"
	"net/http"

	"github.com/americanas-go/log"
	datadog "github.com/jvitoroc/ignite/datadog/dd-trace-go.v1"
	"github.com/jvitoroc/ignite/go-chi/chi.v5"
	c "gopkg.in/DataDog/dd-trace-go.v1/contrib/go-chi/chi.v5"
)

func Register(ctx context.Context) (*chi.Config, error) {
	if !IsEnabled() || !datadog.IsEnabled() {
		return nil, nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling datadog middleware in chi")

	return &chi.Config{
		Middlewares: []func(http.Handler) http.Handler{
			c.Middleware(c.WithServiceName(datadog.Service())),
		},
	}, nil
}
