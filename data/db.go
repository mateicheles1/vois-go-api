package data

import (
	"gogin-api/config"
	"gogin-api/initializers"
	"gogin-api/logs"

	"gorm.io/gorm"
)

func ReturnDB() *gorm.DB {
	config, err := config.LoadConfig("./config/config.json")

	if err != nil {
		logs.ErrorLogger.Fatal().Msgf("Could not load config: %s", err)
		return nil
	}

	db, err := initializers.OpenDB(config)
	if err != nil {
		logs.ErrorLogger.Fatal().Msgf("Could not connect to db: %s", err)
		return nil
	}

	return db
}
