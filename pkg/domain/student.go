package domain

import "echo_crud/pkg/dto"

// Student,tipe data yang mewakili data student yang disimpan dalam sistem
type Student struct {
	Id         int    `json:"id"`
	Fullname   string `json:"fullname"`
	Address    string `json:"address"`
	Birthdate  string `json:"birthdate"`
	Class      string `json:"class"`
	Batch      int    `json:"batch"`
	SchoolName string `json:"school_name"`
}

// StudentRepository, interface yang mendefinisikan operasi-operasi terhadap data student di dalam database
type StudentRepository interface {
	CreateStudent(req Student) error
	GetStudents() ([]Student, error)
	GetStudent(id int) (Student, error)
	UpdateStudent(id int, req Student) error
	DeleteStudent(id int) error
}

// StudentUsecase, interface yang mendefinisikan operasi-operasi terkait student
type StudentUsecase interface {
	CreateStudent(req dto.StudentDTO) error
	GetStudents() ([]Student, error)
	GetStudent(id int) (Student, error)
	UpdateStudent(id int, req dto.StudentDTO) error
	DeleteStudent(id int) error
}
