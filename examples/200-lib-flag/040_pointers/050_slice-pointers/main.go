// slices contais a pointer to the underlying data
// if you copy a slice, you copy a pointer
package main

import "fmt"

func main() {
	// array type
	a := [3]int{1, 2, 3}
	b := a            // copying array a to b
	fmt.Println(a, b) // [1 2 3] [1 2 3]
	a[1] = 99         // only a array will be changed because b is a copy
	fmt.Println(a, b) // [1 99 3] [1 2 3]

	// slice type (we removed only length from array type)
	// a slice is a projection of an underlying array
	// it doesn't contain the data itself
	// it contains a pointer to the first element of the underlying array
	s := []int{1, 2, 3}
	fmt.Println(s)
	c := s            // copying array s to c but s now is a pointer!
	fmt.Println(s, c) // [1 2 3] [1 2 3]
	s[1] = 99         // only a array will be changed because b is a copy
	fmt.Println(s, c) // [1 99 3] [1 99 3]
}
