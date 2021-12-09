package multiserver

import (
	"context"
	"net/http"

	"github.com/americanas-go/ignite/go-chi/chi.v5"
	"github.com/americanas-go/log"
	"github.com/americanas-go/multiserver"
)

// Register registers multi server check router for chi.
func Register(ctx context.Context) (*chi.Config, error) {
	l := NewMultiServer()
	return l.Register(ctx)
}

// MultiServer struct that represents a multiserver.
type MultiServer struct {
	options *Options
}

// NewMultiServerWithOptions returns a new multiserver with options.
func NewMultiServerWithOptions(options *Options) *MultiServer {
	return &MultiServer{options: options}
}

// NewMultiServerWithConfigPath returns a new multiserver with options from config path.
func NewMultiServerWithConfigPath(path string) (*MultiServer, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewMultiServerWithOptions(o), nil
}

// NewMultiServer returns a new multiserver with default options.
func NewMultiServer() *MultiServer {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewMultiServerWithOptions(o)
}

// Register registers a multi server check router in a new chi config.
func (i *MultiServer) Register(ctx context.Context) (*chi.Config, error) {
	if !i.options.Enabled {
		return nil, nil
	}

	logger := log.FromContext(ctx)

	checkRoute := i.options.Route

	logger.Tracef("configuring multi server check router on %s in chi", checkRoute)

	return &chi.Config{
		Routes: []chi.ConfigRouter{
			{
				Method:      http.MethodGet,
				HandlerFunc: Get(ctx),
				Pattern:     checkRoute,
			},
		},
	}, nil
}

// Get returns http handler for multi server check.
func Get(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		status := http.StatusOK

		if err := multiserver.Check(ctx); err != nil {
			status = http.StatusServiceUnavailable
			w.WriteHeader(status)
			w.Write([]byte("service unavailable"))
			return
		}

		w.WriteHeader(status)

	}
}
