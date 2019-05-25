package main

/*
Usage
  -list
		comma separated list of strings

go run *.go -list="First String","Second String","Third String"
*/

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

type list []string

func (l *list) String() string {
	return fmt.Sprint(*l)
}

func (l *list) Set(value string) error {
	// permit to have only one flag
	if len(*l) > 0 {
		return errors.New("list flag already set")
	}

	for _, item := range strings.Split(value, ",") {
		*l = append(*l, item)
	}

	return nil
}

// Define a flag
var listFlag list

func init() {
	flag.Var(&listFlag, "list", "comma separated list")
}

func main() {
	// parsing flag
	flag.Parse()
	fmt.Printf("%v, listFlag (%T)", listFlag, listFlag)
}
