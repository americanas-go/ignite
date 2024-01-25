package otel

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/go-chi/chi.v5"
)

const (
	root       = chi.PluginsRoot + ".otel"
	enabled    = ".enabled"
	serverName = ".serverName"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+enabled, true, "enable/disable otel integration")
	config.Add(path+serverName, "my-server", "define otel server name")
}
