package usecase

import (
	"echo_crud/pkg/domain"
	"echo_crud/pkg/dto"

	"github.com/mitchellh/mapstructure"
)

// tipe yang bertanggung jawab atas operasi-operasi terkait student
type StudentUsecase struct {
	StudentRepository domain.StudentRepository // Mengandung dependency ke repository untuk mengakses data student
}

// NewStudentUsecase, fungsi pembuat untuk membuat instance baru dari StudentUsecase
func NewStudentUsecase(studentRepository domain.StudentRepository) domain.StudentUsecase {
	return &StudentUsecase{
		StudentRepository: studentRepository,
	}
}

// CreateStudent, metode untuk membuat data student baru
func (su StudentUsecase) CreateStudent(req dto.StudentDTO) error {
	var student domain.Student
	if err := mapstructure.Decode(req, &student); err != nil {
		return err
	}
	return su.StudentRepository.CreateStudent(student)
}

// GetStudents, metode yang mengambil daftar student dari database
func (su StudentUsecase) GetStudent() ([]domain.Student, error) {
	return su.StudentRepository.GetStudent()
}

// GetStudent, metode yang mengambil data student berdasarkan ID
func (su StudentUsecase) GetStudentById(id int) (domain.Student, error) {
	return su.StudentRepository.GetStudentById(id)
}

// UpdateStudent, metode untuk memperbarui data student
func (su StudentUsecase) UpdateStudent(id int, req dto.StudentDTO) error {
	var student domain.Student
	if err := mapstructure.Decode(req, &student); err != nil {
		return err
	}
	return su.StudentRepository.UpdateStudent(id, student)
}

// DeleteStudent, metode untuk menghapus data student
func (su StudentUsecase) DeleteStudentById(id int) error {
	return su.StudentRepository.DeleteStudentById(id)
}
