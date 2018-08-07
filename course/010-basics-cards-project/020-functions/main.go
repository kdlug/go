// create a function newCard() which return a card

package main

import (
	"fmt"
)

func main() {
	// call a method and assign a return value to a card variable
	card := newCard()

	fmt.Println(card)

}

// we have to define type of returned value
func newCard() string {
	return "Ace of spades"
}
