package retry

import (
	"context"
	"net/http"
	"time"

	"github.com/americanas-go/log"
	"github.com/go-resty/resty/v2"
)

type Retry struct {
	options *Options
}

func NewRetryWithOptions(options *Options) *Retry {
	return &Retry{options: options}
}

func NewRetryWithConfigPath(path string) (*Retry, error) {
	o, err := NewOptionsWithPath(path)
	if err != nil {
		return nil, err
	}
	return NewRetryWithOptions(o), nil
}

func Register(ctx context.Context, client *resty.Client) error {
	o, err := NewOptions()
	if err != nil {
		return err
	}
	plugin := NewRetryWithOptions(o)
	return plugin.Register(ctx, client)
}

func (p *Retry) Register(ctx context.Context, client *resty.Client) error {

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
