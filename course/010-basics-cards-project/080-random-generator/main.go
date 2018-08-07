// shuffle cards using random number generator
package main

func main() {
	// use deck type, which is slice of strings
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
