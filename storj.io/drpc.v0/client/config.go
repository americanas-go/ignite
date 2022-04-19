package client

import (
	"github.com/americanas-go/config"
)

const (
	root        = "ignite.drpc.client"
	PluginsRoot = root + ".plugins"
	host        = ".host"
	port        = ".port"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+host, "localhost", "defines host")
	config.Add(path+port, 9091, "defines port")
}
