package main

import (
	"context"
	"os"

	"github.com/americanas-go/config"
	c "github.com/americanas-go/ignite/cmd/spf13/cobra.v1"
	"github.com/americanas-go/ignite/http/client/go-resty/resty.v2"
	logplugin "github.com/americanas-go/ignite/http/client/go-resty/resty.v2/plugins/contrib/americanas-go/log.v1"
	ilog "github.com/americanas-go/ignite/log/americanas-go/log.v1"
	"github.com/americanas-go/log"
	r "github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

const (

	// config google client
	google = "app.resty.google"

	// config americanas client
	acom          = "app.resty.acom"
	acomPlugins   = acom + ".plugins"
	acomLogPlugin = acomPlugins + ".log"
)

func init() {

	os.Setenv("IGNITE_LOGRUS_CONSOLE_LEVEL", "INFO")

	os.Setenv("APP_RESTY_GOOGLE_HOST", "http://www.google.com")
	os.Setenv("APP_RESTY_ACOM_HOST", "https://www.americanas.com.br")
	os.Setenv("APP_RESTY_ACOM_PLUGINS_LOG_LEVEL", "INFO")

	resty.ConfigAdd(google)
	resty.ConfigAdd(acom)

	logplugin.ConfigAdd(acomLogPlugin)
}

func main() {

	config.Load()
	ilog.New()

	ctx := context.Background()
	logger := log.FromContext(ctx)

	cmds := []*cobra.Command{
		{
			Use:  "google",
			Long: "google call",
			RunE: func(cmd *cobra.Command, args []string) error {
				return call(ctx, google)
			},
		},
		{
			Use:  "acom",
			Long: "acom call",
			RunE: func(cmd *cobra.Command, args []string) error {

				acomLogP, err := logplugin.NewLogWithConfigPath(acomLogPlugin)
				if err != nil {
					return err
				}

				return call(ctx, acom, acomLogP.Register)
			},
		},
	}

	rootCMD := &cobra.Command{
		Version: "1.0.0",
	}

	if err := c.Run(rootCMD, cmds...); err != nil {
		logger.Errorf(err.Error())
	}

	// go run main.go -> show options
	// go run main.go acom -> call acom
	// go run main.go google -> call google
}

func call(ctx context.Context, path string, plugins ...resty.Plugin) error {

	logger := log.FromContext(ctx)

	var err error

	var client *r.Client
	if client, err = resty.NewClientWithConfigPath(ctx, path, plugins...); err != nil {
		return err
	}

	var response *r.Response
	if response, err = client.R().Get("/"); err != nil {
		return err
	}

	if response != nil {
		logger.Infof(response.String())
	}

	return nil
}
