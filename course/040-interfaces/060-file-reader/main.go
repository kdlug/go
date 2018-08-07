// create a program that reads the content of a text file
// then print its content to a terminal
// the file name should be procided as an argument to the program
// go build -o readfile
package main

import (
	"fmt"
	"io"
	"os"
)

const usage = `
Usage: readfile [filename.txt]
`

func main() {
	// read args
	// first element is a program name
	args := os.Args

	// check if we get exactly 2 arguments
	if len(args) != 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	filename := args[1]

	// check if file exists
	file, err := os.Open(filename) // For read access.
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// copy data from file to stdout
	io.Copy(os.Stdout, file)
}
