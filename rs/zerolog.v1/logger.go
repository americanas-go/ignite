package zerolog

import (
	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/rs/zerolog.v1"
)

func NewLogger() log.Logger {
	options, err := NewOptions()
	if err != nil {
		panic(err)
	}
	return zerolog.NewLoggerWithOptions(options)
}
