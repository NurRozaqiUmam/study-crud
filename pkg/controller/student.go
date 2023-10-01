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

// @Tags STUDENT
// @Summary Create a new student
// @Description Create a new student with the given details
// @Accept json
// @Produce json
// @Param request body dto.StudentDTO true "Create Student"
// @Security BearerAuth
// @Success 200 {object} util.JsonReponse{message=string}
// @Failure 500 {object} util.JsonReponse{message=string}
// @Router /student [post]
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

// @Tags STUDENT
// @Summary Get all students
// @Description Get all students
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} util.JsonReponse
// @Failure 500 {object} util.JsonReponse
// @Router /student [get]
func (sc *StudentController) GetStudent(c echo.Context) error {
	resp, err := sc.StudentUsecase.GetStudent()
	if err != nil {
		return util.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return util.SetResponse(c, http.StatusOK, "success", resp)
}

// @Tags STUDENT
// @Summary Get student by id
// @Description Get student by id
// @Accept json
// @Produce json
// @Param id path int true "Student ID"
// @Security BearerAuth
// @Success 200 {object} util.JsonReponse
// @Failure 500 {object} util.JsonReponse
// @Router /student/{id} [get]
func (sc *StudentController) GetStudentById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	resp, err := sc.StudentUsecase.GetStudentById(id)

	if err != nil {
		return util.SetResponse(c, http.StatusNotFound, "student id not found", nil)
	}
	return util.SetResponse(c, http.StatusOK, "success", resp)
}

// @Tags STUDENT
// @Summary Update student
// @Description Update student with the given details
// @Accept json
// @Produce json
// @Param id path int true "Student ID"
// @Param request body dto.StudentDTO true "Update Student"
// @Security BearerAuth
// @Success 200 {object} util.JsonReponse{message=string}
// @Failure 500 {object} util.JsonReponse{message=string}
// @Router /student/{id} [put]
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

// @Tags STUDENT
// @Summary Delete student
// @Description Delete student with the given details
// @Accept json
// @Produce json
// @Param id path int true "Student ID"
// @Security BearerAuth
// @Success 200 {object} util.JsonReponse{message=string}
// @Failure 500 {object} util.JsonReponse{message=string}
// @Router /student/{id} [delete]
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
