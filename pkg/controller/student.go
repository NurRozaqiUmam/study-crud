package controller

import (
	"echo_crud/pkg/domain"
	"echo_crud/pkg/dto"
	"echo_crud/shared/util"
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
	var request dto.StudentDTO
	if err := c.Bind(&request); err != nil {
		return util.SetResponse(c, http.StatusBadRequest, "bad request", nil)
	}

	if err := request.Validation(); err != nil {
		return util.SetResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	if err := sc.StudentUsecase.CreateStudent(request); err != nil {
		return util.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return util.SetResponse(c, http.StatusOK, "success", nil)
}

// GetStudents, method yang mengambil daftar student dari usecase dan mengembalikan respons HTTP
func (sc *StudentController) GetStudent(c echo.Context) error {
	resp, err := sc.StudentUsecase.GetStudent()
	if err != nil {
		return util.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return util.SetResponse(c, http.StatusOK, "success", resp)
}

// GetStudent, method yang mengambil data student berdasarkan ID dari usecase dan mengembalikan respons HTTP
func (sc *StudentController) GetStudentById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	resp, err := sc.StudentUsecase.GetStudentById(id)

	if err != nil {
		return util.SetResponse(c, http.StatusNotFound, "student id not found", nil)
	}
	return util.SetResponse(c, http.StatusOK, "success", resp)
}

// UpdateStudent, method yang mengupdate data student berdasarkan ID dan data yang diberikan dalam request body
func (sc *StudentController) UpdateStudent(c echo.Context) error {
	var request dto.StudentDTO

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return util.SetResponse(c, http.StatusBadRequest, "bad request", nil)
	}

	_, err = sc.StudentUsecase.GetStudentById(id)
	if err != nil {
		return util.SetResponse(c, http.StatusNotFound, "student not found", nil)
	}

	if err := c.Bind(&request); err != nil {
		return util.SetResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	if err := request.Validation(); err != nil {
		return util.SetResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	if err := sc.StudentUsecase.UpdateStudent(id, request); err != nil {
		return util.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return util.SetResponse(c, http.StatusOK, "success", nil)
}

// DeletetStudent, method yang mendelete data student berdasarkan ID
func (sc *StudentController) DeleteStudent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := sc.StudentUsecase.GetStudentById(id)
	if err != nil {
		return util.SetResponse(c, http.StatusNotFound, "id not found", nil)
	}
	if err := sc.StudentUsecase.DeleteStudentById(id); err != nil {
		return util.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return util.SetResponse(c, http.StatusOK, "success", nil)
}
