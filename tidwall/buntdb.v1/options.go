package buntdb

import "github.com/americanas-go/ignite"

type Options struct {
	Path       string
	SyncPolicy int
	AutoShrink struct {
		Percentage int
		MinSize    int
		Disabled   bool
	}
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
