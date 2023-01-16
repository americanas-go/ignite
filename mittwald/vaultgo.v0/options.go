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

// NewOptionsWithPath unmarshals a given key path into options and returns it.
func NewOptionsWithPath(path string) (opts *Options, err error) {
	return ignite.NewOptionsWithPath[Options](root, path)
}

return ignite.NewOptionsWithPath[Options](root, path)
}
