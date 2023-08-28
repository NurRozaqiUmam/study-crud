package usecase

import (
	"echo_crud/pkg/domain"
	"echo_crud/pkg/dto"

	"github.com/mitchellh/mapstructure"
)

// tipe yang bertanggung jawab atas operasi-operasi terkait user
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
	// mengkonversi UsertDTO menjadi tipe User dan memasukkannya ke dalam database
	mapstructure.Decode(req, &user)
	return uu.UserRepository.CreateUser(user)
}

// GetUsers, metode yang mengambil daftar user dari database
func (uu UserUsecase) GetUsers() ([]domain.User, error) {
	// memanggil metode GetUsers dari UserRepository untuk mengambil data dari database
	return uu.UserRepository.GetUsers()
}

// GetUser, metode yang mengambil data user berdasarkan ID
func (uu UserUsecase) GetUser(id int) (domain.User, error) {
	// memanggil metode GetUser dari UserRepository untuk mengambil data dari database
	return uu.UserRepository.GetUser(id)
}

// UpdateUser, metode untuk memperbarui data user
func (uu UserUsecase) UpdateUser(id int, req dto.UserDTO) error {
	var user domain.User
	// mengkonversi UserDTO menjadi tipe User dan memperbarui data user di database
	mapstructure.Decode(req, &user)
	user.Id = id // menetapkan ID user yang akan diupdate
	return uu.UserRepository.UpdateUser(id, user)
}

// DeleteUser, metode untuk menghapus data user
func (uu UserUsecase) DeleteUser(id int) error {
	// memanggil metode DeleteUser dari UserRepository untuk menghapus data dari database
	return uu.UserRepository.DeleteUser(id)
}
