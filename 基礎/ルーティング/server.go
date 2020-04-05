package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func returnUserID(c echo.Context) error {
	userID := c.Param("id")
	return c.String(http.StatusOK, userID)
}

func main() {
	e := echo.New()
	e.GET("/top", hello)
	e.Any("/user/:id", returnUserID)
	e.Logger.Fatal(e.Start(":1323"))
}
