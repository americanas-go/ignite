package freecache

import "github.com/americanas-go/config"

const (
	root      = "ignite.freecache"
	cacheSize = root + ".cacheSize"
)

func init() {
	config.Add(cacheSize, 100*1024*1024, "The cache size will be set to 512KB at minimum")
}
