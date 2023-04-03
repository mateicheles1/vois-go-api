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
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logs.Logger.Fatal().Msgf("Failed to connect to database due to: %s", err.Error())
	}
}