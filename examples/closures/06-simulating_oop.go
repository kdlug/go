/*
Closures can simulate Object Oriented Programming / Encapsulation

https://play.golang.org/p/h8CFauMTzrW
*/
package main

import (
	"fmt"
)

func counter() (func() int, func() int, func() int) { // returns 3 functions
	i := 0 // variable, outer function scope

	increment := func() int { // anonymous function
		i += 1
		return i
	}

	decrement := func() int { // anonymous function
		i -= 1
		return i
	}

	reset := func() int { // anonymous function
		i = 0
		return i
	}

	return increment, decrement, reset
}

func main() {
	inc, dec, reset := counter()
	fmt.Println("Inc:", inc())
	fmt.Println("Inc:", inc())
	fmt.Println("Inc:", inc())
	fmt.Println("Dec:", dec())
	fmt.Println("Reset:", reset())

}

