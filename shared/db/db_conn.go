package db

import (
	"database/sql"
	"echo_crud/config"
	"fmt"

	_ "github.com/lib/pq"
)

func NewInstanceDb(conf config.Configuration) *sql.DB {
	// mengambil konfigurasi database dari package config

	// membuka koneksi baru ke database PostgreSQL dengan menggunakan informasi konfigurasi
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conf.DBHost, conf.DBPort, conf.DBUsername, conf.DBPassword, conf.DBName))
	if err != nil {
		panic(err) // Jika terjadi kesalahan dalam membuka koneksi, akan menghentikan program dan menampilkan error.
	}

	// melakukan ping ke database untuk memastikan koneksi berhasil
	err = db.Ping()
	if err != nil {
		panic(err) // jika ping ke database gagal, akan menghentikan program dan menampilkan error
	}

	return db // mengembalikan instance koneksi database yang telah dikonfigurasi
}
