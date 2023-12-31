package domain

import "echo_crud/pkg/dto"

// User,tipe data yang mewakili data user yang disimpan dalam sistem
type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

// UserRepository, interface yang mendefinisikan operasi-operasi terhadap data user di dalam database
type UserRepository interface {
	CreateUser(req User) error
	UpdateUser(id int, req User) error
	GetUsers() ([]User, error)
	GetUserById(id int) (User, error)
	DeleteUserById(id int) error
	FindByEmail(email string) (User, error)
}

// UserUsecase, interface yang mendefinisikan operasi-operasi terkait user
type UserUsecase interface {
	CreateUser(req dto.UserDTO) error
	UpdateUser(id int, req dto.UserDTO) error
	GetUsers() ([]User, error)
	GetUserById(id int) (User, error)
	DeleteUserById(id int) error
	UserLogin(req dto.LoginRequest) (interface{}, error)
}
