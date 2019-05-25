package main

import "fmt"

func main() {
	phonePrefixes := map[string]int{
		"Greece":      30,
		"Netherlands": 31,
		"Belgium":     32,
		"France":      33,
		"Spain":       34,
	}

	fmt.Println(phonePrefixes)
	pp := phonePrefixes
	delete(pp, "Spain")
	// Spain key were removed from both maps
	fmt.Println(pp, phonePrefixes)
}
