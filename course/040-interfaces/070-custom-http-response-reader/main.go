package main

import (
	"fmt"
	"io"
	"os"
	"reflect"

	"./checker"
)

func main() {
	urls := []string{
		"http://google.pl",
		"http://somenotexistingurl.dede/code",
		"http://localhost:8080/api",
		"http://myenglishlab.com",
		"http://google.pl/asdasd",
	}

	results := checker.Check(urls)

	// bs := make([]byte, 9999)

	// results.Read(bs)
	// fmt.Println(string(bs))

	io.Copy(os.Stdout, results)

	// Check if Results implements io.Reader interface
	readerType := reflect.TypeOf((*io.Reader)(nil)).Elem()

	resultsType := reflect.TypeOf((*checker.Results)(nil))
	fmt.Println(resultsType.Implements(readerType))
}
