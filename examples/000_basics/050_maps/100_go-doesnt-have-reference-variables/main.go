// Go does not have reference variables
// It is not possible to create a Go program where two variables share the same storage location in memory.
// It is possible to create two variables whose contents point to the same storage location, but that is not the same thing as two variables who share the same storage location.
package main

import "fmt"

func main() {
	var a int
	var b, c = &a, &a // we are assigning a value here which is an address to variable a
	// b and c stores the same value
	fmt.Println(b, c) // 0xc00001e140 0xc00001e140
	fmt.Println(&a)   // which is address of v variable
	// but addresses of b and c variables are different
	fmt.Println(&b, &c) // 0x1040c108 0x1040c110
}
