package main

import (
	"log"

	"github.com/sherwin-77/golang-basic-auth-api/configs"
	"github.com/sherwin-77/golang-basic-auth-api/db"
	"github.com/sherwin-77/golang-basic-auth-api/models"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	configs.LoadConfig()
	db.InitDB()

	password, err := bcrypt.GenerateFromPassword([]byte("secret"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	DB := db.DB

	users := []models.User{
		{
			Username: "admin",
			Email:    "admin@example.com",
			Password: string(password),
		},
		{
			Username: "editor",
			Email:    "editor@example.com",
			Password: string(password),
		},
		{
			Username: "user",
			Email:    "user@example.com",
			Password: string(password),
		},
	}

	DB.Create(&users)

	log.Println("Users table seeded successfully")
}
