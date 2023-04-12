package initializers

import (
	"gogin-api/logs"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	if DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		logs.ErrorLogger.Fatal().Msgf("Failed to connect to database due to: %s", err.Error())
	}
}