package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
	"html/template"
	"io"
	"net/http"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func top(c echo.Context) error {
	return c.Render(http.StatusOK, "top", "")
}

func login(c echo.Context) error {
	name := c.FormValue("username") // クエリパラメータからユーザ名を取得
	fmt.Println(name)
	// 名前が入力されていなければログイン失敗
	if name == "" {
		return c.Redirect(http.StatusPermanentRedirect, "/")
	}

	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		MaxAge: 1000,
		Path:   "/",
	}
	sess.Values["login"] = true // ログイン状態を有効化
	sess.Values["name"] = name  // ユーザ名をセッションデータに格納
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusPermanentRedirect, "/mypage")
}

func logout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Values["login"] = false // ログイン状態を無効化
	sess.Options = &sessions.Options{
		MaxAge: -1,
		Path:   "/",
	}
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusPermanentRedirect, "/")
}

func myPage(c echo.Context) error {
	sess, _ := session.Get("session", c)

	// ログインされていない場合は認証失敗を伝える
	if status, _ := sess.Values["login"]; status != true {
		return c.Redirect(http.StatusPermanentRedirect, "/")
	}

	// ログインされている場合は名前を表示
	name := sess.Values["name"]
	return c.Render(http.StatusOK, "main", name)
}

func main() {
	e := echo.New()
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = t
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.Use(middleware.Logger())
	e.Any("/", top)
	e.POST("/login", login)   // ログイン用ページ
	e.POST("/logout", logout) // ログアウト用ページ
	e.Any("/mypage", myPage)  // マイページ

	e.Logger.Fatal(e.Start(":1323"))
}
