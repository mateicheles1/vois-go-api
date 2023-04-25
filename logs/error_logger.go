package logs

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var LogFile *os.File
var Err error

func errorLogger() zerolog.Logger {

	LogFile, Err = os.OpenFile(
		"./logs/log-file.log",
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0777,
	)
	if Err != nil {
		log.Panic().Msgf("Could not open file due to: %s", Err.Error())
	}

	logger := zerolog.New(LogFile).With().Timestamp().Caller().Logger()
	return logger
}

var ErrorLogger = errorLogger()
