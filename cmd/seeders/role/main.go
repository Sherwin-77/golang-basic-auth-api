package main

import (
	"fmt"
	"runtime/debug"

	"github.com/sherwin-77/golang-basic-auth-api/configs"
	"github.com/sherwin-77/golang-basic-auth-api/db"
	"github.com/sherwin-77/golang-basic-auth-api/models"
)

func main() {
	configs.LoadConfig()
	db.InitDB()

	DB := db.DB
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Seeding roles table failed. Rolling back transaction")
			fmt.Printf("Error: %s\n", debug.Stack())
			tx.Rollback()
		} else {
			fmt.Println("Seeding roles table successful")
		}
	}()

	roles := []models.Role{
		{
			Name:      "Admin",
			AuthLevel: 3,
		},
		{
			Name:      "Editor",
			AuthLevel: 2,
		},
		{
			Name:      "Viewer",
			AuthLevel: 1,
		},
	}

	tx.Create(&roles)

	var admin models.User
	result := tx.Where("email = ?", "admin@example.com").Limit(1).Find(&admin)
	if result.Error != nil {
		fmt.Println("Admin user not found. Skipping role assignment")
	} else {
		tx.Model(&admin).Association("Roles").Append(&roles[0])
	}

	var editor models.User
	result = tx.Where("email = ?", "editor@example.com").Limit(1).Find(&editor)
	if result.Error != nil {
		fmt.Println("Editor user not found. Skipping role assignment")
	} else {
		tx.Model(&editor).Association("Roles").Append(&roles[1])
	}

	var viewer models.User
	result = tx.Where("email = ?", "user@example.com").Limit(1).Find(&viewer)
	if result.Error != nil {
		fmt.Println("Viewer user not found. Skipping role assignment")
	} else {
		tx.Model(&viewer).Association("Roles").Append(&roles[2])
	}

	tx.Commit()
}
