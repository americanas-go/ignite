package logger

import (
	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/go-resty/resty.v2"
)

const (
	root    = resty.PluginsRoot + ".logger"
	enabled = root + ".enabled"
	level   = root + ".level"
)

func init() {
	config.Add(enabled, true, "enable/disable logger")
	config.Add(level, "DEBUG", "sets log level INFO/DEBUG/TRACE")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}

func Level() string {
	return config.String(level)
}
