package freecache

import (
	"time"

	"github.com/americanas-go/config"
)

const (
	root = "ignite.cache.freecache"
	ttl  = ".ttl"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+ttl, 5*time.Minute, "time after which entry can be evicted")
}
