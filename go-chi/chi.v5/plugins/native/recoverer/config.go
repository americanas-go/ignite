package recoverer

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/go-chi/chi.v5"
)

const (
	enabled = chi.PluginsRoot + ".recover.enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable recover middleware")
}

// IsEnabled returns config value from key ignite.chi.plugins.recover.enabled where default is true.
func IsEnabled() bool {
	return config.Bool(enabled)
}
