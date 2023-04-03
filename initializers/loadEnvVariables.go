package initializers

import (
	"gogin-api/logs"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	if err := godotenv.Load(); err != nil {
		logs.Logger.Fatal().Msgf("Error loading env file due to: %s", err.Error())
	}
}