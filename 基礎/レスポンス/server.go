package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func response(c echo.Context) error {
	return c.HTML(http.StatusOK, "<h1>Hello, ema!</h1>")
}

func main() {
	e := echo.New()
	e.GET("/", response)
	e.Start(":1323")
}