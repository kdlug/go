/*
Anonymous function - the same as a regular function, but without a name
It's created dynamically
https://play.golang.org/p/GhwA-ev9fop
*/

package main

import (
	"fmt"
)

// declare regular function (only on package level)
func someFunction() {
	fmt.Println("Regular function")
}

// assign anonymous function to a variable on package level
var anonymous func() = func() {
	fmt.Println("Anonymous function pkg")
}

func main() {
	// call a regular function
	someFunction()

	// declare and call anonymous function
	func() {
		fmt.Println("Anonymous function")
	}() // () is important, it calls a function

	// anonymous function with parameter
	func(msg string) {
		fmt.Println(msg)
	}("Anonymous function with a parameter") // call and pass parameter
	
	// assign an anonymous function to a variable with declared type
	var f1 func() = func() {
		fmt.Println("Anonymous function assigned to a variable f1")
	}
	
	// call an anonymous function
	f1()
	
	// assign an anonymous function to a variable - short declaration
	f2 := func() {
		fmt.Println("Anonymous function assigned to a variable f2")
	}
	// call an anonymous function
	f2()
	
	// call anonymous function assigned to a variable at package level
	anonymous()
	
	// assign a new function to the anonymous variable at runtime (dynamically change)
	anonymous = func() {
		fmt.Println("Anonymous function rewrited")
	}
	// call it
	anonymous()
	
}




