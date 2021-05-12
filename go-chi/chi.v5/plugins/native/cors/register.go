package cors

import (
	"context"
	"net/http"

	"github.com/americanas-go/log"
	"github.com/go-chi/cors"
	"github.com/jvitoroc/ignite/go-chi/chi.v5"
)

func Register(ctx context.Context) (*chi.Config, error) {

	if !IsEnabled() {
		return nil, nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("enabling cors middleware in chi")

	return &chi.Config{
		Middlewares: []func(http.Handler) http.Handler{
			cors.Handler(cors.Options{
				AllowedOrigins:   getAllowedOrigins(),
				AllowedMethods:   getAllowedMethods(),
				AllowedHeaders:   getAllowedHeaders(),
				AllowCredentials: getAllowedCredentials(),
				ExposedHeaders:   getExposedHeaders(),
				MaxAge:           getMaxAge(),
			}),
		},
	}, nil
}
