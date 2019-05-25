package main

/*
Usage
  -deltaT value
		comma separated list of intervals

go run *.go -deltaT=10s,20s,30s,40s

go run *.go -deltaT=10s -deltaT=20s -delataT=30s
invalid value "20s" for flag -deltaT: interval flag already set
*/

import (
	"errors"
	"flag"
	"fmt"
	"strings"
	"time"
)

// user defined type - a slice of durations
type interval []time.Duration

// String is the method to format the flag's value, part of the flag.Value interface
func (i *interval) String() string {
	return fmt.Sprint(*i)
}

// Set is the method to set the flag value, part of the flag.Value interface.
// Set's argument is a string to be parsed to set the flag.
// It's a comma-separated list, so we split it.
func (i *interval) Set(value string) error {
	// permit to have only one flag
	if len(*i) > 0 {
		return errors.New("interval flag already set")
	}

	for _, dt := range strings.Split(value, ",") {
		duration, err := time.ParseDuration(dt)

		if err != nil {
			return err
		}

		*i = append(*i, duration)
	}

	return nil
}

// Define a flag to accumulate durations.
var intervalFlag interval

func init() {
	flag.Var(&intervalFlag, "deltaT", "comma separated list of intervals")
}

func main() {
	// parsing flag
	flag.Parse()
	fmt.Printf("%v, intervalFlag (%T)", intervalFlag, intervalFlag)
}
