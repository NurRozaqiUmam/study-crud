package dto

import validation "github.com/go-ozzo/ozzo-validation"

// UserDTO, tipe data yang digunakan untuk membungkus data user yang dikirimkan melalui permintaan HTTP
type UserDTO struct {
	// field dalam UserDTO yang akan dikirimkan melalui JSON dalam permintaan HTTP
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

// Validation, method yang digunakan untuk melakukan validasi pada data yang ada dalam UserDTO
func (u UserDTO) Validation() error {
	// menggunakan package "validation" untuk melakukan validasi terstruktur pada field dalam UserDTO
	// validasi ini memastikan bahwa data yang diterima dari permintaan HTTP memenuhi kriteria yang diharapkan
	err := validation.ValidateStruct(&u,
		validation.Field(&u.Username, validation.Required),
		validation.Field(&u.Email, validation.Required),
		validation.Field(&u.Password, validation.Required),
		validation.Field(&u.Address, validation.Required))

	if err != nil {
		return err
	}
	return nil
}
