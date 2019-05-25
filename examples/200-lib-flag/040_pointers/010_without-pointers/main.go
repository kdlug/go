package main

import "fmt"

func main() {
	fmt.Println("Hi there")

	x := 0
	passByValue(x)

	fmt.Println("In main function after calling passByValue function:", x)
}

func passByValue(value int) {
	// original value
	fmt.Println("Within function, before change:", value)
	// changing value
	value = 100
	fmt.Println("Within function, after change:", value)
}
