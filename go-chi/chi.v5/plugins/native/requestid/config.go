package requestid

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/go-chi/chi.v5"
)

const (
	enabled = chi.PluginsRoot + ".requestid.enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable requestid middleware")
}

// IsEnabled returns config value from key ignite.chi.plugins.requestid.enabled where default is true.
func IsEnabled() bool {
	return config.Bool(enabled)
}
