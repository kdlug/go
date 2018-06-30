// flag package allows us to parse arguments and flags
//
// program generates a random number from given min and max
// we can run our program with arguments fex.
// go run 10-parsing-command-line-arguments.go -min 1 -max 10
// to see help: go run 10-parsing-command-line-arguments.go -h
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Define flags
	// 100 is default value if you run without parameters
	maxVal := flag.Int("max", 100, "max value") // int value

	// Parse
	flag.Parse()

	// Generate
	// Default number generator generates always the same sequence of values (deterministic)
	// rand := rand.Intn(*maxVal)
	seed := rand.NewSource(time.Now().UnixNano()) // we need to give it a seed
	r := rand.New(seed)

	// Print
	fmt.Println("Your lucky number is:", r.Intn(*maxVal))

}
