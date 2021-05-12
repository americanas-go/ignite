package recoverer

import (
	"github.com/americanas-go/config"
	"github.com/jvitoroc/ignite/go-chi/chi.v5"
)

const (
	enabled = chi.PluginsRoot + ".recover.enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable recover middleware")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
