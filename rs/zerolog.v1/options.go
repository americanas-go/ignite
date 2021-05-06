package zerolog

import (
	"github.com/americanas-go/config"
	"github.com/americanas-go/log/contrib/rs/zerolog.v1"
)

func NewOptions() (*zerolog.Options, error) {
	o := &zerolog.Options{}

	err := config.UnmarshalWithPath(root, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}
