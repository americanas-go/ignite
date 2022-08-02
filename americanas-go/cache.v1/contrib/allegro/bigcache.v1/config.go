package bigcache

import "github.com/americanas-go/ignite/allegro/bigcache.v1"

const (
	root = "ignite.cache.bigcache"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	bigcache.ConfigAdd(root)
}
