package multiserver

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/go-chi/chi.v5"
)

const (
	root    = chi.PluginsRoot + ".multiServer"
	enabled = root + ".enabled"
	route   = root + ".route"
)

func init() {
	config.Add(enabled, true, "enable/disable multi server check route")
	config.Add(route, "/check", "define multi server check url")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}

func getRoute() string {
	return config.String(route)
}
