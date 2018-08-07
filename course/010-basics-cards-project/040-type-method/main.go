// use type deck and create a method to print cards
// go run main.go deck.go
package main

func main() {
	// use deck type, which is slice of strings
	cards := deck{"Five of Diamonds", newCard()}
	// add a new element to a slice
	cards = append(cards, "Six of Spades") // append doesn't modify existing slice, it creates a new one

	cards.print()
}

// we have to define type of returned value
func newCard() string {
	return "Ace of spades"
}
