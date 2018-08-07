package main

import (
	"fmt"
)

// create a newtype of deck
type deck []string

// method iterates over the deck and print cards
func (d deck) print() {
	// iterate over cards
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// returns a value of type deck
// without a receiver because it creates a new deck
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Cubes"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// We don't use indices, so instead of index we replace use _
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards

}

// returns hands cards and remaining cards
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
