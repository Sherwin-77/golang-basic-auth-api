package db

import (
	"fmt"

	"github.com/sherwin-77/golang-basic-auth-api/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	config := configs.GetConfig()

	database, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.PSQL.Host, config.PSQL.Port, config.PSQL.User, config.PSQL.Password, config.PSQL.Database)), &gorm.Config{})
	DB = database

	return err
}
