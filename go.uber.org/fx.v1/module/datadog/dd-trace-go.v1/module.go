package datadog

import (
	"sync"

	datadog "github.com/americanas-go/ignite/datadog/dd-trace-go.v1"
	contextfx "github.com/americanas-go/ignite/go.uber.org/fx.v1/module/context"
	"go.uber.org/fx"
)

var optOnce sync.Once

// Module fx module for datadog options.
func OptionsModule() fx.Option {
	options := fx.Options()

	optOnce.Do(func() {
		options = fx.Options(
			fx.Provide(
				datadog.NewOptions,
			),
		)
	})

	return options
}

var tracerOnce sync.Once

// Module fx module for datadog tracer.
func TracerModule() fx.Option {
	options := fx.Options()

	tracerOnce.Do(func() {
		options = fx.Options(
			contextfx.Module(),
			OptionsModule(),
			fx.Invoke(
				datadog.StartTracerWithOptions,
			),
		)
	})

	return options
}

var profilerOnce sync.Once

// Module fx module for datadog profiler.
func ProfilerModule() fx.Option {
	options := fx.Options()

	profilerOnce.Do(func() {
		options = fx.Options(
			contextfx.Module(),
			OptionsModule(),
			fx.Invoke(
				datadog.StartProfilerWithOptions,
			),
		)
	})

	return options
}
