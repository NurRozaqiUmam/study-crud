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
	// mengkonversi StudentDTO menjadi tipe Student dan memasukkannya ke dalam database
	mapstructure.Decode(req, &student)
	return su.StudentRepository.CreateStudent(student)
}

// GetStudents, metode yang mengambil daftar student dari database
func (su StudentUsecase) GetStudents() ([]domain.Student, error) {
	// memanggil metode GetStudents dari StudentRepository untuk mengambil data dari database
	return su.StudentRepository.GetStudents()
}

// GetStudent, metode yang mengambil data student berdasarkan ID
func (su StudentUsecase) GetStudent(id int) (domain.Student, error) {
	// memanggil metode GetStudent dari StudentRepository untuk mengambil data dari database
	return su.StudentRepository.GetStudent(id)
}

// UpdateStudent, metode untuk memperbarui data student
func (su StudentUsecase) UpdateStudent(id int, req dto.StudentDTO) error {
	var student domain.Student
	// mengkonversi StudentDTO menjadi tipe Student dan memperbarui data student di database
	mapstructure.Decode(req, &student)
	student.Id = id // menetapkan ID student yang akan diupdate
	return su.StudentRepository.UpdateStudent(id, student)
}

// DeleteStudent, metode untuk menghapus data student
func (su StudentUsecase) DeleteStudent(id int) error {
	// memanggil metode DeleteStudent dari StudentRepository untuk menghapus data dari database
	return su.StudentRepository.DeleteStudent(id)
}
