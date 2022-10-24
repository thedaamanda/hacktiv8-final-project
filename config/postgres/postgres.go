package postgres

import (
	"os"
	"project/src/model"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitGorm() *gorm.DB {

	connection := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(connection))

	if err != nil {
		log.Error().Msgf("cant connect to database %s", err)
	}

	db.AutoMigrate(&model.User{}, &model.Photo{}, &model.Comment{}, &model.SocialMedia{})

	return db
}
