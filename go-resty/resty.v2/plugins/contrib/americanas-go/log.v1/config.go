package log

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/go-resty/resty.v2"
)

const (
	root    = resty.PluginsRoot + ".log"
	enabled = root + ".enabled"
	level   = root + ".level"
)

func init() {
	config.Add(enabled, true, "enable/disable logger")
	config.Add(level, "DEBUG", "sets log level INFO/DEBUG/TRACE")
}

// IsEnabled returns config value from key ignite.resty.plugins.log.enabled where default is true.
func IsEnabled() bool {
	return config.Bool(enabled)
}

// Level returns config value from key ignite.resty.plugins.log.level where default is DEBUG.
func Level() string {
	return config.String(level)
}
