package main

import (
	"context"
	"os"

	"github.com/americanas-go/config"
	ilog "github.com/americanas-go/ignite/americanas-go/log.v1"
	"github.com/americanas-go/ignite/lib/resty.v2"
	c "github.com/americanas-go/ignite/spf13/cobra.v1"
	"github.com/americanas-go/log"
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
	os.Setenv("APP_RESTY_ACOM_PLUGINS_LOG_ENABLED", "true")
	os.Setenv("APP_RESTY_ACOM_PLUGINS_LOG_LEVEL", "INFO")
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
				return call(ctx, acom)
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

func call(ctx context.Context, path string) error {

	logger := log.FromContext(ctx)

	var err error

	client, err := resty.NewWithConfigPath(ctx, path)
	if err != nil {
		return err
	}

	response, err := client.R().Get("/")
	if err != nil {
		return err
	}

	if response != nil {
		logger.Infof(response.String())
	}

	return nil
}
