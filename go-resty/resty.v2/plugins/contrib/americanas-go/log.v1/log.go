package log

import (
	"context"
	"encoding/json"
	"time"

	"github.com/americanas-go/log"
	"github.com/go-resty/resty/v2"
)

type l func(format string, args ...interface{})

var lvl string

type Log struct {
	options *Options
}

func NewLogWithConfigPath(path string) (*Log, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewLogWithOptions(o), nil
}

func NewLogWithOptions(options *Options) *Log {
	return &Log{options: options}
}

func Register(ctx context.Context, client *resty.Client) error {
	o, err := NewOptions()
	if err != nil {
		return err
	}

	plugin := NewLogWithOptions(o)
	return plugin.Register(ctx, client)
}

func (i *Log) Register(ctx context.Context, client *resty.Client) error {

	if !i.options.Enabled {
		return nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling logger middleware in resty")

	client.OnBeforeRequest(i.logBeforeResponse)
	client.OnAfterResponse(i.logAfterResponse)

	logger.Debug("logger middleware successfully enabled in resty")

	return nil
}

func (i *Log) m(logger log.Logger) (method l) {

	if lvl == "" {
		lvl = i.options.Level
	}
	switch lvl {
	case "TRACE":
		method = logger.Tracef
	case "INFO":
		method = logger.Infof
	default:
		method = logger.Debugf
	}

	return method
}

func (i *Log) logBeforeResponse(client *resty.Client, request *resty.Request) error {

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

	xx := i.m(logger)

	xx("rest request processing")

	return nil
}

func (i *Log) logAfterResponse(client *resty.Client, response *resty.Response) error {

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
		xx := i.m(logger)
		xx("successful rest request processed")
	}

	return nil
}
