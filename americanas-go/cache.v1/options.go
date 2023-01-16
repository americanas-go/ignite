package grapper

import (
	"github.com/americanas-go/ignite"
)

// Options struct which represents cors plugin from chi options.
type Options struct {
	Name bool
}

// NewOptions returns options from config file or environment vars.
func NewOptions() (*Options, error) {
	return ignite.NewOptionsWithPath[Options](root)
}
