package main

import "fmt"

// create a newtype of deck
type deck []string

// method iterates over the deck and print cards
func (d deck) print() {
	// iterate over cards
	for i, card := range d {
		fmt.Println(i, card)
	}
}
