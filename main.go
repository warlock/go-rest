package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// User is...
type User struct {
	Name  string `json:"name" xml:"name"`
	Email string `json:"email" xml:"email"`
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.File("/favicon.ico", "images/favicon.ico")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})

	e.GET("/html", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<b>Hello<br />World </b>")
	})

	e.GET("/no-content", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})

	e.GET("/json", func(c echo.Context) error {
		u := &User{
			Name:  "Jon",
			Email: "jon@labstack.com",
		}
		return c.JSON(http.StatusOK, u)
	})

	e.Static("/static", "assets")

	e.Start(":8080")
}
