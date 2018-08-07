// Delay
// Function Literal but with bug
// variable are referenced
// so only last link will be checked
// Output:
// http://google.com is up!
// http://facebook.com is up!
// http://amazon.com is up!
// http://amazon.com is up!
// http://amazon.com is up!
/// http://amazon.com is up!
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
		go checkLink(link, c)
	}

	// infinite loop
	// for {
	// 	go checkLink(<-c, c)
	// }
	// alternative syntax
	for l := range c {
		// this is bad solution because it sleeps main routine
		// go checkLink(l, c)
		// time.Sleep(time.Second)

		go func() {
			time.Sleep(time.Second * 5)
			checkLink(l, c)
		}()
	}
}

func checkLink(link string, c chan string) {
	// this is bad solution as well because it will wait before call a request
	// time.Sleep(time.Second)
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link + " is down!")
		return
	}
	fmt.Println(link + " is up!")

	c <- link
}
