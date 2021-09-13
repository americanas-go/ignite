package stripslashes

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/go-chi/chi.v5"
)

const (
	root    = chi.PluginsRoot + ".stripslashes"
	enabled = root + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable stripSlashes middleware")
}

// IsEnabled returns config value from key ignite.chi.plugins.stripslashes.enabled where default is true.
func IsEnabled() bool {
	return config.Bool(enabled)
}
