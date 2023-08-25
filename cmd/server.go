package cmd

import (
	"echo_crud/pkg/router"
	"echo_crud/shared/db"

	"github.com/labstack/echo/v4"
)

// RunServer, fungsi yang menjalankan server
func RunServer() {
	e := echo.New() // membuat instance baru dari Echo framework
	g := e.Group("")

	Apply(e, g)                      // memanggil fungsi Apply untuk mengonfigurasi rute-rute
	e.Logger.Error(e.Start(":5000")) // memulai server dan menangani kesalahan jika terjadi
}

// Apply, fungsi yang mengaplikasikan konfigurasi rute-rute dalam Echo framework
func Apply(e *echo.Echo, g *echo.Group) {
	db := db.NewInstanceDb()          // membuat instance dari database yang akan digunakan
	router.NewStudentRouter(e, g, db) // memanggil fungsi NewStudentRouter untuk mengonfigurasi rute terkait student
}
