package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "ok\n")
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Static("/new", "./static")
	e.File("file/test", "./static/test.html")
	e.GET("/hello", hello)
	e.Logger.Fatal(e.Start(":1323"))
}
