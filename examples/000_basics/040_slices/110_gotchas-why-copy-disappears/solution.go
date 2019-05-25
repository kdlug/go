// Why does the dst disappear?
package main

import "fmt"

func main() {
	var src, dst []int
	src = []int{1, 2, 3}

	// The number of elements copied by the copy function is the minimum of len(dst) and len(src).
	// To make a full copy, you must allocate a big enough destination slice.
	dst = make([]int, len(src))

	n := copy(dst, src) // The return value of the copy function is the number of elements copied
	fmt.Println("dst:", dst, "(copied", n, "numbers)")
}
