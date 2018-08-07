// Init function
// init function is called by Go when a package is initialized.
// It does not take any arguments and doesn’t return any value.
// You can have multiple init functions in a file or a package.
// Main job of init function is to initialize global variables
// Order of the execution of init function in a file will be according to the order of their appearances.
// that can not be initialized in global context. For example, initialization of an array, slice etcs

package main

import (
	"fmt"
)

var index []int

func init() {
	fmt.Println("#1 init()")

	// initialize slice
	for i := 0; i < 10; i++ {
		index = append(index, i)
	}
}

func init() {
	fmt.Println("#2 init()")
	fmt.Println(index)
}

func main() {
	fmt.Println("main()")
	fmt.Println(index)
}
