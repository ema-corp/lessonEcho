package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"net/http"
)

func login(c echo.Context) error {
	name := c.QueryParam("name") // クエリパラメータからユーザ名を取得

	// 名前が入力されていなければログイン失敗
	if name == "" {
		return c.String(http.StatusOK, "ログインに失敗しました。クエリパラメータでユーザ名を設定してください。")
	}

	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		MaxAge:   1000, // 有効時間を1000秒に設定
	}
	sess.Values["login"] = true // ログイン状態を有効化
	sess.Values["name"] = name // ユーザ名をセッションデータに格納
	sess.Save(c.Request(), c.Response())
	return c.String(http.StatusOK, "ログインが完了しました。")
}

func logout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Values["login"] = false // ログイン状態を無効化
	sess.Save(c.Request(), c.Response())
	return c.String(http.StatusOK, "ログアウトが完了しました。")
}

func myPage(c echo.Context) error {
	sess, _ := session.Get("session", c)

	// ログインされていない場合は認証失敗を伝える
	if status, _ := sess.Values["login"]; status != true {
		return c.String(http.StatusUnauthorized, "unauthorized")
	}

	// ログインされている場合は名前を表示
	message := fmt.Sprintf("%v としてログイン中", sess.Values["name"])
	return c.String(http.StatusOK, message)
}

func main() {
	e := echo.New()

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	// 以下省略

	e.GET("/login", login) // ログイン用ページ
	e.GET("/logout", logout) // ログアウト用ページ
	e.GET("/mypage", myPage) // マイページ

	e.Start(":1323")
}