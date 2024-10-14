package services

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sherwin-77/golang-basic-auth-api/db"
	"github.com/sherwin-77/golang-basic-auth-api/models"
)

type TodoService struct {
	BaseService
}

func (s *TodoService) GetTodos(user interface{}) []models.Todo {
	DB := db.DB

	var todos []models.Todo
	if userID := s.GetUUID(user); userID != uuid.Nil {
		DB.Where("user_id = ?", userID.String()).Find(&todos)
	}

	return todos
}

func (s *TodoService) GetTodoByID(id string) models.Todo {
	DB := db.DB

	var todo models.Todo

	todoID := s.GetUUID(id)
	if err := DB.Where("id = ?", todoID.String()).First(&todo).Error; err != nil {
		panic(echo.NewHTTPError(http.StatusNotFound, "Todo not found"))
	}

	return todo
}

func (s *TodoService) CreateTodo(todo models.Todo) models.Todo {
	DB := db.DB

	if err := DB.Create(&todo).Error; err != nil {
		panic(err)
	}

	return todo
}

func (s *TodoService) UpdateTodo(todo *models.Todo) {
	DB := db.DB

	if err := DB.Save(todo).Error; err != nil {
		panic(err)
	}
}

func (s *TodoService) DeleteTodo(todo *models.Todo) {
	DB := db.DB

	if err := DB.Delete(todo).Error; err != nil {
		panic(err)
	}
}
