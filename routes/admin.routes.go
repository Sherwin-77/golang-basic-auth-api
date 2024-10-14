package routes

import (
	"github.com/labstack/echo/v4"
	adminhandlers "github.com/sherwin-77/golang-basic-auth-api/handlers/admin"
	"github.com/sherwin-77/golang-basic-auth-api/routes/middlewares"
)

func RegisterAdminRoutes(e *echo.Group) {
	e.Use(middlewares.Authenticated)
	e.Use(middlewares.AuthLevel(3))

	userHandler := adminhandlers.UserHandler{}

	e.GET("", func(ctx echo.Context) error {
		return ctx.JSON(200, map[string]interface{}{
			"message": "Hello Admin!",
			"data":    nil,
		})
	})

	e.GET("/users", userHandler.GetUsers)
	e.GET("/users/:id", userHandler.GetUserByID)
	e.POST("/users", userHandler.CreateUser)
	e.PATCH("/users/:id", userHandler.UpdateUser)
	e.DELETE("/users/:id", userHandler.DeleteUser)
}
