package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var responseMsg = "Hello World"

func NewRouter() *echo.Echo {
	handler := func(c echo.Context) error {
		return c.String(http.StatusOK, responseMsg)
	}
	e := echo.New()
	e.GET("/", handler)
	return e
}
