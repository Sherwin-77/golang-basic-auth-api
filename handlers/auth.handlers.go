package handlers

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/sherwin-77/golang-basic-auth-api/configs"
	"github.com/sherwin-77/golang-basic-auth-api/models"
	authrequests "github.com/sherwin-77/golang-basic-auth-api/requests/auth"
	"github.com/sherwin-77/golang-basic-auth-api/resources"
	"github.com/sherwin-77/golang-basic-auth-api/services"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	services.UserService
}

func (h *AuthHandler) Login(ctx echo.Context) error {
	authRequest := authrequests.LoginRequest{}
	config := configs.GetConfig()

	if err := ctx.Bind(&authRequest); err != nil {
		panic(err)
	}

	if err := ctx.Validate(&authRequest); err != nil {
		panic(err)
	}

	user, err := h.UserService.GetUserByEmail(authRequest.Email)

	if err != nil {
		user.Password = "$2a$10$pRe6SEQi6edG0bEYzAaMF.S1oszSANbZORukCi7j3QFku5jC1frFW"
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authRequest.Password)); err != nil {
		panic(echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password"))
	}

	var key = []byte(config.Key)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, models.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    config.Name,
			ExpiresAt: time.Now().Add(time.Hour * 8).Unix(),
		},
		UserID: user.ID.String(),
	})

	tokenString, err := token.SignedString(key)
	if err != nil {
		panic(err)
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"token": tokenString,
	})
}

func (h *AuthHandler) Register(ctx echo.Context) error {
	authRequest := authrequests.RegisterRequest{}
	resource := resources.UserResource{}

	if err := ctx.Bind(&authRequest); err != nil {
		panic(err)
	}

	if err := ctx.Validate(&authRequest); err != nil {
		panic(err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(authRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	user := h.UserService.CreateUser(models.User{
		Username: authRequest.Username,
		Email:    authRequest.Email,
		Password: string(hashedPassword),
	})

	return ctx.JSON(http.StatusCreated, resource.Make(user))
}
