package main

import (
	"fmt"
)

func main() {
	// long form
	// declaring a variable and assigning a value
	// var name type
	var card string = "Ace of spades"

	// short form
	// Alternative way, it will be also type string, go will figure out of the type base on the value
	// card := "Ace of spades"
	// it can't be used outside the function body
	// We can initialize a variable outside of a function, we just can't assign a value to it.
	fmt.Println(card)
}
