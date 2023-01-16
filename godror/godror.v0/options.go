package godror

import (
	"time"

	"github.com/americanas-go/ignite"
)

// Options represents a godror options.
type Options struct {
	ConnectString       string
	Username            string
	Password            string
	MaxLifetime         time.Duration
	SessionTimeout      time.Duration
	WaitTimeout         time.Duration
	MaxSessions         int
	SessionIncrement    int
	MinSessions         int
	MaxSessionsPerShard int
}

// NewOptions returns options from config file or environment vars.
func NewOptions() (*Options, error) {
	return ignite.NewOptionsWithPath[Options](root)
}

// NewOptionsWithPath unmarshals options based a given key path.
func NewOptionsWithPath(path string) (opts *Options, err error) {

	opts, err = NewOptions()
	if err != nil {
		return nil, err
	}

	return ignite.MergeOptionsWithPath[Options](opts, path)
}
