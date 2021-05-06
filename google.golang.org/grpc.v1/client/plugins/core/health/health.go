package health

import (
	"context"

	"github.com/americanas-go/health"
	"github.com/americanas-go/log"
	"google.golang.org/grpc"
)

type Health struct {
	options *Options
}

func NewHealthWithOptions(options *Options) *Health {
	return &Health{options: options}
}

func NewHealth() *Health {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewHealthWithOptions(o)
}
func (i *Health) Register(ctx context.Context, conn *grpc.ClientConn) error {

	logger := log.FromContext(ctx).WithTypeOf(*i)

	logger.Trace("integrating grpc with health")

	checker := NewChecker(conn)
	hc := health.NewHealthChecker(i.options.Name, i.options.Description, checker, i.options.Required, i.options.Enabled)
	health.Add(hc)

	logger.Debug("grpc integrated on health with success")

	return nil
}
