package main

import (
	"github.com/labstack/echo"
	"net/http"
	"net/url"
)

func malicious(c echo.Context) error {
	values := url.Values{}
	values.Add("username", "hacking")
	http.PostForm("http://localhost:1323/login", values)
	return c.String(http.StatusOK, "hack")
}
func main() {
	e := echo.New()
	e.Any("/", malicious)
	e.Logger.Fatal(e.Start(":8080"))
}
