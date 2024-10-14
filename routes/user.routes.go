package routes

import (
	"github.com/labstack/echo/v4"
	userhandlers "github.com/sherwin-77/golang-basic-auth-api/handlers/user"
	"github.com/sherwin-77/golang-basic-auth-api/routes/middlewares"
)

func RegisterUserRoutes(e *echo.Group) {
	todoHandler := new(userhandlers.TodoHandler)

	e.Use(middlewares.Authenticated)
	e.Use(middlewares.AuthLevel(2))

	e.GET("/todos", todoHandler.GetTodos)
	e.GET("/todos/:id", todoHandler.GetTodoByID, middlewares.ValidateUUID)
	e.POST("/todos", todoHandler.CreateTodo)
	e.PATCH("/todos/:id", todoHandler.UpdateTodo, middlewares.ValidateUUID)
	e.DELETE("/todos/:id", todoHandler.DeleteTodo, middlewares.ValidateUUID)
}
