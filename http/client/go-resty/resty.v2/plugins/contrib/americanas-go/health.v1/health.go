package health

import (
	"context"

	"github.com/americanas-go/health"
	"github.com/americanas-go/log"
	"github.com/go-resty/resty/v2"
)

// Health represents resty health.
type Health struct {
	options *Options
}

// NewHealthWithConfigPath returns a health with the options from config path.
func NewHealthWithConfigPath(path string) (*Health, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewHealthWithOptions(o), nil
}

// NewHealthWithOptions returns a health with the options provided.
func NewHealthWithOptions(options *Options) *Health {
	return &Health{options: options}
}

// NewHealth returns a health with default options.
func NewHealth() *Health {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return NewHealthWithOptions(o)
}

// Register registers health plugin with default options to client
func Register(ctx context.Context, client *resty.Client) error {
	o, err := NewOptions()
	if err != nil {
		return err
	}

	plugin := NewHealthWithOptions(o)
	return plugin.Register(ctx, client)
}

// Register registers a new checker in the health package.
func (i *Health) Register(ctx context.Context, client *resty.Client) error {

	logger := log.FromContext(ctx).WithTypeOf(*i)

	logger.Trace("integrating resty in health")

	checker := NewChecker(client, i.options)
	hc := health.NewHealthChecker(i.options.Name, i.options.Description, checker, i.options.Required, i.options.Enabled)
	health.Add(hc)

	logger.Debug("resty successfully integrated in health")

	return nil
}
