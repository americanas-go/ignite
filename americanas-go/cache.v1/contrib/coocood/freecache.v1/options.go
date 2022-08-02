package freecache

import (
	"time"
)

// Options represents cache options.
type Options struct {
	TTL time.Duration `config:"ttl"`
}
