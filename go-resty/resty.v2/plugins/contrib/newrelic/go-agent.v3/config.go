package newrelic

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/go-resty/resty.v2"
)

const (
	ConfigRoot = resty.PluginsRoot + ".newrelic"
	enabled    = ConfigRoot + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable newrelic integration")
}

// IsEnabled returns config value from key ignite.resty.plugins.newrelic.enabled where default is true.
func IsEnabled() bool {
	return config.Bool(enabled)
}
