package usecase

import (
	"echo_crud/pkg/domain"
	"echo_crud/pkg/dto"
	"echo_crud/shared/util"
	"errors"

	"github.com/mitchellh/mapstructure"
)

// tipe yang bertanggung jawab atas operasi-operasi terkait user.
type UserUsecase struct {
	UserRepository domain.UserRepository // Mengandung dependency ke repository untuk mengakses data user
}

// NewUserUsecase, fungsi pembuat untuk membuat instance baru dari UserUsecase
func NewUserUsecase(userRepository domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{
		UserRepository: userRepository,
	}
}

// CreateUser, metode untuk membuat data user baru
func (uu UserUsecase) CreateUser(req dto.UserDTO) error {
	var user domain.User
	mapstructure.Decode(req, &user)
	if _, err := uu.UserRepository.FindByEmail(user.Email); err == nil {
		return errors.New("email already exists")
	}
	user.Password = util.EncryptPassword(user.Password)
	return uu.UserRepository.CreateUser(user)
}

func (u UserUsecase) UserLogin(req dto.LoginRequest) (interface{}, error) {
	var loginResponse dto.LoginResponse
	user, err := u.UserRepository.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("email not found")
	}
	passwordValid := util.DecryptPassword(user.Password)
	if passwordValid != req.Password {
		return nil, errors.New("bad credential")
	}
	token, err := util.CreateJwtToken(user)
	if err != nil {
		return nil, err
	}
	mapstructure.Decode(user, &loginResponse)
	loginResponse.Token = token
	return loginResponse, nil
}

// GetUsers, metode yang mengambil daftar user dari database
func (uu UserUsecase) GetUsers() ([]domain.User, error) {
	// memanggil metode GetUsers dari UserRepository untuk mengambil data dari database
	return uu.UserRepository.GetUsers()
}

// GetUser, metode yang mengambil data user berdasarkan ID
func (uu UserUsecase) GetUserById(id int) (domain.User, error) {
	// memanggil metode GetUser dari UserRepository untuk mengambil data dari database
	return uu.UserRepository.GetUserById(id)
}

// UpdateUser, metode untuk memperbarui data user
func (uu UserUsecase) UpdateUser(id int, req dto.UserDTO) error {
	var user domain.User
	mapstructure.Decode(req, &user)

	// Jika kata sandi baru diberikan, enkripsi kata sandi tersebut
	if req.Password != "" {
		user.Password = util.EncryptPassword(req.Password)
	}

	return uu.UserRepository.UpdateUser(id, user)
}

// DeleteUser, metode untuk menghapus data user
func (uu UserUsecase) DeleteUserById(id int) error {
	// memanggil metode DeleteUser dari UserRepository untuk menghapus data dari database
	return uu.UserRepository.DeleteUserById(id)
}
