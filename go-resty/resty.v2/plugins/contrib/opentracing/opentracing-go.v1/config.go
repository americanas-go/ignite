package opentracing

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/go-resty/resty.v2"
)

const (
	root    = resty.PluginsRoot + ".opentracing"
	enabled = root + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable opentracing")
}

// IsEnabled returns config value from key ignite.resty.plugins.opentracing.enabled where default is true.
func IsEnabled() bool {
	return config.Bool(enabled)
}
