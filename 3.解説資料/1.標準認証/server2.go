package main

import (
	"crypto/subtle"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

// ハンドラ
func basicAuth2(c echo.Context) error {
	return c.String(http.StatusOK, "Authentication success.")
}

// 独自のSkipper
func mySkipper(c echo.Context) bool {
	// HTTPメソッドがGET以外の場合はミドルウェアをスキップ（実行しない）
	if c.Request().Method != "GET" {
		fmt.Println("Not execute.")
		return true
	}
	return false
}

// 独自のバリデータ
func myValidator(username, password string, c echo.Context) (bool, error) {
	// バリデータ関数が利用されているかのチェック
	fmt.Println("Execute authentication.")

	// 文字列比較の処理時間で認証情報を推測（タイミング攻撃）されないように、ConstantTimeCompare関数を利用
	if subtle.ConstantTimeCompare([]byte(username), []byte("ema")) == 1 &&
		subtle.ConstantTimeCompare([]byte(password), []byte("echo")) == 1 {
		return true, nil
	}
	return false, nil
}

func main() {
	e := echo.New()

	// BasicAuthWithConfigを直接使う
	e.Use(middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{
		Skipper:   mySkipper,
		Validator: myValidator,
		Realm:     "myRealm",
	}))

	e.Any("/", basicAuth2)
	e.Start(":1323")
}
