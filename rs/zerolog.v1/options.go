package zerolog

import (
	"github.com/americanas-go/ignite"
	"github.com/americanas-go/log/contrib/rs/zerolog.v1"
)

// NewOptions returns options from config file or environment vars.
func NewOptions() (*zerolog.Options, error) {
	return ignite.NewOptionsWithPath[zerolog.Options](root)
}
