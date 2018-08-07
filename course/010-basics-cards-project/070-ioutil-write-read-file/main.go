// func writeFile(filename string, data []byte, perm os.FileMode) error
// our challenge is to extract data - convert deck to []byte
// type conversion
// []byte("Hi there") - converts a string to slice of bytes
// deck -> []byte
// deck -> []string -> join to a single string -> []byte
// func readFile(filename string) ([]byte, error)
package main

import "fmt"

func main() {
	// use deck type, which is slice of strings
	cards := newDeck()
	cards.print()

	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
	fmt.Println(remainingCards.toString())
	cards.saveToFile("my_cards")
	cff := newDeckFromFile("my_cards")
	cff.print()
}
