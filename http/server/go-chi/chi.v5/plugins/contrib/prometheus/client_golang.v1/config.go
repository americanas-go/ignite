package prometheus

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/http/server/go-chi/chi.v5"
)

const (
	root    = chi.PluginsRoot + ".prometheus"
	enabled = ".enabled"
	route   = ".route"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+enabled, true, "enable/disable prometheus integration")
	config.Add(path+route, "/metrics", "define prometheus metrics url")
}
