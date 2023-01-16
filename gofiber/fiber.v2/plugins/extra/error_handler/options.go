package error_handler

import "github.com/americanas-go/ignite"

// Options error handler plugin for fiber options.
type Options struct {
	Enabled bool
	Logger  struct {
		Print4xx        bool
		Print5xx        bool
		PrintStackTrace bool
	}
}

// NewOptions returns options from config file or environment vars.
func NewOptions() (*Options, error) {
	return ignite.NewOptionsWithPath[Options](root)
}

// NewOptionsWithPath unmarshals a given key path into options and returns it.
func NewOptionsWithPath(path string) (opts *Options, err error) {
	return ignite.NewOptionsWithPath[Options](root, path)
}

return ignite.NewOptionsWithPath[Options](root, path)
}
