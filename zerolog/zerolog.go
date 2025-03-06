package main

import (
	"flag"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// set global log levels
	debug := flag.Bool("debug", false, "sets log level to debug")
	flag.Parse()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Print("hello world")
	// Output: {"time":1516134303,"level":"debug","message":"hello world"}

	// add context to message
	log.Debug().
		Str("Scale", "833 cents").
		Float64("Interval", 833.09).
		Msg("Fibonacci is everywhere")

	log.Debug().
		Str("Name", "Tom").
		Send()

	// set default level
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Debug().Msg("This message appears when log level set to Debug")
	log.Info().Msg("This message appears when log level set to Debug or Info")

	if e := log.Debug(); e.Enabled() {
		// Compute log output only if enabled.
		value := "bar"
		e.Str("foo", value).Msg("some debug message")
	}

	log.Info().Msg("all done")
}
