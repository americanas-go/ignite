package datadog

import (
	"context"
	"sync"

	"github.com/americanas-go/ignite/net/http/client"
	"github.com/americanas-go/log"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
)

// StartProfiler starts the profiler.
func StartProfiler(ctx context.Context, profileOptions ...profiler.Option) {

	o, err := NewOptions()
	if err != nil {
		panic(err)
	}

	StartProfilerWithOptions(ctx, o, profileOptions...)
}

var profilerOnce sync.Once

// StartProfilerWithOptions start the profiler with options.
func StartProfilerWithOptions(ctx context.Context, options *Options, profileOptions ...profiler.Option) {

	if !IsProfilerEnabled() {
		return
	}

	profilerOnce.Do(func() {

		logger := log.FromContext(ctx)

		httpClient := client.NewClientWithOptions(ctx, &options.HttpClient)

		var tags []string

		for _, v := range options.Tags {
			tags = append(tags, v)
		}

		o := []profiler.Option{
			profiler.WithAgentAddr(options.Addr),
			profiler.WithEnv(options.Env),
			profiler.WithService(options.Service),
			profiler.WithVersion(options.Version),
			profiler.WithHTTPClient(httpClient),
			profiler.WithTags(tags...),
			profiler.WithProfileTypes(
				profiler.CPUProfile,
				profiler.HeapProfile,
				profiler.GoroutineProfile,
				profiler.MetricsProfile,
				// The profiles below are disabled by default to keep overhead
				// low, but can be enabled as needed.

				// profiler.BlockProfile,
				// profiler.MutexProfile,
			),
		}

		o = append(o, profileOptions...)

		if err := profiler.Start(o...); err != nil {
			logger.Errorf("datadog profiler not started. %s", err.Error())
		} else {
			logger.Infof("started a datadog profiler: %s", options.Service)
		}

	})

}
