package grapper

import "github.com/americanas-go/config"

const (
	root        = "ignite.grapper"
	PluginsRoot = root + ".plugins"
	name        = root + ".name"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+name, "default", "defines default wrapper name")
}
