package main

import (
	"net/http"

	"github.com/americanas-go/config"
	"github.com/americanas-go/log"
	"github.com/jvitoroc/ignite/labstack/echo.v4"
	e "github.com/labstack/echo/v4"
)

func Get(c e.Context) (err error) {

	logger := log.FromContext(c.Request().Context())

	resp := Response{
		Message: "Hello Google!!",
	}

	err = config.Unmarshal(&resp)
	if err != nil {
		logger.Errorf(err.Error())
	}

	return echo.JSON(c, http.StatusOK, resp, err)
}
