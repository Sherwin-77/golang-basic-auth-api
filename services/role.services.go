package services

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sherwin-77/golang-basic-auth-api/db"
	"github.com/sherwin-77/golang-basic-auth-api/models"
)

type RoleService struct {
	BaseService
}

func (s *RoleService) GetRoles() []models.Role {
	DB := db.DB

	var roles []models.Role
	DB.Find(&roles)

	return roles
}

func (s *RoleService) GetRoleByID(id string) models.Role {
	DB := db.DB

	var role models.Role

	roleID := s.GetUUID(id)
	if err := DB.Where("id = ?", roleID).First(&role).Error; err != nil {
		panic(echo.NewHTTPError(http.StatusNotFound, "Role not found"))
	}

	return role
}

func (s *RoleService) CreateRole(role models.Role) models.Role {
	DB := db.DB

	if err := DB.Create(&role).Error; err != nil {
		panic(err)
	}

	return role
}

func (s *RoleService) UpdateRole(role *models.Role) {
	DB := db.DB

	if err := DB.Save(role).Error; err != nil {
		panic(err)
	}
}

func (s *RoleService) DeleteRole(role *models.Role) {
	DB := db.DB

	if err := DB.Delete(role).Error; err != nil {
		panic(err)
	}
}
