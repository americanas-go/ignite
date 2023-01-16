package cors

import "github.com/americanas-go/ignite"

// Options cors plugin for fiber options.
type Options struct {
	Enabled bool
	Allowed struct {
		Origins     []string
		Headers     []string
		Methods     []string
		Credentials bool
	}
	Exposed struct {
		Headers []string
	}
	MaxAge int
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
