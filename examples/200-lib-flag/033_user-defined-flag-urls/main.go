package main

/*
Creating URL flag which is type of []url.URL so we can easily get scheme, host, path of that value
fmt.Printf(`{scheme: %q, host: %q, path: %q}`, u.Scheme, u.Host, u.Path)

Usage
  -url
		url to parse

go run *.go -url="https://golang.org/pkg/flag/" -url="https://golang.org/"
*/

import (
	"errors"
	"flag"
	"fmt"
	"net/url"
	"strings"
)

type UrlsValue struct {
	urls []*url.URL // a pointer to url.URL
}

func (v *UrlsValue) String() string {
	// converting []*url.URL to []string
	s := []string{}
	for _, value := range v.urls {
		s = append(s, value.String())
	}

	return strings.Join(s, " ")
}

// Set function is called if flag in command line is defined
func (v *UrlsValue) Set(s string) error {
	// permit to have only one flag
	if len(v.urls) > 0 {
		return errors.New("url flag already set")
	}

	// when s=url1, url2, url3
	urls := strings.Split(s, ",")

	for _, value := range urls {
		parsedURL, err := url.Parse(value)
		if err != nil {
			return err
		}
		v.urls = append(v.urls, parsedURL)
	}

	return nil
}

var urlsFlag UrlsValue

func init() {
	flag.Var(&urlsFlag, "url", "comma separated URLs to parse")
}

func main() {
	flag.Parse()

	fmt.Printf("Type: %T, Value: %v\n", urlsFlag, urlsFlag)

	for _, item := range urlsFlag.urls {
		fmt.Printf(`{scheme: %q, host: %q, path: %q}`, item.Scheme, item.Host, item.Path)
	}

}
