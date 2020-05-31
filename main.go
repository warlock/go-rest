package main

import (
	"net/http"
	"strconv"

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

	/*
		// Login personalitzat
			myLog, err := os.OpenFile("logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
				log.Fatalf("No es pot llegir el document: %v", err)
			}
			defer myLog.Close()

			logConfig := middleware.LoggerConfig {
				Output: myLog,
				Format: "method=${method}, uri=${uri}, status=${status}\n"
			}

			e.Use(middleware.LoggerWithConfig(logConfig))
	*/

	e.Use(middleware.Logger())

	// Reinicia en cas de que faci un panic
	e.Use(middleware.Recover())

	// Redireccio https
	//e.Pre(middleware.HTTPSRedirect())

	e.File("/favicon.ico", "images/favicon.ico")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})

	e.GET("/html", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<b>Hello<br />World </b>")
	})

	e.GET("/file", func(c echo.Context) error {
		return c.File("images/favicon.ico")
	})

	e.GET("/fileinline", func(c echo.Context) error {
		return c.Inline("images/favicon.ico", "file.ico")
		// Content disposition Inline
	})

	e.GET("/fileattachment", func(c echo.Context) error {
		return c.Attachment("images/favicon.ico", "file.ico")
		// Content disposition attachment
	})

	e.GET("/no-content", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})

	e.GET("redirect", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "http://girona.dev")
	})

	e.GET("redirect2", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/json")
	})

	e.GET("/json", func(c echo.Context) error {
		u := User{
			Name:  "Jon",
			Email: "jon@labstack.com",
		}
		return c.JSON(http.StatusOK, u)
	})

	e.GET("/jsonlist", func(c echo.Context) error {
		ps := make([]User, 0)
		u := User{
			Name:  "Jon",
			Email: "jon@labstack.com",
		}
		ps = append(ps, u)
		ps = append(ps, u)
		ps = append(ps, u)
		return c.JSON(http.StatusOK, ps)
	})

	e.GET("/xml", func(c echo.Context) error {
		u := User{
			Name:  "Jon",
			Email: "jon@labstack.com",
		}
		return c.XML(http.StatusOK, u)
	})

	e.GET("/name/:name", func(c echo.Context) error {
		p := c.Param("name")
		if p != "" {
			return c.String(http.StatusOK, "Hello")

		}
		return c.String(http.StatusOK, "Hello world")
	})

	e.GET("/operacio", func(c echo.Context) error {
		d := c.QueryParam("d")
		f, _ := strconv.Atoi(d)
		a := 25000 / f
		return c.String(http.StatusOK, strconv.Itoa(a))
	})

	e.Static("/static", "assets")

	e.Start(":8080")

	/*
		   e.Logger.Fatal(
		   	e.StartTLS(":443", "./cert/cert.pem","./cert/key.key")
			 )

			 go func () {
				 e.Logger.Fatal(
					 e.Start(":80")
				 )
			 }
	*/

}
