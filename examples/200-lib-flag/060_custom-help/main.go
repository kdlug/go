package main

import (
	"flag"
	"fmt"
)

var n int
var b bool
var s string

func init() {
	flag.IntVar(&n, "number", 0, "Prints number") // flag name, default option, flag description
	flag.StringVar(&s, "string", "", "Prints string")
	flag.BoolVar(&b, "bool", false, "Prints bool")
}

func main() {
	flag.Args()

	fmt.Printf("n(%T), s(%T), b(%T)\n", n, s, b)
}
