package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func hello(c echo.Context) error {
	return c.Render(http.StatusOK)
}

func main() {

	e := echo.New()
	e.GET("/", hello)
	e.Logger.Fatal(e.Start(":1323"))
}
