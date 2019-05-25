package main

import (
	"flag"
	"fmt"
)

var gopherType string

func init() {
	const (
		defaultGopher = "pocket"
		usage         = "the variety of gopher"
	)

	// Two flags sharing a variable, so we can have a shorthand
	// The order of initialization is undefined, so make sure both use the
	// same default value. They must be set up with an init function.
	flag.StringVar(&gopherType, "gopher_type", defaultGopher, usage)
	flag.StringVar(&gopherType, "g", defaultGopher, usage+" (shorthand)")
}

func main() {
	// parsing flag should be in main
	flag.Parse()
	fmt.Printf("%v, gopherType (%T)", gopherType, gopherType)
}
