package router

import (
	"database/sql"
	"echo_crud/pkg/controller"
	"echo_crud/pkg/repository"
	"echo_crud/pkg/usecase"

	"github.com/labstack/echo/v4"
)

// NewStudentRouter, fungsi untuk mengkonfigurasikan rute-rute terkait student dalam Echo framework
func NewStudentRouter(e *echo.Echo, g *echo.Group, db *sql.DB) {
	// membuat instance dari StudentRepository yang berinteraksi dengan database
	sr := repository.NewStudentRepository(db)
	// membuat instance dari StudentUsecase yang berfungsi sebagai perantara antara repository dan controller
	su := usecase.NewStudentUsecase(sr)
	// membuat instance dari StudentController yang mengatur logika untuk endpoint
	sc := &controller.StudentController{
		StudentUsecase: su,
	}

	// Mengatur rute HTTP menggunakan Echo untuk mendefinisikan endpoint
	g.POST("/student", sc.CreateStudent)       // POST /student, memanggil fungsi CreateStudent di StudentController untuk membuat data baru
	g.GET("/student", sc.GetStudent)           // GET /student, memanggil fungsi GetStudents di StudentController
	g.GET("/student/:id", sc.GetStudentById)   // GET /student/:id, memanggil fungsi GetStudent di StudentController dengan parameter ID
	g.PUT("/student/:id", sc.UpdateStudent)    // PUT /student/:id, memanggil fungsi UpdateStudent di StudentController dengan parameter ID
	g.DELETE("/student/:id", sc.DeleteStudent) // DELETE /student/:id, memanggil fungsi DeleteStudent di StudentController dengan parameter ID
}
