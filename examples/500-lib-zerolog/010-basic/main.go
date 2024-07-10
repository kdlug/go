package main

import (
	"errors"
	"flag"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	isDebug = flag.Bool("debug", false, "sets log level to debug")
)

func init() {
	flag.Parse()

	// it's good to set up logs in init

	// UNIX Time is faster and smaller than most timestamps
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// change level for DEBUG mode
	if *isDebug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Debug().Msg("debug mode is enabled")
	}
	// Output: {"level":"debug","time":1678183138,"message":"debug mode is enabled"}
}
func main() {
	//logger := log.Logger
	// By default log writes to os.Stderr
	log.Print("hello world") // The default log level for log.Print is debug so its the same as: Logger.Debug().Msg("hello world)
	// Output: {"level":"debug","time":1678183138,"message":"hello world"}

	// leveled log
	// The chain must have either the Msg or Msgf method call.
	// If you forget to add either of these, the log will not occur and there is no compile time error
	log.Warn().Msg("some warning")
	// Output: {"level":"warn","time":1678183138,"message":"some warning"}

	// formatted info
	log.Info().Msgf("info: %s", "zerolog")
	// Output: {"level":"info","time":1678183138,"message":"info: zerolog"}

	// zerolog allows data to be added to log messages in the form of key:value pairs.
	// The data added to the message adds "context" about the log event
	log.Debug().
		Str("Scale", "833 cents").
		Float64("Interval", 833.09).
		Msg("Fibonacci is everywhere")

	log.Debug().
		Str("Name", "Tom").
		Send()
	// Output: {"level":"debug","Scale":"833 cents","Interval":833.09,"time":1562212768,"message":"Fibonacci is everywhere"}
	// Output: {"level":"debug","Name":"Tom","time":1562212768}

	if e := log.Debug(); e.Enabled() {
		// Compute log output only if enabled.
		value := "bar"
		e.Str("foo", value).Msg("some debug message")
	}
	// Output: {"level":"debug","foo":"bar","time":1678183498,"message":"some debug message"}

	//logger.UpdateContext(func(c zerolog.Context) zerolog.Context {
	//	return c.
	//		Str("address", address).
	//		Str("http-server", "API")
	//})

	// logging without lever or error msg
	log.Log().
		Str("foo", "bar").
		Msg("")
	// Output: {"foo":"bar","time":1678183581}

	// errors loging
	var err error

	err = errors.New("seems we have an error here")
	log.Error().Err(err).Msg("")
	// Outout: {"level":"error","error":"seems we have an error here","time":1678183697}

	// err := errors.New("panic error example message")
	if err != nil {
		log.Panic().Err(err).Msg("initialize config") // exits app
	}

}
