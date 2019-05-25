// Why does the dst disappear?
package main

import "fmt"

func main() {
	var src, dst []int
	src = []int{1, 2, 3}

	// Using append
	// You could also use the append function to make a copy by appending to a nil slice.
	// Note that the capacity of the slice allocated by append may be a bit larger than len(src).
	dst = append(dst, src...)

	// Why we can't just use this ??? (check)
	// dst = src

	fmt.Println("dst:", dst)
}
