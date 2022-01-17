package redis

import (
	"context"
	"os"
	"reflect"
	"strings"

	"github.com/americanas-go/ignite"
	"github.com/americanas-go/ignite/lib/go-redis.v7/plugins"
	iredis "github.com/americanas-go/ignite/lib/go-redis.v7/redis"
	"github.com/go-redis/redis/v7"
)

// creates a new resty client or cluster client with default options.
func New[R redis.UniversalClient](ctx context.Context) (R, error) {
	if enableCluster[R]() {
		setEnv(iredis.OptionsRoot)
	}

	w, e := ignite.Setup(ctx, plugins.All...)
	return w.UniversalClient().(R), e
}

// creates a new resty client with options from config path.
func NewWithConfigPath[R redis.UniversalClient](ctx context.Context, path string) (R, error) {
	if enableCluster[R]() {
		setEnv(path)
	}
	w, e := ignite.SetupWithConfigPath(ctx, path, plugins.All...)
	return w.UniversalClient().(R), e
}

// creates a new resty client with options.
func NewWithOptions[R redis.UniversalClient](ctx context.Context, o *iredis.Options) (R, error) {
	o.Cluster.Enabled = enableCluster[R]()
	w, e := ignite.SetupWithOptions(ctx, o, plugins.All...)
	return w.UniversalClient().(R), e
}

// creates a new resty client options with values from default path.
func NewOptions() (*iredis.Options, error) {
	return ignite.Load[*iredis.Options]()
}

// creates a new resty client options with values from path.
func NewOptionsWithPath(path string) (*iredis.Options, error) {
	return ignite.LoadWithPath[*iredis.Options](path)
}

func enableCluster[R redis.UniversalClient]() bool {
	var r R
	switch reflect.TypeOf(r).String() {
	case "*redis.ClusterClient":
		return true
	}
	return false
}

func setEnv(path string) {
	os.Setenv(strings.ToUpper(strings.ReplaceAll(path, " ", "_"))+"_CLUSTER_ENABLED", "true")
}
