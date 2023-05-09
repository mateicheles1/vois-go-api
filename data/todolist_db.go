package data

import (
	"gogin-api/config"
	"gogin-api/logs"
	"gogin-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB(config *config.Config) *gorm.DB {

	db, err := gorm.Open(postgres.Open(config.DB.ConnectionString()), &gorm.Config{})

	if err != nil {
		logs.ErrorLogger.Fatal().Msgf("Could not connect to DB due to: %s", err)
	}

	if err = db.AutoMigrate(&models.ToDoList{}, &models.ToDo{}); err != nil {
		logs.ErrorLogger.Fatal().Msgf("Failed to migrate to db: %s", err)
	}

	return db
}
