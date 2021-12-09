package newrelic

import (
	"net/http"

	"github.com/americanas-go/config"
)

const (
	root                            = "ignite.newrelic"
	appName                         = root + ".appName"
	license                         = root + ".license"
	enabled                         = root + ".enabled"
	debug                           = root + ".debug"
	tracerEnabled                   = root + ".tracerEnabled"
	labels                          = root + ".labels"
	serverlessModeEnabled           = root + ".serverless.enabled"
	serverlessModeAccountID         = root + ".serverless.accountid"
	serverlessModeTrustedAccountKey = root + ".serverless.trustedAccountKey"
	serverlessModePrimaryAppID      = root + ".serverless.primaryAppId"
	serverlessModeApdexThreshold    = root + ".serverless.apdexThreshold"
	errorCollectorIgnoreStatusCodes = root + ".errorCollector.ignoreStatusCodes"
)

func init() {

	config.Add(appName, "", "application name for newrelic")
	config.Add(license, "", "newrelic license key")
	config.Add(enabled, true, "enables newrelic")
	config.Add(debug, false, "enables debug")
	config.Add(tracerEnabled, false, "enables newrelic distributed tracer")
	config.Add(labels, map[string]string{}, "newrelic labels")
	config.Add(serverlessModeEnabled, false, "enables newrelic serverless mode")
	config.Add(serverlessModeAccountID, "", "newrelic serverless mode account id")
	config.Add(serverlessModeTrustedAccountKey, "", "newrelic serverless mode trusted account key")
	config.Add(serverlessModePrimaryAppID, "", "newrelic serverless mode primary app id")
	config.Add(serverlessModeApdexThreshold, "", "newrelic serverless mode apdex threshold")
	config.Add(errorCollectorIgnoreStatusCodes,
		[]int{
			http.StatusBadRequest,
			http.StatusUnauthorized,
			http.StatusPaymentRequired,
			http.StatusForbidden,
			http.StatusNotFound,
			http.StatusMethodNotAllowed,
			http.StatusNotAcceptable,
			http.StatusProxyAuthRequired,
			http.StatusRequestTimeout,
			http.StatusConflict,
			http.StatusGone,
			http.StatusLengthRequired,
			http.StatusPreconditionFailed,
			http.StatusRequestEntityTooLarge,
			http.StatusRequestURITooLong,
			http.StatusUnsupportedMediaType,
			http.StatusRequestedRangeNotSatisfiable,
			http.StatusExpectationFailed,
			http.StatusTeapot,
			http.StatusMisdirectedRequest,
			http.StatusUnprocessableEntity,
			http.StatusLocked,
			http.StatusFailedDependency,
			http.StatusTooEarly,
			http.StatusUpgradeRequired,
			http.StatusPreconditionRequired,
			http.StatusTooManyRequests,
			http.StatusRequestHeaderFieldsTooLarge,
			http.StatusUnavailableForLegalReasons,
		},
		"newrelic serverless mode apdex threshold")
}

// IsEnabled returns enabled config value.
func IsEnabled() bool {
	return config.Bool(enabled)
}

// Debug returns debug config value.
func Debug() bool {
	return config.Bool(debug)
}
