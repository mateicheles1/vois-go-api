package initializers

import (
	"gogin-api/config"
	"gogin-api/logs"
	"gogin-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDB(config *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.ConnectionString()), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(&models.ToDoList{}, &models.ToDo{}); err != nil {
		logs.ErrorLogger.Fatal().Msgf("Failed to migrate: %s", err)
	}
	return db, nil
}
