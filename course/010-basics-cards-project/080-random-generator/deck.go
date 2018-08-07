package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

func (d deck) toString() string {
	// convert back the type
	// join and add a separator
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		// #1 option - log the error return a new deck
		// #2 option - log the error and quit
		fmt.Println("Error", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)

}

// randomize
// pseudo random generator rand.Intn depends on random seed value (source of randomness)
func (d deck) shuffle() {
	// create a new Rand object
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := range d {
		np := r.Intn(len(d) - 1)
		d[i], d[np] = d[np], d[i]
	}
}
