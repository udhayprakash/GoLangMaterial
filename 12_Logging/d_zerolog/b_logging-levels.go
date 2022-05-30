package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// Setting the log level
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	log.Print("print log")

	log.Trace().Msg("trace log")
	log.Debug().Msg("debug log")
	log.Info().Msg("info log")
	log.Warn().Msg("warn log")
	log.Error().Msg("error log")
	log.Fatal().Msg("fatal log")
	log.Panic().Msg("panic log")

	log.Print("Last Statement") 

}
