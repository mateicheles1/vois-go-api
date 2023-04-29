package logs

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func errorLogger() (zerolog.Logger, *os.File) {
	logFile, err := os.OpenFile(
		"./logs/log-file.log",
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0777,
	)
	if err != nil {
		log.Panic().Msgf("Could not open file due to: %s", err.Error())
	}

	logger := zerolog.New(logFile).With().Timestamp().Caller().Logger()
	return logger, logFile
}

var ErrorLogger, LogFile = errorLogger()
