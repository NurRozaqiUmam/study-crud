package util

import "github.com/labstack/echo/v4"

// tipe struct yang digunakan untuk menghasilkan respons JSON
type JsonReponse struct {
	Code    int         `json:"code"`    // Kode status HTTP dalam respons
	Message string      `json:"message"` // pesan yang menjelaskan respons
	Data    interface{} `json:"data"`    // data yang akan disertakan dalam respons
}

// SetResponse, fungsi yang membantu menghasilkan respons JSON dengan format yang sama
func SetResponse(c echo.Context, statusCode int, message string, data interface{}) error {
	return c.JSON(statusCode, JsonReponse{
		Code:    statusCode,
		Message: message,
		Data:    data,
	})
}
