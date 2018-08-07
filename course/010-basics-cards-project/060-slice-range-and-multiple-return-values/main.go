// Range of slice
// sliceVar[startIndexIncluding:upToNotIncluding]
// fruits := []{"apple", "banana", "kiwi"}
// fruits[0:2] // apple, banana
// is equivalent to fruits[:2]
// add deck function which returns 2 values of type deck
package main

func main() {
	// use deck type, which is slice of strings
	cards := newDeck()
	cards.print()

	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()

}
