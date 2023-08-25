package dto

import validation "github.com/go-ozzo/ozzo-validation"

// StudentDTO, tipe data yang digunakan untuk membungkus data student yang dikirimkan melalui permintaan HTTP
type StudentDTO struct {
	// field dalam StudentDTO yang akan dikirimkan melalui JSON dalam permintaan HTTP
	Fullname   string `json:"fullname"`
	Address    string `json:"address"`
	Birthdate  string `json:"birthdate"`
	Class      string `json:"class"`
	Batch      int    `json:"batch"`
	SchoolName string `json:"school_name"`
}

// Validation, method yang digunakan untuk melakukan validasi pada data yang ada dalam StudentDTO
func (s StudentDTO) Validation() error {
	// menggunakan package "validation" untuk melakukan validasi terstruktur pada field dalam StudentDTO
	// validasi ini memastikan bahwa data yang diterima dari permintaan HTTP memenuhi kriteria yang diharapkan
	err := validation.ValidateStruct(&s,
		validation.Field(&s.Fullname, validation.Required),
		validation.Field(&s.Birthdate, validation.Required),
		validation.Field(&s.Class, validation.Required),
		validation.Field(&s.Batch, validation.Required),
		validation.Field(&s.SchoolName, validation.Required))

	if err != nil {
		return err
	}
	return nil
}
