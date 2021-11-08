package godror

import (
	"time"

	"github.com/americanas-go/config"
)

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

func NewOptions() (*Options, error) {
	o := &Options{}

	err := config.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}

func NewOptionsWithPath(path string) (opts *Options, err error) {

	opts, err = NewOptions()
	if err != nil {
		return nil, err
	}

	err = config.UnmarshalWithPath(path, opts)
	if err != nil {
		return nil, err
	}

	return opts, nil
}
