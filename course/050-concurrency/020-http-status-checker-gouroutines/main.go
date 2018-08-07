// concurrency - we can have multiple threads executing code (scheduler).
// If one thread blocks another one is picked up and worked on
// multiple goroutines may be running on a single CPU
// Parralelism - multiple threads are executed at the exact same time. Requires multiple CPUs
package main

import (
	"fmt"
	"net/http"
)

// main routine
// if main routine finish it exits program, it doesn't care of child routines
func main() {
	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://facebook.com",
	}

	for _, link := range links {
		// child goroutine
		go checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link + " is down!")
		return
	}
	fmt.Println(link + " is up!")
}
