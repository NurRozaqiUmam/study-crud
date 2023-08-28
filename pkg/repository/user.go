package repository

import (
	"database/sql"
	"echo_crud/pkg/domain"
	"echo_crud/pkg/helpers"
	"errors"
)

// tipe yang bertanggung jawab atas interaksi dengan database terkait data user
type UserRepository struct {
	db *sql.DB // instance dari koneksi database
}

// NewUserRepository, fungsi pembuat untuk membuat instance baru dari UserRepository
func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &UserRepository{
		db: db,
	}
}

// CreateUser, metode untuk membuat data user baru dalam database
func (ur UserRepository) CreateUser(req domain.User) error {
	// Cek apakah email sudah terdaftar
	if ur.isEmailExists(req.Email) {
		return errors.New("email already exists")
	}

	// Hash password sebelum memasukkan ke database
	hashedPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		return err
	}

	sql := `INSERT INTO users (username, email, password, address) VALUES ($1, $2, $3, $4)`
	stmt, err := ur.db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(req.Username, req.Email, hashedPassword, req.Address)
	if err2 != nil {
		return err2
	}
	return nil
}

func (ur UserRepository) isEmailExists(email string) bool {
	sql := "SELECT COUNT(*) FROM users WHERE email = $1"
	var count int
	err := ur.db.QueryRow(sql, email).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

// GetUser, metode yang mengambil daftar user dari database
func (ur UserRepository) GetUsers() ([]domain.User, error) {
	// query SQL untuk mengambil semua data user
	sql := `SELECT * FROM users`
	rows, err := ur.db.Query(sql)
	var users []domain.User
	for rows.Next() {
		var user domain.User
		// membaca baris hasil query dan memasukkan ke dalam struktur user
		err2 := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Address)
		if err2 != nil {
			return nil, err2
		}
		users = append(users, user)
	}
	return users, err
}

// GetUser, metode yang mengambil data user berdasarkan ID dari database
func (ur UserRepository) GetUser(id int) (domain.User, error) {
	var user domain.User
	// query SQL untuk mengambil data user berdasarkan ID
	sql := `SELECT * FROM users WHERE id = $1`
	err := ur.db.QueryRow(sql, id).Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Address)
	return user, err
}

// UpdateUser, metode untuk memperbarui data user dalam database
func (ur UserRepository) UpdateUser(id int, req domain.User) error {
	// Hash password baru sebelum memperbarui data user
	hashedPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		return err
	}

	// Query SQL untuk memperbarui data user berdasarkan ID
	sql := `UPDATE users SET username = $1, email = $2, password = $3, address = $4 WHERE id = $5`
	stmt, err := ur.db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(req.Username, req.Email, hashedPassword, req.Address, req.Id)
	if err2 != nil {
		return err2
	}
	return nil
}

// DeleteUser, metode untuk menghapus data user dari database
func (ur UserRepository) DeleteUser(id int) error {
	// query SQL untuk menghapus data user berdasarkan ID
	sql := `DELETE FROM users WHERE id = $1`
	_, err := ur.db.Exec(sql, id)
	return err
}
