package contrib

import (
	"context"
	"encoding/json"
	"time"

	iresty "github.com/americanas-go/ignite/lib/resty.v2/resty"
	"github.com/americanas-go/log"
	"github.com/go-resty/resty/v2"
)

// Adds logging for made requests and received responses.
func Log(ctx context.Context, w *iresty.Wrapper) error {
	o := w.Options.Plugins.Log
	if !o.Enabled {
		return nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling logger middleware in resty")
	lvl = o.Level
	w.Instance.OnBeforeRequest(logBeforeResponse)
	w.Instance.OnAfterResponse(logAfterResponse)

	logger.Debug("logger middleware successfully enabled in resty")

	return nil
}

var lvl string

func m(logger log.Logger) func(format string, args ...interface{}) {
	switch lvl {
	case "TRACE":
		return logger.Tracef
	case "INFO":
		return logger.Infof
	default: // DEBUG
		return logger.Debugf
	}
}

func logBeforeResponse(client *resty.Client, request *resty.Request) error {

	logger := log.FromContext(request.Context())

	requestHeaders, _ := json.Marshal(request.Header)

	requestBody, _ := json.Marshal(request.Body)

	logger = logger.
		WithFields(
			log.Fields{
				"rest_client_host":     client.HostURL,
				"rest_request_body":    string(requestBody),
				"rest_request_url":     request.URL,
				"rest_request_headers": string(requestHeaders),
				"rest_request_method":  request.Method,
			})

	xx := m(logger)

	xx("rest request processing")

	return nil
}

func logAfterResponse(client *resty.Client, response *resty.Response) error {

	logger := log.FromContext(response.Request.Context())

	responseHeaders, _ := json.Marshal(response.Header())

	statusCode := response.StatusCode()

	logger = logger.WithFields(
		log.Fields{
			"rest_response_body":        string(response.Body()),
			"rest_response_headers":     string(responseHeaders),
			"rest_response_time":        response.Time().Seconds() * float64(time.Second/time.Millisecond),
			"rest_response_status_code": statusCode,
		})

	if statusCode > 500 {
		logger.Errorf("rest request processed with error")
	} else if statusCode > 400 {
		logger.Warnf("rest request processed with warning")
	} else {
		xx := m(logger)
		xx("successful rest request processed")
	}

	return nil
}
