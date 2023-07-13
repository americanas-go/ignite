package health

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/americanas-go/ignite/go-chi/chi.v5"
	"github.com/americanas-go/log"
	response "github.com/americanas-go/rest-response"
)

func Register(ctx context.Context) (*chi.Config, error) {
	o, err := NewOptions()
	if err != nil {
		return nil, err
	}
	health := NewHealthWithOptions(o)
	return health.Register(ctx)
}

type Health struct {
	options *Options
}

func NewHealthWithOptions(options *Options) *Health {
	return &Health{options: options}
}

func NewHealthWithConfigPath(path string) (*Health, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewHealthWithOptions(o), nil
}

func NewHealth() *Health {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewHealthWithOptions(o)
}

func (i *Health) Register(ctx context.Context) (*chi.Config, error) {
	if !i.options.Enabled {
		return nil, nil
	}

	logger := log.FromContext(ctx)

	logger.Tracef("configuring health router on %s in chi", i.options.Route)

	return &chi.Config{
		Routes: []chi.ConfigRouter{
			{
				Method:      http.MethodGet,
				HandlerFunc: handler(ctx),
				Pattern:     i.options.Route,
			},
		},
	}, nil
}

func handler(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, httpCode := response.NewHealth(ctx)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(httpCode)
		json.NewEncoder(w).Encode(resp)
	}
}
