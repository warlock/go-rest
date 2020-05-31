package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// User is...
type User struct {
	Name  string `json:"name" xml:"name"`
	Email string `json:"email" xml:"email"`
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
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

	e.Use(middleware.CORS())

	/*
			e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		  AllowOrigins: []string{"https://labstack.com", "https://labstack.net"},
		  AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		}))
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

	e.GET("/login", func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		if username != "josep" && password != "shh" {
			return c.String(http.StatusOK, "No no")
		}

		claims := &jwtCustomClaims{
			"Josep",
			true,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})

	})

	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte("secretpass"),
		TokenLookup: "header:Authorization",
	}))

	e.GET("/restricted", func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*jwtCustomClaims)
		name := claims.Name
		return c.String(http.StatusOK, "Hello Secure: "+name+" !!")
	})

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
