package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/wralith/mirket/src/userservice/app/config"
)

// TODO: Additional levels
func InitLogger(c *config.LoggerConfig) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// Human readable form instead of json if pretty
	if c.Pretty {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	if c.Level == "debug" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}
