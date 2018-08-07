// a map is a collection of key-value pairs
// all keys and values must be the same type
// a map is a refference type (struct is a value type)
// declaring a map
package main

import (
	"fmt"
)

func main() {
	// #1
	colors := map[string]string{
		"red":   "#ff0000", // comma is mandatory here
		"green": "#00ff00",
	}

	fmt.Println(colors)

	// #2
	var fruits map[string]string // a zero value will be assigned (an empty map)
	fmt.Println(fruits)

	// #3 Equivalent to #2
	cars := make(map[string]string)
	fmt.Println(cars)

	// Adding elements to a map
	cars["honda"] = "green"
	fmt.Println(cars)

	// ?? panic: assignment to entry in nil map
	// fruits["kiwi"] = "green"
	// fmt.Println(fruits)

	// delete a key/value pair
	delete(cars, "honda")
	fmt.Println(cars)
}
