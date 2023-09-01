package repository

import (
	"database/sql"
	"echo_crud/pkg/domain"
)

// tipe yang bertanggung jawab atas interaksi dengan database terkait data student
type StudentRepository struct {
	db *sql.DB // instance dari koneksi database
}

// NewStudentRepository, fungsi pembuat untuk membuat instance baru dari StudentRepository
func NewStudentRepository(db *sql.DB) domain.StudentRepository {
	return &StudentRepository{
		db: db,
	}
}

// CreateStudent, metode untuk membuat data student baru dalam database
func (sr *StudentRepository) CreateStudent(req domain.Student) error {
	sql := `INSERT INTO student (fullname, address, birthdate, class, batch, school_name) values ($1, $2, $3, $4, $5, $6)`
	stmt, err := sr.db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err2 := stmt.Exec(req.Fullname, req.Address, req.Birthdate, req.Class, req.Batch, req.SchoolName)
	if err2 != nil {
		return err2
	}
	return nil
}

// GetStudents, metode yang mengambil daftar student dari database
func (sr *StudentRepository) GetStudent() ([]domain.Student, error) {
	sql := `SELECT * FROM student`
	rows, err := sr.db.Query(sql)
	if err != nil {
		return nil, err
	}
	var students []domain.Student
	for rows.Next() {
		student := domain.Student{}
		err2 := rows.Scan(&student.Id, &student.Fullname, &student.Address, &student.Birthdate, &student.Class, &student.Batch, &student.SchoolName)
		if err2 != nil {
			return students, err2
		}
		students = append(students, student)
	}
	return students, err
}

// GetStudent, metode yang mengambil data student berdasarkan ID dari database
func (sr *StudentRepository) GetStudentById(id int) (student domain.Student, err error) {
	sql := `SELECT * FROM student WHERE id = $1`
	err = sr.db.QueryRow(sql, id).Scan(&student.Id, &student.Fullname, &student.Address, &student.Birthdate, &student.Class, &student.Batch, &student.SchoolName)
	return student, err
}

// UpdateStudent, metode untuk memperbarui data student dalam database
func (sr *StudentRepository) UpdateStudent(id int, req domain.Student) error {
	sql := `UPDATE student SET fullname = $1, address = $2, birthdate = $3, class = $4, batch = $5, school_name = $6 WHERE id = $7`
	_, err2 := sr.db.Exec(sql, req.Fullname, req.Address, req.Birthdate, req.Class, req.Batch, req.SchoolName, id)
	if err2 != nil {
		return err2
	}
	return nil
}

// DeleteStudent, metode untuk menghapus data student dari database
func (sr *StudentRepository) DeleteStudentById(id int) error {
	sql := `DELETE FROM student WHERE id = $1`
	_, err2 := sr.db.Exec(sql, id)
	if err2 != nil {
		return err2
	}
	return nil
}
