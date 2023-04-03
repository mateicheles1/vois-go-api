package initializers

import (
	"gogin-api/logs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := "host=ruby.db.elephantsql.com user=dshfydzd password=YVKi7Ce2-a3kdDuTAk7s1ZIah74lO_q- dbname=dshfydzd port=5432"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logs.Logger.Fatal().Msgf("Failed to connect to database due to: %s", err.Error())
	}
}