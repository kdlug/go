// Go does not have pass-by-reference semantics because Go does not have reference variables.
// A map value is a pointer to a runtime.hmap structure.
// func makemap(t *maptype, hint int64, h *hmap, bucket unsafe.Pointer) *hmap
// https://dave.cheney.net/2017/04/30/if-a-map-isnt-a-reference-variable-what-is-it
package main

import (
	"fmt"
)

// creates a map
func fn(m map[int]int) {
	m = make(map[int]int) // m is not nil,
	// fmt.Println(m == nil) // false
}

func main() {
	var m map[int]int // m is nill

	fn(m)

	// If the map m was a C++ style reference variable, the m declared in main and the m declared in fn would occupy the same storage location in memory.
	// But, because the assignment to m inside fn has no effect on the value of m in main, we can see that maps are not reference variables.
	fmt.Println(m == nil) // true
}
