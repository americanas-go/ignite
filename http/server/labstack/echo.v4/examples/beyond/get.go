package main

import (
	"net/http"

	"github.com/americanas-go/config"
	e "github.com/labstack/echo/v4"
)

func Get(c e.Context) (err error) {

	resp := Response{
		Message: "Hello Google!!",
	}

	err = config.Unmarshal(&resp)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}
