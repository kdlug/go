// strings.EqualFold() Function in Golang reports whether s and t, interpreted as UTF-8 strings,
// are equal under Unicode case-folding, which is a more general form of case-insensitivity.
package main

import (
	"fmt"
	"strings"
)

func main() {

	// case insensitive comparing and returns true.
	fmt.Println(strings.EqualFold("Geeks", "geeks"))

	// case insensitive comparing and returns true.
	fmt.Println(strings.EqualFold("COMPUTERSCIENCE",
		"computerscience"))

	// OUTPUT:
	// true
	// true
}
