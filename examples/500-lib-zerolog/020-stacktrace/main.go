package main

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	err := outer()
	log.Error().Stack().Err(err).Msg("")

	// Output: {"level":"error","stack":[{"func":"inner","line":"20","source":"main.go"},{"func":"middle","line":"24","source":"main.go"},{"func":"outer","line":"32","source":"main.go"},{"func":"main","line":"15","source":"main.go"},{"func":"main","line":"250","source":"proc.go"},{"func":"goexit","line":"1172","source":"asm_arm64.s"}],"error":"seems we have an error here","time":1678183948}
}

func inner() error {
	return errors.New("seems we have an error here")
}

func middle() error {
	err := inner()
	if err != nil {
		return err
	}
	return nil
}

func outer() error {
	err := middle()
	if err != nil {
		return err
	}
	return nil
}
