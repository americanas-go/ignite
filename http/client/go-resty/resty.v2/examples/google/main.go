package main

import (
	"context"

	"github.com/americanas-go/config"
	"github.com/americanas-go/ignite/http/client/go-resty/resty.v2"
	"github.com/americanas-go/ignite/http/client/go-resty/resty.v2/plugins/contrib/americanas-go/health.v1"
	ilog "github.com/americanas-go/ignite/log/americanas-go/log.v1"
	"github.com/americanas-go/log"
	r "github.com/go-resty/resty/v2"
)

func main() {

	var err error

	config.Load()

	ctx := context.Background()

	ilog.New()

	logger := log.FromContext(ctx)

	options := health.Options{
		Name:        "Google Inc",
		Host:        "http://google.com",
		Endpoint:    "/status",
		Enabled:     true,
		Description: "Search Engine",
		Required:    true,
	}

	healthIntegrator := health.NewHealthWithOptions(&options)

	client := resty.NewClientWithOptions(ctx, &resty.Options{}, healthIntegrator.Register)
	request := client.R().EnableTrace()

	var resp *r.Response
	resp, err = request.Get("http://google.com")
	if err != nil {
		logger.Fatalf(err.Error())
	}

	if resp != nil {
		logger.Infof(resp.String())
	}
}
