package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/wralith/mirket/src/api-gateway/app/config"
)

func InitLogger(c *config.LoggerConfig) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if c.Pretty {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	if c.Level == "debug" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}
