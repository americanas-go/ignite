package log

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/go-chi/chi.v5"
)

const (
	root    = chi.PluginsRoot + ".logger"
	enabled = root + ".enabled"
	level   = root + ".level"
)

func init() {
	config.Add(enabled, true, "enable/disable logger middleware")
	config.Add(level, "INFO", "sets log level INFO/DEBUG/TRACE")
}

// IsEnabled returns config value from key ignite.chi.plugins.logger.enabled where default is true.
func IsEnabled() bool {
	return config.Bool(enabled)
}

// Level returns config value from key ignite.chi.plugins.logger.level where default is INFO.
func Level() string {
	return config.String(level)
}
