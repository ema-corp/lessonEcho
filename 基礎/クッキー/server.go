package main

import (
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func createCookie(c echo.Context) error {
	// クッキー情報を格納するオブジェクトを生成
	cookie := new(http.Cookie)

	// クッキーの名前を設定
	cookie.Name = "username"

	// クッキーの値を設定
	cookie.Value = "ema"

	// クッキーの有効期限を設定
	cookie.Expires = time.Now().Add(24 * time.Hour)

	// クッキーをセット
	c.SetCookie(cookie)

	return c.String(http.StatusOK, "Create a cookie.")
}

// ハンドラ内でCookieを読む
func readCookie(c echo.Context) error {
	// クッキーの名前を指定して、そのクッキー情報を読み込む
	cookie, _ := c.Cookie("username")

	// クッキー情報から読み取ったユーザ名を送信
	return c.String(http.StatusOK, "Your name is "+cookie.Value)
}

func main() {
	e := echo.New()
	e.GET("/", createCookie)
	e.GET("/read", readCookie)
	e.Start(":1323")
}