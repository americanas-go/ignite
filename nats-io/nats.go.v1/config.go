package nats

import (
	"github.com/americanas-go/config"
	"github.com/nats-io/nats.go"
)

const (
	root               = "ignite.nats"
	maxReconnects      = ".maxReconnects"
	reconnectWait      = ".reconnectWait"
	url                = ".url"
	reconnectJitter    = ".reconnectJitter"
	reconnectJitterTLS = ".reconnectJitterTLS"
	timeout            = ".timeout"
	pingInterval       = ".pingInterval"
	maxPingOut         = ".maxPingOut"
	maxChanLen         = ".maxChanLen"
	reconnectBufSize   = ".reconnectBufSize"
	drainTimeout       = ".drainTimeout"
	PluginsRoot        = root + ".plugins"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+maxReconnects, nats.DefaultMaxReconnect, "define max reconnects to nats server")
	config.Add(path+reconnectWait, nats.DefaultReconnectWait, "define reconnects waiting before reconnect to nats server")
	config.Add(path+url, nats.DefaultURL, "define nats server url")
	config.Add(path+reconnectJitter, nats.DefaultReconnectJitter, "")
	config.Add(path+reconnectJitterTLS, nats.DefaultReconnectJitterTLS, "")
	config.Add(path+timeout, nats.DefaultTimeout, "")
	config.Add(path+pingInterval, nats.DefaultPingInterval, "")
	config.Add(path+maxPingOut, nats.DefaultMaxPingOut, "")
	config.Add(path+maxChanLen, nats.DefaultMaxChanLen, "")
	config.Add(path+reconnectBufSize, nats.DefaultReconnectBufSize, "")
	config.Add(path+drainTimeout, nats.DefaultDrainTimeout, "")
}
