// Delay
// Function Literal
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://facebook.com",
	}

	c := make(chan string)

	for _, link := range links {
		// child goroutine
		fmt.Println("#1 loop only runs goroutines")
		go checkLink(link, c)
	}

	for l := range c {
		fmt.Println("#2 loop runs and receives goroutines")
		go func(l string, c chan string) {
			time.Sleep(time.Second * 5)
			checkLink(l, c)
		}(l, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link + " is down!")
		return
	}
	fmt.Println(link + " is up!")

	c <- link
}
