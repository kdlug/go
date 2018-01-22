/*
Closures
A special type of anonymous function that references variables declared outside of the function itself
It's similar to how a regular function can reference global variables
In other words a closure is created when a function returns another function (anonymous) that access the variables in the sope of it's first function
https://play.golang.org/p/mPdN3ucAb3M
*/
package main

import (
	"fmt"
)

func counter() func() int { // it returns an anonymous function
	i := 0 // we want to have an access to this variable outside

	return func() int { // closure
		i++
		return i // variable i is being closed in a function; inner function is holding variables defined in outer function
	}
}

func main() {
	counter()
	// fmt.Println(i) // it still doesn't work

	// assign function to a variable, it returns the value of i
	firstCounter := counter()
	fmt.Println("First call:", firstCounter())  // First call: 1
	fmt.Println("Second call:", firstCounter()) // Second call: 2
	fmt.Println("Third call:", firstCounter())  // Third call: 3

	// assign to another variable and run again - counter will restart
	secondCounter := counter()
	fmt.Println("First call:", secondCounter()) // First call: 1
}

