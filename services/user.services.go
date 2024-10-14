package services

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sherwin-77/golang-basic-auth-api/db"
	"github.com/sherwin-77/golang-basic-auth-api/models"
)

type UserService struct {
	BaseService
}

func (s *UserService) GetUsers() []models.User {
	DB := db.DB

	var users []models.User
	DB.Find(&users)

	return users
}

func (s *UserService) GetUserByID(id string) models.User {
	DB := db.DB

	var user models.User

	userID := s.GetUUID(id)
	if err := DB.Where("id = ?", userID).First(&user).Error; err != nil {
		panic(echo.NewHTTPError(http.StatusNotFound, "User not found"))
	}

	return user
}

func (s *UserService) GetUserByEmail(email string) (models.User, error) {
	DB := db.DB

	var user models.User

	err := DB.Where("email = ?", email).First(&user).Error

	return user, err
}

func (s *UserService) CreateUser(user models.User) models.User {
	DB := db.DB

	if err := DB.Create(&user).Error; err != nil {
		panic(err)
	}

	return user
}

func (s *UserService) UpdateUser(user *models.User) {
	DB := db.DB

	if err := DB.Save(user).Error; err != nil {
		panic(err)
	}
}

func (s *UserService) DeleteUser(user *models.User) {
	DB := db.DB

	if err := DB.Delete(user).Error; err != nil {
		panic(err)
	}
}
