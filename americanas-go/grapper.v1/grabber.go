package grapper

import "github.com/americanas-go/grapper"

func NewWrapper[R any](plugins ...func(name string) grapper.Middleware[R]) (*grapper.Wrapper[R], error) {
	o, err := NewOptions()
	if err != nil {
		return nil, err
	}

	return NewWrapperWithOptions(o, plugins...)
}

func NewWrapperWithPath[R any](path string, plugins ...func(name string) grapper.Middleware[R]) (*grapper.Wrapper[R], error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}

	return NewWrapperWithOptions(o, plugins...)
}

func NewWrapperWithOptions[R any](options *Options, plugins ...func(name string) grapper.Middleware[R]) (*grapper.Wrapper[R], error) {
	var m []grapper.Middleware[R]

	for _, f := range plugins {
		if p := f(options.Name); p != nil {
			m = append(m, p)
		}
	}

	return grapper.New(options.Name, m...), nil
}
