// A type assertion provides access to an interface’s concrete value
// A type assertion doesn’t really convert an interface to another data type,
// but it provides access to an interface’s concrete value, which is typically what you want.
package main

import "fmt"

func main() {
	// Type assertion x.(T) asserts that the concrete value stored in the x is type of T and that x is not null
	// If T is not an interface, it asserts that the dynamic type of x is identical to T
	// if T is an interface, it asserts that the dynamic type of x implements T
	var x interface{} = "foo"

	var s string = x.(string)
	fmt.Println(s) // foo

	s, ok := x.(string)
	fmt.Println(s, ok) // foo, true

	n, ok := x.(int)
	fmt.Println(n, ok) // 0, false

	// n = x.(int) // ILLEGAL - interface conversion: interface {} is string, not int
}
