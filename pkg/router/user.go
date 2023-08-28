package router

import (
	"database/sql"
	"echo_crud/pkg/controller"
	"echo_crud/pkg/middleware"
	"echo_crud/pkg/repository"
	"echo_crud/pkg/usecase"

	"github.com/labstack/echo/v4"
)

// NewStudentRouter, fungsi untuk mengkonfigurasikan rute-rute terkait student dalam Echo framework
func NewUserRouter(e *echo.Echo, g *echo.Group, db *sql.DB) {
	// membuat instance dari UserRepository yang berinteraksi dengan database
	ur := repository.NewUserRepository(db)
	// membuat instance dari UserUsecase yang berfungsi sebagai perantara antara repository dan controller
	uu := usecase.NewUserUsecase(ur)
	// membuat instance dari UserController yang mengatur logika untuk endpoint
	uc := &controller.UserController{
		UserUsecase: uu,
	}

	// Mengatur rute HTTP menggunakan Echo untuk mendefinisikan endpoint
	e.POST("/user", uc.CreateUser)                                   // POST /user, memanggil fungsi CreateStudent di StudentController untuk membuat data baru
	e.GET("/user", uc.GetUsers, middleware.IsAuthenticated)          // GET /user, memanggil fungsi GetStudents di StudentController
	e.GET("/user/:id", uc.GetUser, middleware.IsAuthenticated)       // GET /user/:id, memanggil fungsi GetStudent di StudentController dengan parameter ID
	e.PUT("/user/:id", uc.UpdateUser, middleware.IsAuthenticated)    // PUT /user/:id, memanggil fungsi UpdateStudent di StudentController dengan parameter ID
	e.DELETE("/user/:id", uc.DeleteUser, middleware.IsAuthenticated) // DELETE /user/:id, memanggil fungsi DeleteStudent di StudentController dengan parameter ID
	e.POST("/login", controller.CheckLogin)
}
