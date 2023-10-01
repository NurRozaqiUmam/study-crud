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

// @Tags USER
// @Summary Create a new user
// @Description Create a new user with the given details
// @Accept json
// @Produce json
// @Param request body dto.UserDTO true "Create User"
// @Success 200 {object} util.JsonReponse{message=string}
// @Failure 500 {object} util.JsonReponse{message=string}
// @Router /register [post]
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

// @Tags USER
// @Summary Get all users
// @Description Get all users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} util.JsonReponse
// @Failure 500 {object} util.JsonReponse
// @Router /user [get]
func (uc *UserController) GetUsers(c echo.Context) error {
	resp, err := uc.UserUsecase.GetUsers()
	if err != nil {
		return util.SetResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}
	return util.SetResponse(c, http.StatusOK, "success view all user", resp)
}

// @Tags USER
// @Summary Get user by id
// @Description Get user by id
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Security BearerAuth
// @Success 200 {object} util.JsonReponse
// @Failure 404 {object} util.JsonReponse
// @Router /user/{id} [get]
func (uc *UserController) GetUserById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	resp, err := uc.UserUsecase.GetUserById(id)

	if err != nil {
		return util.SetResponse(c, http.StatusNotFound, "id user not found", nil)
	}
	return util.SetResponse(c, http.StatusOK, "success search user by id", resp)
}

// @Tags USER
// @Summary Update user by id
// @Description Update user by id
// @Accept json
// @Produce json
// @Param id path int true "User ID"
//
// @Security BearerAuth
// @Success 200 {object} util.JsonReponse
// @Failure 400 {object} util.JsonReponse
// @Router /user/{id} [put]
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

// @Tags USER
// @Summary Delete user by id
// @Description Delete user by id
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Security BearerAuth
// @Success 200 {object} util.JsonReponse
// @Failure 404 {object} util.JsonReponse
// @Router /user/{id} [delete]
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

// @Tags USER
// @Summary Login user
// @Description Login user
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Login User"
// @Success 200 {object} util.JsonReponse
// @Failure 400 {object} util.JsonReponse
// @Router /login [post]
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
