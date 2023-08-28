package controller

import (
	"echo_crud/pkg/domain"
	"echo_crud/pkg/dto"
	"echo_crud/shared/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// tipe struct yang mengendalikan HTTP handler terkait entitas student
type StudentController struct {
	StudentUsecase domain.StudentUsecase
}

// CreateStudent, method yang membuat data student baru berdasarkan data yang diberikan dalam request body
func (sc *StudentController) CreateStudent(c echo.Context) error {
	var studentdto dto.StudentDTO
	if err := c.Bind(&studentdto); err != nil {
		return response.SetResponse(c, http.StatusBadRequest, "bad request", nil)
	}
	if err := studentdto.Validation(); err != nil {
		return response.SetResponse(c, http.StatusBadRequest, err.Error(), nil)
	}
	if err := sc.StudentUsecase.CreateStudent(studentdto); err != nil {
		return response.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return response.SetResponse(c, http.StatusOK, "success added new student", nil)
}

// GetStudents, method yang mengambil daftar student dari usecase dan mengembalikan respons HTTP
func (sc *StudentController) GetStudents(c echo.Context) error {
	resp, err := sc.StudentUsecase.GetStudents()
	if err != nil {
		return response.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return response.SetResponse(c, http.StatusOK, "success search student", resp)
}

// GetStudent, method yang mengambil data student berdasarkan ID dari usecase dan mengembalikan respons HTTP
func (sc *StudentController) GetStudent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	resp, err := sc.StudentUsecase.GetStudent(id)
	if err != nil {
		return response.SetResponse(c, http.StatusNotFound, "id student not found", nil)
	}
	return response.SetResponse(c, http.StatusOK, "success search by id", resp)
}

// UpdateStudent, method yang mengupdate data student berdasarkan ID dan data yang diberikan dalam request body
func (sc *StudentController) UpdateStudent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// memeriksa apakah student dengan ID yang diberikan ada, dengan melakukan pengecekan terlebih dahulu dengan memanggil funsi GetStudent sebelum memanggil UpdateStudent
	_, err := sc.StudentUsecase.GetStudent(id)
	if err != nil {
		return response.SetResponse(c, http.StatusNotFound, "failed to update, id student not found", nil)
	}

	var studentdto dto.StudentDTO
	if err := c.Bind(&studentdto); err != nil {
		return response.SetResponse(c, http.StatusBadRequest, "bad request", nil)
	}
	if err := studentdto.Validation(); err != nil {
		return response.SetResponse(c, http.StatusBadRequest, err.Error(), nil)
	}
	if err := sc.StudentUsecase.UpdateStudent(id, studentdto); err != nil {
		return response.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return response.SetResponse(c, http.StatusOK, "success update student by id", nil)
}

// DeletetStudent, method yang mendelete data student berdasarkan ID
func (sc *StudentController) DeleteStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.SetResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	// Memeriksa apakah pengguna dengan ID yang diberikan ada di dalam database
	_, err = sc.StudentUsecase.GetStudent(id)
	if err != nil {
		return response.SetResponse(c, http.StatusNotFound, "failed to delete, id student not found", nil)
	}

	// Melakukan operasi penghapusan
	err = sc.StudentUsecase.DeleteStudent(id)
	if err != nil {
		return response.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	return response.SetResponse(c, http.StatusOK, "success delete user", nil)
}
