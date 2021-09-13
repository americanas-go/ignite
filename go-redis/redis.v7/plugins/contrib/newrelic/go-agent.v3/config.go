package newrelic

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/go-redis/redis.v7"
)

const (
	root    = redis.PluginsRoot + ".newrelic"
	enabled = root + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable redis integration")
}

// IsEnabled returns config value from key ignite.redis.plugins.newrelic.enabled where default is true.
func IsEnabled() bool {
	return config.Bool(enabled)
}
