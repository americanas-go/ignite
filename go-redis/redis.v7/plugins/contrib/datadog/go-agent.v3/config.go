package newrelic

import (
	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/go-redis/redis.v7"
)

const (
	root    = redis.PluginsRoot + ".datadog"
	enabled = root + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable datadog integration")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
