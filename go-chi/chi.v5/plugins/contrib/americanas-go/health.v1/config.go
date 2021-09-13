package health

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/go-chi/chi.v5"
)

const (
	root    = chi.PluginsRoot + ".health"
	enabled = root + ".enabled"
	route   = root + ".route"
)

func init() {
	config.Add(enabled, true, "enable/disable health route")
	config.Add(route, "/health", "define status url")
}

// IsEnabled returns config value from key ignite.chi.plugins.health.enabled where default is true.
func IsEnabled() bool {
	return config.Bool(enabled)
}

func getRoute() string {
	return config.String(route)
}
