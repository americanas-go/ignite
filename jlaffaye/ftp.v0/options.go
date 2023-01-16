package ftp

import (
	"github.com/americanas-go/ignite"
)

// Options ftp server connection options.
type Options struct {
	Addr     string
	User     string
	Password string
	Timeout  int
	Retry    int
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
