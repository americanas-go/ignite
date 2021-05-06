package nats

import (
	"time"

	"github.com/americanas-go/config"
	"github.com/nats-io/nats.go"
)

const (
	root          = "ignite.nats"
	maxReconnects = root + ".maxReconnects"
	reconnectWait = root + ".reconnectWait"
	url           = root + ".url"
	ExtRoot       = root + ".ext"
)

func init() {
	config.Add(maxReconnects, 1000, "define max reconnects to nats server")
	config.Add(reconnectWait, 1*time.Second, "define reconnects waiting before reconnect to nats server")
	config.Add(url, nats.DefaultURL, "define nats server url")
}
