package main

/*
Creating URL flag which is type of url.URL so we can easily get scheme, host, path of that value
fmt.Printf(`{scheme: %q, host: %q, path: %q}`, u.Scheme, u.Host, u.Path)

Usage
  -url
		url to parse

go run *.go -url="https://golang.org/pkg/flag/""
*/

import (
	"flag"
	"fmt"
	"net/url"
)

// 1. We need to define custom type which implements the Value interface
//
// type Value interface {
//     String() string
//     Set(string) error
// }
// type customValue string
//
// 2. String() // for setting custom value
// func (c *customValue) String() string { return string(*c) }
//
// 3. Set() // for setting a value to type f.ex.
// func (c *customValue) Set(val string) error {
//	*c = customValue(val)
//	return nil
//}

// URLValue defines a custom type which will implements Value interface
// and holds URL
// We can use the predefined url.URL struct in net/url package
type URLValue struct {
	URL *url.URL // a pointer to url.URL
}

// We you should define String() method which converts customValue to string
// It's called when no parameter is provided in the command line
// It defines a custom value (which is a string)
func (v *URLValue) String() string {
	//return string((*u).URL)
	// we have to check if v.URL is empty and return "" empty string in that case
	// because we will get error: invalid memory address or nil pointer dereference
	if v.URL == nil {
		v.URL = &url.URL{} // initialize
	}

	return v.URL.String() // url implements Stringer interface so we can use it
}

// Function set is called if flag in command line is defined
func (v *URLValue) Set(s string) error {
	// we have to parse first rawurl into a URL structure
	// using url.Parse(string) function
	u, err := url.Parse(s)
	if err != nil {
		return err
	}

	v.URL = u

	return nil
}

var urlFlag URLValue

func init() {
	// Flag package provides an easy way to extend by using flag.Var function.
	// It creates a flag for any type that obey flag.Value interface
	// func Var(value Value, name string, usage string)
	// default value will be value.String()
	flag.Var(&urlFlag, "url", "URL to parse")
}

func main() {
	flag.Parse()

	fmt.Printf("Type: %T, Value: %v\n", urlFlag, urlFlag)
	fmt.Printf(`{scheme: %q, host: %q, path: %q}`, urlFlag.URL.Scheme, urlFlag.URL.Host, urlFlag.URL.Path)
}
