// create deck_test.go (with _test.go suffix) file
// run: go test
package main

func main() {
	// use deck type, which is slice of strings
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
