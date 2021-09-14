package retry

import (
	"context"
	"net/http"
	"time"

	"github.com/americanas-go/log"
	r "github.com/go-resty/resty/v2"
)

type Retry struct {
	options *Options
}

// NewRetryWithOptions returns a new Retry with options.
func NewRetryWithOptions(options *Options) *Retry {
	return &Retry{options: options}
}

// NewRetry returns a new Retry.
func NewRetry() *Retry {
	o, err := NewOptions()
	if err != nil {
		log.Fatalf(err.Error())
	}
	return NewRetryWithOptions(o)
}

// Registry registers retry in resty.
func (p *Retry) Register(ctx context.Context, client *r.Client) error {

	if !p.options.Enabled {
		return nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("configuring retry in resty")

	client.
		SetRetryCount(p.options.Count).
		SetRetryWaitTime(p.options.WaitTime).
		SetRetryMaxWaitTime(p.options.MaxWaitTime).
		AddRetryCondition(statusCodeRetryCondition).
		AddRetryCondition(addTimeoutRetryCondition(client.GetClient().Timeout))

	logger.Debug("retry successfully configured in resty")

	return nil
}

func statusCodeRetryCondition(r *r.Response, err error) bool {
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

func addTimeoutRetryCondition(timeout time.Duration) func(r *r.Response, err error) bool {

	return func(resp *r.Response, err error) bool {

		if resp.Time() > timeout {
			return true
		}

		return false
	}
}
