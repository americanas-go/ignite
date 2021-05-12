package opentracing

import (
	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/go-resty/resty.v2"
)

const (
	root    = resty.PluginsRoot + ".opentracing"
	enabled = root + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable opentracing")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
