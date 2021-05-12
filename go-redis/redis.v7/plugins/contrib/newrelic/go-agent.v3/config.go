package newrelic

import (
	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/go-redis/redis.v7"
)

const (
	root    = redis.PluginsRoot + ".newrelic"
	enabled = root + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable redis integration")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
