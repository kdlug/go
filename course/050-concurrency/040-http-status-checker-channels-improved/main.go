package main

import (
	"fmt"
	"net/http"
)

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

	// we have to read all values from the channel
	// it's blocking operation
	// number of receiver values should be equal number of sent values
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
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
