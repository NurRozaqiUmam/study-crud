package controller

import (
	"echo_crud/pkg/domain"
	"echo_crud/pkg/dto"
	"echo_crud/shared/util"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// tipe struct yang mengendalikan HTTP handler terkait entitas user
type UserController struct {
	UserUsecase domain.UserUsecase
}

// CreateUser, method yang membuat data user baru berdasarkan data yang diberikan dalam request body
func (uc *UserController) CreateUser(c echo.Context) error {
	var request dto.UserDTO
	if err := c.Bind(&request); err != nil {
		return util.SetResponse(c, http.StatusBadRequest, "bad request", nil)
	}

	if err := request.Validation(); err != nil {
		return util.SetResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	if err := uc.UserUsecase.CreateUser(request); err != nil {
		return util.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return util.SetResponse(c, http.StatusOK, "success added user", nil)
}

// GetUser, method yang mengambil daftar user dari usecase dan mengembalikan respons HTTP
func (uc *UserController) GetUsers(c echo.Context) error {
	resp, err := uc.UserUsecase.GetUsers()
	if err != nil {
		return util.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return util.SetResponse(c, http.StatusOK, "success view all user", resp)
}

// GetSUser, method yang mengambil data user berdasarkan ID dari usecase dan mengembalikan respons HTTP
func (uc *UserController) GetUserById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	resp, err := uc.UserUsecase.GetUserById(id)

	if err != nil {
		return util.SetResponse(c, http.StatusNotFound, "id user not found", nil)
	}
	return util.SetResponse(c, http.StatusOK, "success search user by id", resp)
}

// UpdateUser, method yang mengupdate data user berdasarkan ID dan data yang diberikan dalam request body
func (uc *UserController) UpdateUser(c echo.Context) error {
	var request dto.UserDTO

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return util.SetResponse(c, http.StatusBadRequest, "bad request", nil)
	}

	_, err = uc.UserUsecase.GetUserById(id)
	if err != nil {
		return util.SetResponse(c, http.StatusNotFound, "failed to update, user not found", nil)
	}

	if err := c.Bind(&request); err != nil {
		return util.SetResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	if err := request.Validation(); err != nil {
		return util.SetResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	if err := uc.UserUsecase.UpdateUser(id, request); err != nil {
		return util.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return util.SetResponse(c, http.StatusOK, "success update", nil)
}

// DeletetUser, method yang mendelete data user berdasarkan ID
func (uc *UserController) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	_, err := uc.UserUsecase.GetUserById(id)
	if err != nil {
		return util.SetResponse(c, http.StatusNotFound, "failed to delete, user not found", nil)
	}
	if err := uc.UserUsecase.DeleteUserById(id); err != nil {
		return util.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return util.SetResponse(c, http.StatusOK, "success delete", nil)
}

func (uc *UserController) Login(c echo.Context) error {
	var request dto.LoginRequest
	if err := c.Bind(&request); err != nil {
		return util.SetResponse(c, http.StatusBadRequest, "bad request", nil)
	}
	resp, err := uc.UserUsecase.UserLogin(request)
	if err != nil {
		return util.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return util.SetResponse(c, http.StatusOK, "success login", resp)
}
