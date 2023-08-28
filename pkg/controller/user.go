package controller

import (
	"echo_crud/pkg/domain"
	"echo_crud/pkg/dto"
	"echo_crud/shared/response"
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
	var userdto dto.UserDTO
	if err := c.Bind(&userdto); err != nil {
		return response.SetResponse(c, http.StatusBadRequest, "bad request", nil)
	}
	if err := userdto.Validation(); err != nil {
		return response.SetResponse(c, http.StatusBadRequest, err.Error(), nil)
	}
	if err := uc.UserUsecase.CreateUser(userdto); err != nil {
		return response.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return response.SetResponse(c, http.StatusOK, "user created successfully", nil)
}

// GetUser, method yang mengambil daftar user dari usecase dan mengembalikan respons HTTP
func (uc *UserController) GetUsers(c echo.Context) error {
	resp, err := uc.UserUsecase.GetUsers()
	if err != nil {
		return response.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return response.SetResponse(c, http.StatusOK, "success search user", resp)
}

// GetSUser, method yang mengambil data user berdasarkan ID dari usecase dan mengembalikan respons HTTP
func (uc *UserController) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	resp, err := uc.UserUsecase.GetUser(id)
	if err != nil {
		return response.SetResponse(c, http.StatusNotFound, "id user not found", nil)
	}
	return response.SetResponse(c, http.StatusOK, "success search user by id", resp)
}

// UpdateUser, method yang mengupdate data user berdasarkan ID dan data yang diberikan dalam request body
func (uc *UserController) UpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// memeriksa apakah user dengan ID yang diberikan ada, dengan melakukan pengecekan terlebih dahulu dengan memanggil fungsi GetUser sebelum memanggil UpdateUser
	_, err := uc.UserUsecase.GetUser(id)
	if err != nil {
		return response.SetResponse(c, http.StatusNotFound, "failed to update, id user not found", nil)
	}

	var userdto dto.UserDTO
	if err := c.Bind(&userdto); err != nil {
		return response.SetResponse(c, http.StatusBadRequest, "bad request", nil)
	}
	if err := userdto.Validation(); err != nil {
		return response.SetResponse(c, http.StatusBadRequest, err.Error(), nil)
	}
	if err := uc.UserUsecase.UpdateUser(id, userdto); err != nil {
		return response.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return response.SetResponse(c, http.StatusOK, "success update user by id", nil)
}

// DeletetUser, method yang mendelete data user berdasarkan ID
func (uc *UserController) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.SetResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	// Memeriksa apakah user dengan ID yang diberikan ada di dalam database
	_, err = uc.UserUsecase.GetUser(id)
	if err != nil {
		return response.SetResponse(c, http.StatusNotFound, "failed to delete, id user not found", nil)
	}

	// Melakukan operasi penghapusan
	err = uc.UserUsecase.DeleteUser(id)
	if err != nil {
		return response.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	return response.SetResponse(c, http.StatusOK, "success delete user", nil)
}
