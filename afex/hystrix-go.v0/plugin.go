package hystrix

import (
	"context"
)

// Plugin defines a function to process plugin.
type Plugin func(context.Context) error
