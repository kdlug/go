// Why does the dst disappear?
package main

import "fmt"

func main() {
	var src, dst []int
	src = []int{1, 2, 3}
	copy(dst, src)           // Copy elements to dst from src.
	fmt.Println("dst:", dst) // dst: []
}
