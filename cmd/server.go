package cmd

import (
	"echo_crud/config"
	"echo_crud/pkg/router"
	"echo_crud/shared/db"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// RunServer, fungsi yang menjalankan server
func RunServer() {
	e := echo.New()
	g := e.Group("")
	conf := config.GetConfig()
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		TokenLookup: "header:Authorization",
		AuthScheme:  "Bearer",
		ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
			keyFunc := func(t *jwt.Token) (interface{}, error) {
				if t.Method.Alg() != "HS512" {
					return nil, fmt.Errorf("jwt token is formatted incorrectly")
				}
				return []byte(conf.SignKey), nil
			}
			token, err := jwt.Parse(auth, keyFunc)
			if err != nil {
				return nil, err
			}
			if !token.Valid {
				return nil, err
			}
			return token, err
		},
	}))
	Apply(e, g, conf)
	e.Logger.Error(e.Start(":5000"))
}

// Apply, fungsi yang mengaplikasikan konfigurasi rute-rute dalam Echo framework
func Apply(e *echo.Echo, g *echo.Group, conf config.Configuration) {
	db := db.NewInstanceDb(conf)      // membuat instance dari database yang akan digunakan
	router.NewUserRouter(e, g, db)    // memanggil fungsi NewUserRouter untuk mengonfigurasi rute terkait user
	router.NewStudentRouter(e, g, db) // memanggil fungsi NewStudentRouter untuk mengonfigurasi rute terkait student
}
