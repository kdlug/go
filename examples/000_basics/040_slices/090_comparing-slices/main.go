// Invalid operation s1 == s1 (slice can only be compared to nil)

package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1

	/*
		if s1 == s2 { // invalid operation: s1 == s2 (slice can only be compared to nil)
			fmt.Println("Slices are equals")
		}
	*/

	// To compare two slices we have to use our own method
	if Equal(s1, s2) {
		fmt.Println("Slices are equals")
	}

}

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
