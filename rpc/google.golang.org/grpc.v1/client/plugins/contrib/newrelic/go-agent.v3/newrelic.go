package newrelic

import (
	"context"

	newrelic "github.com/americanas-go/ignite/apm/newrelic/go-agent.v3"
	"github.com/americanas-go/log"
	"github.com/newrelic/go-agent/v3/integrations/nrgrpc"
	"google.golang.org/grpc"
)

// Register registers a new newrelic plugin for grpc client.
func Register(ctx context.Context) ([]grpc.DialOption, []grpc.CallOption) {
	o, err := NewOptions()
	if err != nil {
		return nil, nil
	}
	h := NewNewrelicWithOptions(o)
	return h.Register(ctx)
}

// Newrelic represents newrelic plugin for grpc client.
type Newrelic struct {
	options *Options
}

// NewNewrelicWithOptions returns a new newrelic plugin with options.
func NewNewrelicWithOptions(options *Options) *Newrelic {
	return &Newrelic{options: options}
}

// NewNewrelicWithConfigPath returns a new newrelic plugin with options from config path.
func NewNewrelicWithConfigPath(path string) (*Newrelic, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewNewrelicWithOptions(o), nil
}

// NewNewrelic returns a new newrelic plugin with default options.
func NewNewrelic() *Newrelic {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewNewrelicWithOptions(o)
}

// Register registers this newrelic plugin for grpc client.
func (i *Newrelic) Register(ctx context.Context) ([]grpc.DialOption, []grpc.CallOption) {

	if !i.options.Enabled || !newrelic.IsEnabled() {
		return nil, nil
	}

	logger := log.FromContext(ctx)
	logger.Debug("newrelic interceptor successfully enabled in grpc client")

	return []grpc.DialOption{
		grpc.WithChainUnaryInterceptor(nrgrpc.UnaryClientInterceptor),
		grpc.WithChainStreamInterceptor(nrgrpc.StreamClientInterceptor),
	}, nil

}
