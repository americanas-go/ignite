package main

import (
	"context"
	"os"

	"github.com/americanas-go/config"
	ilog "github.com/americanas-go/ignite/americanas-go/log.v1"
	"github.com/americanas-go/ignite/go-resty/resty.v2"
	"github.com/americanas-go/log"
)

const (

	// config google client
	googleConfigPath          = "app.resty.google"
	googlePluginsConfigPath   = googleConfigPath + ".plugins"
	googleLogPluginConfigPath = googlePluginsConfigPath + ".log"

	// config americanas client
	acomConfigPath          = "app.resty.acom"
	acomPluginsConfigPath   = acomConfigPath + ".plugins"
	acomLogPluginConfigPath = acomPluginsConfigPath + ".log"
)

func init() {

	os.Setenv("IGNITE_LOGRUS_CONSOLE_LEVEL", "INFO")

	os.Setenv("APP_RESTY_GOOGLE_HOST", "http://www.google.com")
	os.Setenv("APP_RESTY_GOOGLE_PLUGINS_LOG_ENABLED", "true")
	os.Setenv("APP_RESTY_GOOGLE_PLUGINS_LOG_LEVEL", "INFO")
	os.Setenv("APP_RESTY_ACOM_HOST", "https://www.americanas.com.br")
	os.Setenv("APP_RESTY_ACOM_PLUGINS_LOG_ENABLED", "true")
	os.Setenv("APP_RESTY_ACOM_PLUGINS_LOG_LEVEL", "INFO")
}

func main() {

	config.Load()
	ilog.New()

	ctx := context.Background()
	logger := log.FromContext(ctx)

	var err error

	// ACOM CALL
	clientAcom, err := resty.NewWithConfigPath(ctx, acomConfigPath)
	if err != nil {
		log.Fatal(err)
	}

	responseAcom, err := clientAcom.R().Get("/")
	if err != nil {
		log.Fatal(err)
	}

	if responseAcom != nil {
		logger.Infof(responseAcom.String())
	}

	// GOOGLE CALL
	clientGoogle, err := resty.NewWithConfigPath(ctx, googleConfigPath)
	if err != nil {
		log.Fatal(err)
	}

	responseGoogle, err := clientGoogle.R().Get("/")
	if err != nil {
		log.Fatal(err)
	}

	if responseGoogle != nil {
		logger.Infof(responseGoogle.String())
	}

}
