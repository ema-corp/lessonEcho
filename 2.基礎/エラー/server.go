package main

import (
	"crypto/subtle"
	"github.com/labstack/echo"
	"net/http"
)

func errorHandling(c echo.Context) error {
	return c.String(http.StatusOK, "Success!")
}

func main() {
	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 標準認証パラメータを取得
			username, password, _ := c.Request().BasicAuth()

			// 認証を実行
			if subtle.ConstantTimeCompare([]byte(username), []byte("ema")) == 1 &&
				subtle.ConstantTimeCompare([]byte(password), []byte("echo")) == 1 {
				// 認証が成功した場合は通常通りハンドラを実行させる
				return next(c)
			}

			// 認証が失敗した場合はエラーを通知する
			return echo.NewHTTPError(http.StatusUnauthorized, "Please provide correct username or password.")
		}
	})
	e.GET("/", errorHandling)
	e.Start(":1323")
}