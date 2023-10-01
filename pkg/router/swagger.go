package router

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewSwaggerRouter(e *echo.Echo, g *echo.Group, db *sql.DB) {
	// swagger route
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
