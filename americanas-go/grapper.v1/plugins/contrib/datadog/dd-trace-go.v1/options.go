package datadog

import (
	"strings"

	"github.com/americanas-go/ignite"
)

// Options struct which represents cors plugin from chi options.
type Options struct {
	Enabled bool
}

// NewOptions returns options from config path.
func NewOptions(name string) (opts *Options, err error) {
	opts = &Options{}
	path := strings.Join([]string{root, ".", name}, "")
	return ignite.NewOptionsWithPath[Options](path)
}
