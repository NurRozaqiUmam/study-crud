package models

import (
	"database/sql"
	"echo_crud/pkg/helpers"
	"echo_crud/shared/db"
	"fmt"
)

type Useer struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

func CheckLogin(email, password string) (bool, error) {
	var obj Useer
	var pwd string

	db := db.NewInstanceDb()

	sqlStatement := "SELECT * FROM users WHERE email = $1"

	err := db.QueryRow(sqlStatement, email).Scan(
		&obj.Id, &obj.Username, &obj.Email, &pwd, &obj.Address,
	)

	if err == sql.ErrNoRows {
		fmt.Println("Email not found")
		return false, err
	}

	if err != nil {
		fmt.Println("Query error:", err)
		return false, err
	}

	match, err := helpers.CheckPasswordHash(password, pwd)

	if !match {
		fmt.Println("Hash and Password doesn't match")
		return false, err
	}

	return true, nil

}
