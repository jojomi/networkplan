package main

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.With().Caller().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})

	if err := getCmdRoot().Execute(); err != nil {
		log.Fatal().Err(err).Msg("")
		os.Exit(1)
	}
}
