// Maps are like hash map or associative array (in other languages)
package main

import "fmt"

func main() {
	phonePrefixes := make(map[string]int)
	phonePrefixes = map[string]int{ // we are mapping one key type to one value type
		"Greece":      30,
		"Netherlands": 31,
		"Belgium":     32,
		"France":      33,
		"Spain":       34, // if we use key pairs notation in new line, comma is required here
	}

	fmt.Println(phonePrefixes)
}
