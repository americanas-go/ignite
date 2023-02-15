package bigcache

import (
	"time"

	"github.com/americanas-go/config"
)

const (
	root               = "ignite.bigcache"
	shards             = ".shards"
	lifeWindow         = ".lifeWindow"
	cleanWindow        = ".cleanWindow"
	maxEntriesInWindow = ".maxEntriesInWindow"
	maxEntrySize       = ".maxEntrySize"
	verbose            = ".verbose"
	hardMaxCacheSize   = ".hardMaxCacheSize"
	statsEnabled       = ".statsEnabled"
)

func init() {
	ConfigAdd(root)
}

func ConfigAdd(path string) {
	config.Add(path+shards, 1024, "number of cache shards, value must be a power of two")
	config.Add(path+lifeWindow, 5*time.Minute, "time after which entry can be evicted")
	config.Add(path+cleanWindow, 0, "interval between removing expired entries (clean up). if set to <= 0 then no action is performed. Setting to < 1 second is counterproductive — bigcache has a one second resolution")
	config.Add(path+maxEntriesInWindow, 1000*10*60, "max number of entries in life window. Used only to calculate initial size for cache shards. when proper value is set then additional memory allocation does not occur.")
	config.Add(path+maxEntrySize, 1*1024*1024, "max size of entry in bytes. Used only to calculate initial size for cache shards")
	config.Add(path+verbose, false, "verbose mode prints information about new memory allocation")
	config.Add(path+statsEnabled, false, "if true calculate the number of times a cached resource was requested")
	config.Add(path+hardMaxCacheSize, 1, "hardMaxCacheSize is a limit for cache size in MB. Cache will not allocate more memory than this limit. when the limit is higher than 0 and reached then the oldest entries are overridden for the new ones")
}
