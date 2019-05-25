// Using returning variable name can be confusing
package main

import "fmt"

func main() {
	// sum returns a pointer so we have to dereference it
	s := sum(1, 2, 3, 4, 5)
	fmt.Println("The sum is", s)
}

func sum(values ...int) (result int) { // we specified a variable name and type which will be returned
	// we get a slice of int :)
	fmt.Printf("%T %v\n", values, values)
	// we can omit initialization because it was done automatically
	// result := 0
	for _, v := range values {
		result += v
	}

	return // we don't have to specify name of variable, we specified it in function declaration
}
