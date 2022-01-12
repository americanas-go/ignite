package ignite

import "context"

// represents a ignite plugin which can be applied over an ignite instance.
type Plugin[I IgniteInstance[O], O IgniteOptions] func(context.Context, I) error
