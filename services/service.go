package services

import (
	"github.com/google/uuid"
	"github.com/sherwin-77/golang-basic-auth-api/db"
	"github.com/sherwin-77/golang-basic-auth-api/models"
)

type BaseService struct {
}

func (s *BaseService) GetUUID(user interface{}) uuid.UUID {
	var userId uuid.UUID

	switch u := user.(type) {
	case models.BaseModel:
		userId = u.ID
	case string:
		userId = uuid.MustParse(u)
	case uuid.UUID:
		userId = u
	default:
		userId = uuid.Nil
	}

	return userId
}

func (s *BaseService) PreloadModel(preloads []string, model interface{}) {
	DB := db.DB

	for _, preload := range preloads {
		DB = DB.Preload(preload)
	}

	DB.First(model)
}
