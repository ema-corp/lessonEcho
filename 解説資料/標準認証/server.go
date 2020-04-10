package main

import (
	"crypto/subtle"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func basicAuth(c echo.Context) error {
	return c.String(http.StatusOK, "Authentication success.")
}

func main() {
	e := echo.New()

	// 標準認証を利用する
	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// 文字列比較の処理時間で認証情報を推測（タイミング攻撃）されないように、ConstantTimeCompare関数を利用
		if subtle.ConstantTimeCompare([]byte(username), []byte("ema")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("echo")) == 1 {
			return true, nil
		}
		return false, nil
	}))

	e.GET("/", basicAuth)
	e.Start(":1323")
}