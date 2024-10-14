package userhandlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sherwin-77/golang-basic-auth-api/models"
	userrequest "github.com/sherwin-77/golang-basic-auth-api/requests/user"
	"github.com/sherwin-77/golang-basic-auth-api/resources"
	"github.com/sherwin-77/golang-basic-auth-api/services"
)

type TodoHandler struct {
	services.TodoService
	services.UserService
}

func (h *TodoHandler) GetTodos(ctx echo.Context) error {
	resource := resources.TodoResource{}

	if userID, ok := ctx.Get("user_id").(string); ok {
		todos := h.TodoService.GetTodos(userID)

		return ctx.JSON(http.StatusOK, resource.Collections(todos))
	}

	return echo.NewHTTPError(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
}

func (h *TodoHandler) GetTodoByID(ctx echo.Context) error {
	resource := resources.TodoResource{}

	if userID, ok := ctx.Get("user_id").(string); ok {
		id := ctx.Param("id")
		todo := h.TodoService.GetTodoByID(id)

		if todo.UserID.String() != userID {
			return echo.NewHTTPError(http.StatusForbidden, http.StatusText(http.StatusForbidden))
		}

		return ctx.JSON(http.StatusOK, resource.Make(todo))
	}

	return echo.NewHTTPError(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
}

func (h *TodoHandler) CreateTodo(ctx echo.Context) error {
	resource := resources.TodoResource{}
	todoRequest := userrequest.TodoRequest{}

	if userID, ok := ctx.Get("user_id").(string); ok {
		if err := ctx.Bind(&todoRequest); err != nil {
			panic(err)
		}

		if err := ctx.Validate(&todoRequest); err != nil {
			panic(err)
		}

		todo := h.TodoService.CreateTodo(models.Todo{
			UserID:      uuid.MustParse(userID),
			Title:       todoRequest.Title,
			Description: todoRequest.Description,
			IsCompleted: todoRequest.IsCompleted,
		})

		return ctx.JSON(http.StatusCreated, resource.Make(todo))
	}

	return echo.NewHTTPError(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
}

func (h *TodoHandler) UpdateTodo(ctx echo.Context) error {
	resource := resources.TodoResource{}
	todoRequest := userrequest.UpdateTodoRequest{}

	if userID, ok := ctx.Get("user_id").(string); ok {
		if err := ctx.Bind(&todoRequest); err != nil {
			panic(err)
		}

		if err := ctx.Validate(&todoRequest); err != nil {
			panic(err)
		}

		todo := h.TodoService.GetTodoByID(todoRequest.ID)

		if todo.UserID.String() != userID {
			return echo.NewHTTPError(http.StatusForbidden, http.StatusText(http.StatusForbidden))
		}

		todo.Title = todoRequest.Title
		todo.Description = todoRequest.Description
		todo.IsCompleted = todoRequest.IsCompleted

		h.TodoService.UpdateTodo(&todo)

		return ctx.JSON(http.StatusOK, resource.Make(todo))
	}

	return echo.NewHTTPError(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
}

func (h *TodoHandler) DeleteTodo(ctx echo.Context) error {

	if userID, ok := ctx.Get("user_id").(string); ok {
		id := ctx.Param("id")
		todo := h.TodoService.GetTodoByID(id)

		if todo.UserID.String() != userID {
			return echo.NewHTTPError(http.StatusForbidden, http.StatusText(http.StatusForbidden))
		}

		h.TodoService.DeleteTodo(&todo)

		return ctx.NoContent(http.StatusNoContent)
	}

	return echo.NewHTTPError(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
}
