package zerolog

import (
	"github.com/americanas-go/log"
	"github.com/americanas-go/log/contrib/rs/zerolog.v1"
)

// NewLogger returns a new  zerolog logger.
func NewLogger() log.Logger {
	options, err := NewOptions()
	if err != nil {
		panic(err)
	}
	return zerolog.NewLoggerWithOptions(options)
}
