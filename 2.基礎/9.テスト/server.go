package main

import (
	"github.com/labstack/echo"
	"net/http"
)

type (
	User struct {
		Name  string `json:"name" form:"name"`
		Email string `json:"email" form:"email"`
	}
	Handler struct {
		db map[string]*User
	}
)

func (h *Handler) CreateUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

func (h *Handler) GetUser(c echo.Context) error {
	email := c.Param("email")
	user := h.db[email]
	if user == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	return c.JSON(http.StatusOK, user)
}

func main() {
	e := echo.New()
	//h := Handler{
	//	db: map[string]*User{
	//		"testUser":&User{
	//			Name:  "ema1",
	//			Email: "ema1@exsample.com",
	//		}},
	//}
	h := Handler{}
	e.GET("/users/:email", h.GetUser)
	e.POST("/", h.CreateUser)
	e.Logger.Fatal(e.Start(":1323"))
}
