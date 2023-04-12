package logs

import (
	"os"

	"github.com/rs/zerolog"
)

func infoLogger() zerolog.Logger {
	logger := zerolog.New(os.Stderr).With().Timestamp().Stack().Logger()
	return logger
}

var InfoLogger = infoLogger()