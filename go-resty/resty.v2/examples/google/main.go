package main

import (
	"context"

	"github.com/americanas-go/config"
	ilog "github.com/americanas-go/ignite/americanas-go/log.v1"
	"github.com/americanas-go/ignite/go-resty/resty.v2"
	"github.com/americanas-go/log"
)

func main() {

	var err error

	config.Load()

	ctx := context.Background()

	ilog.New()

	logger := log.FromContext(ctx)

	options, err := resty.NewOptions()
	if err != nil {
		logger.Fatal(err.Error())
	}
	healthOptions := &(options.Plugins.Health)
	healthOptions.Name = "Google Inc"
	healthOptions.Host = "http://google.com"
	healthOptions.Endpoint = "/status"
	healthOptions.Enabled = true
	healthOptions.Description = "Search Engine"
	healthOptions.Required = true

	client, err := resty.NewWithOptions(ctx, options)
	if err != nil {
		logger.Fatal(err.Error())
	}
	request := client.R().EnableTrace()

	resp, err := request.Get("http://google.com")
	if err != nil {
		logger.Fatalf(err.Error())
	}

	if resp != nil {
		logger.Infof(resp.String())
	}
}
