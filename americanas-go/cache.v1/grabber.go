package grapper

import "github.com/americanas-go/grapper"

func New[R any](name string, plugins ...func(name string) grapper.Middleware[R]) *grapper.Wrapper[R] {
	var m []grapper.Middleware[R]
	for _, f := range plugins {
		if p := f(name); p != nil {
			m = append(m, p)
		}
	}
	return grapper.New(name, m...)
}
