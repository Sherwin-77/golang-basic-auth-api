package adminhandlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sherwin-77/golang-basic-auth-api/models"
	adminrequests "github.com/sherwin-77/golang-basic-auth-api/requests/admin"
	"github.com/sherwin-77/golang-basic-auth-api/resources"
	"github.com/sherwin-77/golang-basic-auth-api/services"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	services.UserService
	services.RoleService
}

func (h *UserHandler) GetUsers(ctx echo.Context) error {
	resource := resources.UserIndexResource{}

	users := h.UserService.GetUsers()

	return ctx.JSON(http.StatusOK, resource.Collections(users))
}

func (h *UserHandler) GetUserByID(ctx echo.Context) error {
	resource := resources.UserResource{}

	id := ctx.Param("id")
	user := h.UserService.GetUserByID(id)

	h.UserService.PreloadModel([]string{"Roles", "Todos"}, &user)

	return ctx.JSON(http.StatusOK, resource.Make(user))
}

func (h *UserHandler) CreateUser(ctx echo.Context) error {
	resource := resources.UserResource{}
	userRequest := adminrequests.UserRequest{}

	if err := ctx.Bind(&userRequest); err != nil {
		panic(err)
	}

	if err := ctx.Validate(&userRequest); err != nil {
		panic(err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	user := h.UserService.CreateUser(models.User{
		Username: userRequest.Username,
		Email:    userRequest.Email,
		Password: string(hashedPassword),
	})

	return ctx.JSON(http.StatusCreated, resource.Make(user))
}

func (h *UserHandler) UpdateUser(ctx echo.Context) error {
	resource := resources.UserResource{}
	userRequest := adminrequests.UpdateUserRequest{}

	if err := ctx.Bind(&userRequest); err != nil {
		panic(err)
	}

	if err := ctx.Validate(&userRequest); err != nil {
		panic(err)
	}

	user := h.UserService.GetUserByID(userRequest.ID)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	user.Username = userRequest.Username
	user.Email = userRequest.Email
	user.Password = string(hashedPassword)

	h.UserService.UpdateUser(&user)

	return ctx.JSON(http.StatusOK, resource.Make(user))
}

func (h *UserHandler) DeleteUser(ctx echo.Context) error {
	userID := ctx.Param("id")

	user := h.UserService.GetUserByID(userID)

	h.UserService.DeleteUser(&user)

	return ctx.NoContent(http.StatusNoContent)
}
