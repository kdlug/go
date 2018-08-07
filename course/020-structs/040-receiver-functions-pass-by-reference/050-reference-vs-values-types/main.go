// Slices (a reference type)

// Everytime we make a slice we have to types: a slice and underlying array

// ss := []string{"banana", "apple", "kiwi"}
// updateSlice(s []string)
//
// RAM
//                  Address | Variable | Value
//                  -------------------------------------------
//                | 0000    |          |                        |
//                | 0001    |   ss     | [len][cap][ptr]--------|---
//                                                                  |
//                | 0002    |          | []string{"banana",...} |<-- (underlying array)
//                                                                  |
//                | 0003    |    s     | [len][cap][ptr]--------|---
//
//
// value types | reference types
// -----------------------------
//    int      |     slices
//   float     |      maps
//   string    |     channels
//    bool     |     pointers
//   struct    |    functions
package main

import (
	"fmt"
)

func main() {
	ss := []string{"banana", "apple", "kiwi"}
	updateSlice(ss)
	fmt.Println(ss) // [orange apple kiwi]

}

// slice data structure is copied but it's still pointing to the original array
func updateSlice(s []string) {
	s[0] = "orange"
}
