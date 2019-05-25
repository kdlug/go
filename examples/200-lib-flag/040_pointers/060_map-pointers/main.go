// maps contais a pointer to the underlying data
// if you copy a map, you copy a pointer
package main

import "fmt"

func main() {
	// map type
	a := map[string]string{
		"foo": "bar",
		"baz": "buz",
	}

	b := a // copying map a to b (the value is a pointer!)

	fmt.Println(a, b)
	a["foo"] = "oops"
	fmt.Println(a, b) // map[baz:buz foo:oops] map[baz:buz foo:oops]
}
