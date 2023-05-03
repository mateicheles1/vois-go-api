package main

import (
	"gogin-api/config"
	"gogin-api/initializers"
	"gogin-api/logs"
	"gogin-api/routes"

	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	config, err := config.LoadConfig("./config/config.json")

	if err != nil {
		logs.ErrorLogger.Error().Msgf("Could not load config: %s", err)
		return
	}

	db, err := initializers.OpenDB(config)
	if err != nil {
		logs.ErrorLogger.Fatal().Msgf("Could not connect to db: %s", err)
	}
	DB = db
}

func main() {
	routes.SetupRoutes()
}
