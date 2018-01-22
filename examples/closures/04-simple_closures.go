/*
Closures
https://play.golang.org/p/pSGcfr5pj0h
*/
package main

import (
	"fmt"
)

func simpleClosure() func() int {
	i := 0
	fmt.Println("External i:", i) // this is run once only when function is assigned to a variable (f := simpleClosure())
	
	return func() int { // this is run each time when f() is called, i is persistent
		i++
		fmt.Println("Closure i:", i)
		return i
	}
}



func main() {
	fmt.Println("Example 1")
	// runs only first function, without closure
	simpleClosure() // External i: 0
	// runs also closure, but i data isn't persisted
	simpleClosure()() // External i: 0, Closure i: 1
	simpleClosure()() // External i: 0, Closure i: 1
	
	fmt.Println("\nExample 2")
	// assign to a variable, i data is persisted
	f := simpleClosure() // External i: 0
	f() // Closure i: 1
	f() // Closure i: 2

}
