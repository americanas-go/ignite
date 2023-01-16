package hystrix

import (
	"strings"

	"github.com/americanas-go/config"
)

// Options struct which represents cors plugin from chi options.
type Options struct {
	Enabled bool
}

// NewOptions returns options from config path.
func NewOptions(name string) (opts *Options, err error) {
	opts = &Options{}
	path := strings.Join([]string{root, ".", name}, "")
	err = config.UnmarshalWithPath(path, opts)
	if err != nil {
		return nil, err
	}

	return opts, nil
}
