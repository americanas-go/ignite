package health

import (
	"context"
	"strconv"
	"strings"

	"github.com/americanas-go/errors"
	"github.com/go-resty/resty/v2"
)

type Checker struct {
	client  *resty.Client
	options *Options
}

func (c *Checker) Check(ctx context.Context) (err error) {

	request := c.client.R()

	var response *resty.Response

	response, err = request.Get(strings.Join([]string{c.options.Host, c.options.Endpoint}, ""))
	if err != nil {
		return errors.Internalf(err.Error())
	}

	if response.IsError() {
		return errors.New(strconv.Itoa(response.StatusCode()))
	}

	return err
}

func NewChecker(client *resty.Client, options *Options) *Checker {
	return &Checker{client: client, options: options}
}
