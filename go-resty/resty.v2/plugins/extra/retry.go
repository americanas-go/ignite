package extra

import (
	"context"
	"net/http"
	"time"

	iresty "github.com/americanas-go/ignite/go-resty/resty.v2/resty"
	"github.com/americanas-go/log"
	"github.com/go-resty/resty/v2"
)

// Retries the request when http response status code is:
//   429 Too Many Requests
//   500 Internal Server Error
//   504 Gateway timeout
//   503 Service Unavailable
func Retry(ctx context.Context, w *iresty.Wrapper) error {
	options := w.Options.Plugins.Retry
	if !options.Enabled {
		return nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("configuring retry in resty")

	w.Instance.
		SetRetryCount(options.Count).
		SetRetryWaitTime(options.WaitTime).
		SetRetryMaxWaitTime(options.MaxWaitTime).
		AddRetryCondition(statusCodeRetryCondition).
		AddRetryCondition(addTimeoutRetryCondition(w.Options.ConnectionTimeout))

	logger.Debug("retry successfully configured in resty")

	return nil
}

func statusCodeRetryCondition(r *resty.Response, err error) bool {
	switch statusCode := r.StatusCode(); statusCode {

	case http.StatusTooManyRequests:
		return true
	case http.StatusInternalServerError:
		return true
	case http.StatusGatewayTimeout:
		return true
	case http.StatusServiceUnavailable:
		return true
	default:
		return false
	}
}

func addTimeoutRetryCondition(timeout time.Duration) func(r *resty.Response, err error) bool {

	return func(resp *resty.Response, err error) bool {

		if resp.Time() > timeout {
			return true
		}

		return false
	}
}
