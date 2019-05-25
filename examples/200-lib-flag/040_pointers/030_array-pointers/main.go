package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := &a[0]
	c := &a[1]
	fmt.Printf("%v %p %p\n", a, b, c) // %p prints the value of the pointer
	// notice that c address is 8 higher than b, because we're working on array of integers and integer is 8 bytes long

	// Go language doesn't have a pointer arithmetic (we can hack it and import unsafe package)
	// Fe.x want point to head on the array (like in other lang)
	// c := &a[1] - 4 // it won't work because of mismatching types *int and int
}
