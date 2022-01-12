package ignite

import (
	"context"
)

// represents an ignite instance. It could be an instance of a client, connection, server, and so on.
type IgniteInstance[O IgniteOptions] interface {
	// initializes instance
	Init(context.Context, O) error
}

// sets up a ignite instance with default options.
func Setup[I IgniteInstance[O], O IgniteOptions](ctx context.Context, plugins ...Plugin[I, O]) (i I, e error) {
	// load options
	o, e := Load[O]()
	if e != nil {
		return
	}
	return SetupWithOptions(ctx, o, plugins...)
}

// sets up a ignite instance with options from config path.
func SetupWithConfigPath[I IgniteInstance[O], O IgniteOptions](ctx context.Context, path string, plugins ...Plugin[I, O]) (i I, e error) {
	// load options
	o, e := LoadWithPath[O](path)
	if e != nil {
		return
	}
	return SetupWithOptions(ctx, o, plugins...)
}

// // sets up a ignite instance with options.
func SetupWithOptions[I IgniteInstance[O], O IgniteOptions](ctx context.Context, o O, plugins ...Plugin[I, O]) (i I, e error) {
	// setting up ignite instance
	i = New[I]()
	//initialize instance
	e = i.Init(ctx, o)
	// register plugins
	for _, p := range plugins {
		if e = p(ctx, i); e != nil {
			return
		}
	}
	// initalize
	return
}
