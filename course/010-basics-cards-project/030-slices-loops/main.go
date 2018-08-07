// slices and loops
// array - fixed lenght
// slices - can grow or shrink

package main

import (
	"fmt"
)

func main() {
	// slice of strings
	cards := []string{"Five of Diamonds", newCard()}
	// add a new element to a slice
	cards = append(cards, "Six of Spades") // append doesn't modify existing slice, it creates a new one
	fmt.Println(cards)

	// iterate over cards
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

// we have to define type of returned value
func newCard() string {
	return "Ace of spades"
}
