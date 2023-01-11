package hystrix

import (
	"strings"

	"github.com/americanas-go/config"
)

type Options struct {
	Enabled                bool
	Timeout                int
	RequestVolumeThreshold int
	ErrorPercentThreshold  int
	MaxConcurrentRequests  int
	SleepWindow            int
}

// NewOptionsFromCommand unmarshals options based a given key path.
func NewOptionsFromCommand(cmd string) (opts *Options, err error) {
	opts = new(Options)
	path := strings.Join([]string{cmdRoot, cmd}, ".")
	return NewOptionsWithPath(path)
}

// NewOptionsWithPath unmarshals options based a given key path.
func NewOptionsWithPath(path string) (opts *Options, err error) {
	opts = new(Options)

	err = config.UnmarshalWithPath(path, opts)
	if err != nil {
		return nil, err
	}

	return opts, nil
}
