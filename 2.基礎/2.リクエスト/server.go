package main

import (
	"github.com/labstack/echo"
)

func request(c echo.Context) error {
	method := c.Request().Method
	println("HTTPメソッド：", method)
	return nil
}

func main() {
	e := echo.New()
	e.GET("/", request)
	e.Start(":1323")
}
