package newrelic

import (
	"context"
	"time"

	"github.com/americanas-go/config"
	"github.com/americanas-go/log"
	"github.com/newrelic/go-agent/v3/newrelic"
)

var app *newrelic.Application

func Application() *newrelic.Application {
	if app == nil {
		var err error
		if app, err = NewApplication(context.Background()); err != nil {
			panic(err)
		}
	}
	return app
}

func NewApplication(ctx context.Context) (*newrelic.Application, error) {

	if !IsEnabled() {
		return nil, nil
	}

	logger := log.FromContext(ctx)

	enabled := config.Bool(enabled)
	appName := config.String(appName)
	a, err := newrelic.NewApplication(
		newrelic.ConfigAppName(appName),
		newrelic.ConfigLicense(config.String(license)),
		newrelic.ConfigEnabled(enabled),
		newrelic.ConfigDistributedTracerEnabled(config.Bool(tracerEnabled)),
		newrelic.ConfigLogger(NewLogger()),
		// newrelic.ConfigDebugLogger(log.GetLogger().Output()),
		func(cfg *newrelic.Config) {
			cfg.ErrorCollector.IgnoreStatusCodes = config.Ints(errorCollectorIgnoreStatusCodes)
			cfg.Labels = config.StringMap(labels)
			cfg.ServerlessMode.Enabled = config.Bool(serverlessModeEnabled)
			cfg.ServerlessMode.AccountID = config.String(serverlessModeAccountID)
			cfg.ServerlessMode.TrustedAccountKey = config.String(serverlessModeTrustedAccountKey)
			cfg.ServerlessMode.PrimaryAppID = config.String(serverlessModePrimaryAppID)
			if apdex, err := time.ParseDuration(config.String(serverlessModeApdexThreshold) + "s"); nil == err {
				cfg.ServerlessMode.ApdexThreshold = apdex
			}
		},
	)

	if err != nil {
		return nil, err
	}

	logger.Debugf("started a new NewRelic application: %s", appName)

	app = a

	return app, nil
}
