package contrib

import (
	"context"
	"strconv"
	"strings"

	"github.com/americanas-go/errors"
	"github.com/americanas-go/health"
	iresty "github.com/americanas-go/ignite/go-resty/resty.v2/resty"
	"github.com/americanas-go/log"
	"github.com/go-resty/resty/v2"
)

// Adds health check.
func Health(ctx context.Context, w *iresty.Wrapper) error {
	o := w.Options.Plugins.Health
	if !o.Enabled {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("integrating resty in health")

	hc := health.NewHealthChecker(
		o.Name,
		o.Description,
		&checker{
			client: w.Instance,
			url:    strings.Join([]string{o.Host, o.Endpoint}, ""),
		},
		o.Required,
		o.Enabled,
	)
	health.Add(hc)

	logger.Debug("resty successfully integrated in health")

	return nil
}

type checker struct {
	client *resty.Client
	url    string
}

func (c *checker) Check(ctx context.Context) (err error) {

	request := c.client.R()

	var response *resty.Response

	response, err = request.Get(c.url)

	if err != nil {
		return errors.Internalf(err.Error())
	}

	if response.IsError() {
		return errors.New(strconv.Itoa(response.StatusCode()))
	}

	return err
}
