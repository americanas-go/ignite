package grapper

import (
	"context"
	"github.com/americanas-go/grapper"
)

func NewAnyErrorWrapper[R any](ctx context.Context, plugins ...func(ctx context.Context, name string) grapper.AnyErrorMiddleware[R]) (*grapper.AnyErrorWrapper[R], error) {
	o, err := NewOptions()
	if err != nil {
		return nil, err
	}

	return NewAnyErrorWrapperWithOptions(ctx, o, plugins...)
}

func NewAnyErrorWrapperWithPath[R any](ctx context.Context, path string, plugins ...func(ctx context.Context, name string) grapper.AnyErrorMiddleware[R]) (*grapper.AnyErrorWrapper[R], error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}

	return NewAnyErrorWrapperWithOptions(ctx, o, plugins...)
}

func NewAnyErrorWrapperWithOptions[R any](ctx context.Context, options *Options, plugins ...func(ctx context.Context, name string) grapper.AnyErrorMiddleware[R]) (*grapper.AnyErrorWrapper[R], error) {
	var m []grapper.AnyErrorMiddleware[R]

	for _, f := range plugins {
		if p := f(ctx, options.Name); p != nil {
			m = append(m, p)
		}
	}

	return grapper.NewAnyErrorWrapper(ctx, options.Name, m...), nil
}

func NewAnyWrapper[R any](ctx context.Context, plugins ...func(ctx context.Context, name string) grapper.AnyMiddleware[R]) (*grapper.AnyWrapper[R], error) {
	o, err := NewOptions()
	if err != nil {
		return nil, err
	}

	return NewAnyWrapperWithOptions(ctx, o, plugins...)
}

func NewAnyWrapperWithPath[R any](ctx context.Context, path string, plugins ...func(ctx context.Context, name string) grapper.AnyMiddleware[R]) (*grapper.AnyWrapper[R], error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}

	return NewAnyWrapperWithOptions(ctx, o, plugins...)
}

func NewAnyWrapperWithOptions[R any](ctx context.Context, options *Options, plugins ...func(ctx context.Context, name string) grapper.AnyMiddleware[R]) (*grapper.AnyWrapper[R], error) {
	var m []grapper.AnyMiddleware[R]

	for _, f := range plugins {
		if p := f(ctx, options.Name); p != nil {
			m = append(m, p)
		}
	}

	return grapper.NewAnyWrapper(ctx, options.Name, m...), nil
}

func NewErrorWrapper(ctx context.Context, plugins ...func(ctx context.Context, name string) grapper.ErrorMiddleware) (*grapper.ErrorWrapper, error) {
	o, err := NewOptions()
	if err != nil {
		return nil, err
	}

	return NewErrorWrapperWithOptions(ctx, o, plugins...)
}

func NewErrorWrapperWithPath(ctx context.Context, path string, plugins ...func(ctx context.Context, name string) grapper.ErrorMiddleware) (*grapper.ErrorWrapper, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}

	return NewErrorWrapperWithOptions(ctx, o, plugins...)
}

func NewErrorWrapperWithOptions(ctx context.Context, options *Options, plugins ...func(ctx context.Context, name string) grapper.ErrorMiddleware) (*grapper.ErrorWrapper, error) {
	var m []grapper.ErrorMiddleware

	for _, f := range plugins {
		if p := f(ctx, options.Name); p != nil {
			m = append(m, p)
		}
	}

	return grapper.NewErrorWrapper(ctx, options.Name, m...), nil
}
