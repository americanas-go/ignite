package zerolog

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/log/contrib/rs/zerolog.v1"
)

// NewOptions returns options from config file or environment vars.
func NewOptions() (*zerolog.Options, error) {
	o := &zerolog.Options{}

	err := config.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
