package vault

import "github.com/americanas-go/ignite"

// Options vault client options.
type Options struct {
	Addr   string
	Type   string
	CaPath string
	Token  string
	K8s    struct {
		Role string
		Jwt  struct {
			File    string
			Content string
		}
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
