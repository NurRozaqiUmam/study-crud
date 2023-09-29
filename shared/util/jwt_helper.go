package util

import (
	"echo_crud/config"
	"echo_crud/pkg/domain"
	"encoding/base64"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtClaims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	jwt.StandardClaims
}

func CreateJwtToken(user domain.User) (string, error) {
	conf := config.GetConfig()
	claimss := JwtClaims{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Address:  user.Address,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claimss)
	token, err := rawToken.SignedString([]byte(conf.SignKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func EncryptPassword(password string) string {
	return base64.StdEncoding.EncodeToString([]byte(password))
}
func DecryptPassword(password string) string {
	data, err := base64.StdEncoding.DecodeString(password)
	if err != nil {
		panic(err)
	}
	return string(data)
}
