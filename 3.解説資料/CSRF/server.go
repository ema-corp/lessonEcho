package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func csrf(c echo.Context) error {
	return c.String(http.StatusOK, "Hello.")
}

func main() {
	e := echo.New()

	// CSRFを利用する
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:X-XSRF-TOKEN",
	}))

	e.GET("/", csrf)
	e.Start(":1323")
}