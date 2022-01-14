package redis

import (
	"context"

	"github.com/americanas-go/ignite"
	"github.com/americanas-go/ignite/lib/go-redis.v7/plugins"
	iredis "github.com/americanas-go/ignite/lib/go-redis.v7/redis"
	"github.com/go-redis/redis/v7"
)

// creates a new resty client or cluster client with default options.
func New[R redis.UniversalClient](ctx context.Context) (r R, e error) {
	w, e := ignite.Setup(ctx, plugins.All...)
	r = w.UniversalClient().(R)
	return
}

// creates a new resty client with options from config path.
func NewWithConfigPath[R redis.UniversalClient](ctx context.Context, path string) (r R, e error) {
	w, e := ignite.SetupWithConfigPath(ctx, path, plugins.All...)
	r = w.UniversalClient().(R)
	return
}

// creates a new resty client with options.
func NewWithOptions[R redis.UniversalClient](ctx context.Context, o *iredis.Options) (r R, e error) {
	w, e := ignite.SetupWithOptions(ctx, o, plugins.All...)
	r = w.UniversalClient().(R)
	return
}

// creates a new resty client options with values from default path.
func NewOptions() (*iredis.Options, error) {
	return ignite.Load[*iredis.Options]()
}

// creates a new resty client options with values from path.
func NewOptionsWithPath(path string) (*iredis.Options, error) {
	return ignite.LoadWithPath[*iredis.Options](path)
}
