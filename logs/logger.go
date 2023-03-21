package logs

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Logger() zerolog.Logger {

	logFile, err := os.OpenFile(
        "./logs/log-file.log",
        os.O_RDWR|os.O_CREATE|os.O_APPEND,
        0664,
    )
	if err != nil {
		log.Panic().Msgf("couldn't open file; %s", err)
	}

	logger := zerolog.New(logFile).With().Timestamp().Caller().Logger()
	return logger
}