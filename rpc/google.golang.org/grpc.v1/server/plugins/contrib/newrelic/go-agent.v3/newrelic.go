package newrelic

import (
	"context"

	newrelic "github.com/americanas-go/ignite/apm/newrelic/go-agent.v3"
	"github.com/americanas-go/log"
	"github.com/newrelic/go-agent/v3/integrations/nrgrpc"
	"google.golang.org/grpc"
)

// Register registers a new compressor plugin for grpc server.
// Compressor represents compressor plugin for grpc server.
// NewCompressorWithOptions returns a new compressor plugin with options.
// NewCompressorWithConfigPath returns a new compressor plugin with options from config path.
// NewCompressor returns a new compressor plugin with default options.
// Register registers this compressor plugin for grpc server.
func Register(ctx context.Context) []grpc.ServerOption {
	o, err := NewOptions()
	if err != nil {
		return nil
	}
	h := NewNewrelicWithOptions(o)
	return h.Register(ctx)
}

type Newrelic struct {
	options *Options
}

func NewNewrelicWithOptions(options *Options) *Newrelic {
	return &Newrelic{options: options}
}

func NewNewrelicWithConfigPath(path string) (*Newrelic, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewNewrelicWithOptions(o), nil
}

func NewNewrelic() *Newrelic {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewNewrelicWithOptions(o)
}

func (i *Newrelic) Register(ctx context.Context) []grpc.ServerOption {

	if !i.options.Enabled || !newrelic.IsEnabled() {
		return nil
	}

	app := newrelic.Application()

	logger := log.FromContext(ctx)
	logger.Debug("newrelic interceptor successfully enabled in grpc server")

	return []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(nrgrpc.UnaryServerInterceptor(app)),
		grpc.ChainStreamInterceptor(nrgrpc.StreamServerInterceptor(app)),
	}

}
