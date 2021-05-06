package main

import (
	"time"

	"github.com/americanas-go/config"
)

const (

	// config google client

	GoogleHost              = "app.client.resty.google.host"
	GoogleDebug             = "app.client.resty.google.debug"
	GoogleRequestTimeout    = "app.client.resty.google.request.timeout"
	GoogleRetryCount        = "app.client.resty.google.retry.count"
	GoogleRetryWaitTime     = "app.client.resty.google.retry.waittime"
	GoogleRetryMaxWaitTime  = "app.client.resty.google.retry.maxwaittime"
	GoogleHealthEnabled     = "app.client.resty.google.health.enabled"
	GoogleHealthDescription = "app.client.resty.google.health.description"
	GoogleHealthEndpoint    = "app.client.resty.google.health.endpoint"
	GoogleHealthRequired    = "app.client.resty.google.health.required"

	// config americanas client

	ACOMHost              = "app.client.resty.acom.host"
	ACOMDebug             = "app.client.resty.acom.debug"
	ACOMRequestTimeout    = "app.client.resty.acom.request.timeout"
	ACOMRetryCount        = "app.client.resty.acom.retry.count"
	ACOMRetryWaitTime     = "app.client.resty.acom.retry.waittime"
	ACOMRetryMaxWaitTime  = "app.client.resty.acom.retry.maxwaittime"
	ACOMHealthEnabled     = "app.client.resty.acom.health.enabled"
	ACOMHealthDescription = "app.client.resty.acom.health.description"
	ACOMHealthEndpoint    = "app.client.resty.acom.health.endpoint"
	ACOMHealthRequired    = "app.client.resty.acom.health.required"
)

func init() {

	config.Add(GoogleHost, "http://www.google.com", "defines host")
	config.Add(GoogleDebug, false, "defines client debug request")
	config.Add(GoogleRequestTimeout, 2*time.Second, "defines client http request timeout (ms)")
	config.Add(GoogleRetryCount, 0, "defines client max http retries")
	config.Add(GoogleRetryWaitTime, 200*time.Millisecond, "defines client retry wait time (ms)")
	config.Add(GoogleRetryMaxWaitTime, 2*time.Second, "defines client max retry wait time (ms)")
	config.Add(GoogleHealthEnabled, true, "enable/disable health")
	config.Add(GoogleHealthDescription, "google endpoint", "defines health description")
	config.Add(GoogleHealthEndpoint, "http://www.google.com", "defines health endpoint")
	config.Add(GoogleHealthRequired, true, "enable/disable health required dependency")

	config.Add(ACOMHost, "http://www.americanas.com", "defines host")
	config.Add(ACOMDebug, false, "defines client debug request")
	config.Add(ACOMRequestTimeout, 2*time.Second, "defines client http request timeout (ms)")
	config.Add(ACOMRetryCount, 0, "defines client max http retries")
	config.Add(ACOMRetryWaitTime, 200*time.Millisecond, "defines client retry wait time (ms)")
	config.Add(ACOMRetryMaxWaitTime, 2*time.Second, "defines client max retry wait time (ms)")
	config.Add(ACOMHealthEnabled, true, "enable/disable health")
	config.Add(ACOMHealthDescription, "google endpoint", "defines health description")
	config.Add(ACOMHealthEndpoint, "http://www.google.com", "defines health endpoint")
	config.Add(ACOMHealthRequired, true, "enable/disable health required dependency")

}
