package newrelic

import (
	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/panjf2000/ants.v2"
)

const (
	root    = ants.PluginsRoot + ".newrelic"
	enabled = root + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable newrelic integration")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
