//package main
//
//import (
//	"github.com/gorilla/sessions"
//	"github.com/labstack/echo"
//	"github.com/labstack/echo-contrib/session"
//	"strconv"
//
//	"fmt"
//	"net/http"
//)
//
//func setSession(c echo.Context) error {
//	sess, _ := session.Get("session", c)
//	sess.Options = &sessions.Options{
//		Path:     "/",
//		MaxAge:   7,
//		HttpOnly: true,
//	}
//	loadCount := sess.Values["loadCount"]
//	if loadCount == nil {
//		loadCount = 0
//	}
//	sess.Values["loadCount"] = countUp(loadCount.(int))
//	message := fmt.Sprintf("あなたはこのページを %v 回表示しています。この値はsessionを参照しています。" , sess.Values["loadCount"])
//	sess.Save(c.Request(), c.Response())
//
//	// ここからクッキー
//	cookie, _ := c.Cookie("loadCount")
//	if cookie == nil {
//		cookie := new(http.Cookie)
//		cookie.Name = "loadCount"
//		cookie.Value = "0"
//		c.SetCookie(cookie)
//	}
//	cookieLoadCount :=  cookie.Value
//	cc, _ := strconv.Atoi(cookieLoadCount)
//	cookie.Value = strconv.Itoa(countUp(cc))
//	c.SetCookie(cookie)
//
//	return c.HTML(http.StatusOK,
//		"<div>"+ message +"</div>" +
//		"<br>" +
//		"<div>あなたはこのページを <span id='result'></span> 回表示しています。この値はcookieを参照しています。</div>\n" +
//		"<input id='cookie' />\n" +
//		"<button id='button'>クッキーの値を変更</button>\n" +
//		"<script>\n" +
//		"var cookies = document.cookie\n" +
//		"var cookiesArray = cookies.split(';')\n" +
//		"for(var c of cookiesArray){\n" +
//		"	var cArray = c.split('=')\n" +
//		"	if( cArray[0] == ' loadCount'){\n" +
//		"		document.getElementById('cookie').value = cArray[1]\n" +
//		"		document.getElementById('result').textContent = cArray[1]\n" +
//		"	}\n" +
//		"}\n" +
//		"document.getElementById('button').onclick = function() {\n" +
//		"	b = document.getElementById('cookie').value\n" +
//		"	document.cookie = 'loadCount='+b\n" +
//		"	document.getElementById('result').textContent = b\n" +
//		"}\n" +
//		"</script>")
//}
//
//func countUp(counter int) int {
//	return counter + 1
//}
//
//func main() {
//	e := echo.New()
//	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
//	e.GET("/session", setSession)
//	e.Start(":1323")
//}
//
