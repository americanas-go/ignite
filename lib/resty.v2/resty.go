package resty

import (
	"context"

	"github.com/americanas-go/ignite"
	"github.com/americanas-go/ignite/lib/resty.v2/plugins"
	iresty "github.com/americanas-go/ignite/lib/resty.v2/resty"
	"github.com/go-resty/resty/v2"
)

// creates a new resty client with default options.
func New(ctx context.Context) (*resty.Client, error) {
	w, e := ignite.Setup(ctx, plugins.All...)
	if e != nil {
		return nil, e
	}
	return w.Instance, nil
}

// creates a new resty client with options from config path.
func NewWithConfigPath(ctx context.Context, path string) (*resty.Client, error) {
	w, e := ignite.SetupWithConfigPath(ctx, path, plugins.All...)
	if e != nil {
		return nil, e
	}
	return w.Instance, nil
}

// creates a new resty client with options.
func NewWithOptions(ctx context.Context, o *iresty.Options) (*resty.Client, error) {
	w, e := ignite.SetupWithOptions(ctx, o, plugins.All...)
	if e != nil {
		return nil, e
	}
	return w.Instance, nil
}

// creates a new resty client options with values from default path.
func NewOptions() (*iresty.Options, error) {
	return ignite.Load[*iresty.Options]()
}

// creates a new resty client options with values from path.
func NewOptionsWithPath(path string) (*iresty.Options, error) {
	return ignite.LoadWithPath[*iresty.Options](path)
}
