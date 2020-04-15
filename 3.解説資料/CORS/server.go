package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func cors(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func main() {
	e := echo.New()

	// ミドルウェアを利用してCORSを設定
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://ema.com", "https://echo.com"},  // 許可するオリジンを設定。すべてのオリジンを許可する場合は「*」を設定
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept}, // 許可するヘッダを設定。
	}))

	e.GET("/", cors)

	e.Start(":1323")
}
