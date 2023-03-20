package logs

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func Logger() zerolog.Logger {
	logFile, err := os.OpenFile(
        "./logs/log-file.log",
        os.O_APPEND|os.O_CREATE|os.O_WRONLY,
        0664,
    )
	if err != nil {
		log.Fatal().
		Err(err).
		Msg("Couldn't open file")
	}
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	logger := zerolog.New(logFile).With().Timestamp().Caller().Logger()
	return logger
}