package data

import (
	"gogin-api/config"
	"gogin-api/logs"
	"gogin-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() *gorm.DB {
	config, err := config.LoadConfig("./config/config.json")

	if err != nil {
		logs.ErrorLogger.Error().Msgf("Could not load config due to: %s", err)
		return nil
	}

	db, err := gorm.Open(postgres.Open(config.ConnectionString()), &gorm.Config{})

	if err != nil {
		logs.ErrorLogger.Fatal().Msgf("Could not connect to DB due to: %s", err)
		return nil
	}

	if err = db.AutoMigrate(&models.ToDoList{}, &models.ToDo{}); err != nil {
		logs.ErrorLogger.Fatal().Msgf("Failed to migrate db: %s", err)
		return nil
	}

	return db
}
