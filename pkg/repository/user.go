package repository

import (
	"database/sql"
	"echo_crud/pkg/domain"
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
	sql := `INSERT INTO users (username, email, password, address) values ($1, $2, $3, $4)`
	_, err2 := ur.db.Exec(sql, req.Username, req.Email, req.Password, req.Address)
	if err2 != nil {
		return err2
	}
	return nil
}

// FindByEmail, metode yang mencari daftar user berdasarkan email dari database
func (ur UserRepository) FindByEmail(email string) (domain.User, error) {
	var user domain.User
	sql := `SELECT * FROM users WHERE email = $1`
	err := ur.db.QueryRow(sql, email).Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Address)
	return user, err
}

// GetUser, metode yang mengambil daftar user dari database
func (ur UserRepository) GetUsers() ([]domain.User, error) {
	sql := `SELECT * FROM users`
	rows, err := ur.db.Query(sql)
	if err != nil {
		return nil, err
	}
	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err2 := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Address)
		if err2 != nil {
			return users, err2
		}
		users = append(users, user)
	}
	return users, err
}

// GetUser, metode yang mengambil data user berdasarkan ID dari database
func (ur UserRepository) GetUserById(id int) (domain.User, error) {
	var user domain.User
	sql := `SELECT * FROM users WHERE id = $1`
	err := ur.db.QueryRow(sql, id).Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Address)
	return user, err
}

// UpdateUser, metode untuk memperbarui data user dalam database
func (ur UserRepository) UpdateUser(id int, req domain.User) error {
	sql := `UPDATE users SET username = $1, email = $2, password = $3, address = $4 WHERE id = $5`
	_, err2 := ur.db.Exec(sql, req.Username, req.Email, req.Password, req.Address, id)
	if err2 != nil {
		return err2
	}
	return nil
}

// DeleteUser, metode untuk menghapus data user dari database
func (ur UserRepository) DeleteUserById(id int) error {
	sql := `DELETE FROM users WHERE id = $1`
	_, err2 := ur.db.Exec(sql, id)
	if err2 != nil {
		return err2
	}
	return nil
}
