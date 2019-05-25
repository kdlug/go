// Map variables copies 
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

	pp := phonePrefixes
	delete(pp, "Greece")

	fmt.Println(pp, phonePrefixes) // Greece key was removed from both maps!
}
