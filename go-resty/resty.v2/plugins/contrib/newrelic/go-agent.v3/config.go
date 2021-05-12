package newrelic

import (
	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/go-resty/resty.v2"
)

const (
	ConfigRoot = resty.PluginsRoot + ".newrelic"
	enabled    = ConfigRoot + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable newrelic integration")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
