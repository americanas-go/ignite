package freecache

import (
	"time"

	"github.com/americanas-go/config"
)

const (
	root      = "ignite.freecache"
	cacheSize = ".cacheSize"
	ttl       = ".ttl"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+cacheSize, 100*1024*1024, "The cache size will be set to 512KB at minimum")
	config.Add(path+ttl, 5*time.Minute, "The cache size will be set to 512KB at minimum")
}
