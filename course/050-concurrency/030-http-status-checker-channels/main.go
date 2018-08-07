// We use channel to make sure that child go routines are executed before our program exits
// Channel is for communication between different gorutines f.ex main and child
// channels are typed

// channel <- 5 - send value 5 into a channel
// number <- channel - wait for a value to be sent into a channel, when we get one, assign the value to the number
// pass to fanction fmt.Println(<-channel)

// In this program we're only receiving one communication over the channel
// http://google.com is up!
// Yep, it's up!

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

	// channel with values of type string
	// will be used to communicate between main function and checkLink function
	c := make(chan string)

	for _, link := range links {
		// child goroutine
		go checkLink(link, c)
	}

	// Receive msg from channel
	// blocking line of code
	// main routine is sleep until it gets a value from channel
	// but only the first goroutine
	// main routine wakes up, print value end exits
	fmt.Println(<-c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link + " is down!")
		return
	}
	fmt.Println(link + " is up!")
	// send value into a channel
	c <- "Yep, it's up!"
}
