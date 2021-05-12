package newrelic

import (
	"context"

	"github.com/americanas-go/log"
	newrelic "github.com/jvitoroc/ignite/newrelic/go-agent.v3"
	"github.com/newrelic/go-agent/v3/integrations/nrgrpc"
	"google.golang.org/grpc"
)

func Register(ctx context.Context) []grpc.ServerOption {

	if !IsEnabled() || !newrelic.IsEnabled() {
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
