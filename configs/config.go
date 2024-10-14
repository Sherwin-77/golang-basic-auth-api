package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Name string
	Key  string
	Port string
	PSQL struct {
		Host     string
		Port     string
		User     string
		Password string
		Database string
	}
}

func GetConfig() *Config {
	config := &Config{
		Name: os.Getenv("APP_NAME"),
		Key:  os.Getenv("APP_KEY"),
		Port: os.Getenv("APP_PORT"),
	}

	config.PSQL.Host = os.Getenv("PSQL_HOST")
	config.PSQL.Port = os.Getenv("PSQL_PORT")
	config.PSQL.User = os.Getenv("PSQL_USER")
	config.PSQL.Password = os.Getenv("PSQL_PASSWORD")
	config.PSQL.Database = os.Getenv("PSQL_DATABASE")

	return config
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	return GetConfig()
}
