package routes

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Group) {
	auth := e.Group("/auth")
	RegisterAuthRoutes(auth)

	admin := e.Group("/admin")
	RegisterAdminRoutes(admin)

	user := e.Group("/user")
	RegisterUserRoutes(user)
}
