package repository

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Instance *gorm.DB
	err      error
)

func Connect(connectionString string) {
	Instance, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Error().Strack().Err(err).Msg("Error connecting to database.")
		panic("Cannot connect to DB")
	}
	log.Info().Print("Connected to Database...")
}

func Migrate() {
	Instance.AutoMigrate(&entities.Product{})
	log.Info().Print("Database Migration Completed...")
}
