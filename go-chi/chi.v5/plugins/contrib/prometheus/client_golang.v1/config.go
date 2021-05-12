package prometheus

import (
	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/go-chi/chi.v5"
)

const (
	root    = chi.PluginsRoot + ".prometheus"
	enabled = root + ".enabled"
	route   = root + ".route"
)

func init() {
	config.Add(enabled, true, "enable/disable prometheus integration")
	config.Add(route, "/metrics", "define prometheus metrics url")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}

func getRoute() string {
	return config.String(route)
}
