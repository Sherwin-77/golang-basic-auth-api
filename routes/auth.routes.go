package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/sherwin-77/golang-basic-auth-api/handlers"
)

func RegisterAuthRoutes(e *echo.Group) {
	authHandler := handlers.AuthHandler{}

	e.POST("/login", authHandler.Login)
	e.POST("/register", authHandler.Register)
}
