package nats

import (
	"time"

	"github.com/americanas-go/config"
	"github.com/nats-io/nats.go"
)

const (
	root          = "ignite.nats"
	maxReconnects = ".maxReconnects"
	reconnectWait = ".reconnectWait"
	url           = ".url"
	PluginsRoot   = root + ".plugins"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+maxReconnects, 1000, "define max reconnects to nats server")
	config.Add(path+reconnectWait, 1*time.Second, "define reconnects waiting before reconnect to nats server")
	config.Add(path+url, nats.DefaultURL, "define nats server url")
}
