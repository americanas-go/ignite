package ants

import (
	"context"
	"sync"

	"github.com/americanas-go/log"
	"github.com/panjf2000/ants/v2"
)

// Middleware interface for task middlewares
type Middleware interface {

	// Before method you perform before the task
	Before(ctx context.Context) context.Context

	// After method you perform after the task
	After(ctx context.Context)
}

// Task interface for tasks
type Task func(ctx context.Context) context.Context

// Wrapper
type Wrapper struct {
	// pool instance of ants pool
	pool *ants.Pool

	// middlewares list of middleware included in the wrapper
	middlewares []Middleware
}

// NewWrapper generates an instance of giants wrapper.
func NewWrapper(pool *ants.Pool, middlewares ...Middleware) *Wrapper {
	log.Trace("creating ants wrapper")
	return &Wrapper{pool: pool, middlewares: middlewares}
}

// Submit sends a task to ant, needs a sync.WaitGroup
func (a *Wrapper) Submit(ctx context.Context, task Task, wg *sync.WaitGroup) error {

	logger := log.FromContext(ctx).WithTypeOf(*a)
	logger.Trace("submit ants task")

	wg.Add(1)

	err := ants.Submit(func() {

		a.exec(ctx, task)

		wg.Done()

		logger.Debug("ants task executed")
	})

	return err
}

// AsyncSubmit sends a task to ant
func (a *Wrapper) AsyncSubmit(ctx context.Context, task Task) error {

	logger := log.FromContext(ctx).WithTypeOf(*a)
	logger.Trace("submit async ants task")

	err := ants.Submit(func() {

		a.exec(ctx, task)

		logger.Debug("async ants task executed")
	})

	return err
}

// exec performs the task and the middleware
func (a *Wrapper) exec(ctx context.Context, task Task) {

	for _, m := range a.middlewares {
		ctx = m.Before(ctx)
	}

	ctx = task(ctx)

	for _, m := range a.middlewares {
		m.After(ctx)
	}
}
