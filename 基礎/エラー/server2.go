package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func errorHandling2(c echo.Context) error {
	return c.String(http.StatusOK, "Success!")
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	c.JSON(code, code)
}

func main() {
	e := echo.New()
	e.HTTPErrorHandler = customHTTPErrorHandler
	e.GET("/", errorHandling2)
	e.Start(":1323")
}