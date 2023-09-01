package router

import (
	"database/sql"
	"echo_crud/pkg/controller"
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
	e.POST("/register", uc.CreateUser) // POST /user, memanggil fungsi CreateStudent di StudentController untuk membuat data baru
	e.POST("/login", uc.Login)
	g.GET("/user", uc.GetUsers)          // GET /user, memanggil fungsi GetStudents di StudentController
	g.GET("/user/:id", uc.GetUserById)   // GET /user/:id, memanggil fungsi GetStudent di StudentController dengan parameter ID
	g.PUT("/user/:id", uc.UpdateUser)    // PUT /user/:id, memanggil fungsi UpdateStudent di StudentController dengan parameter ID
	g.DELETE("/user/:id", uc.DeleteUser) // DELETE /user/:id, memanggil fungsi DeleteStudent di StudentController dengan parameter ID
}
