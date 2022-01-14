package plugins

import (
	"github.com/americanas-go/ignite"
	"github.com/americanas-go/ignite/lib/go-redis.v7/plugins/contrib"
	"github.com/americanas-go/ignite/lib/go-redis.v7/redis"
)

// all go-redis plugins
var All = []ignite.Plugin[*redis.Wrapper, *redis.Options]{
	contrib.Datadog, contrib.Health, contrib.Newrelic,
}
