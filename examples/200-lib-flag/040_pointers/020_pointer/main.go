package main

import "fmt"

func main() {
	x := 0

	fmt.Println("In main function before calling passByReference function value [x]:", x)
	fmt.Println("In main function before calling passByReference function address [&x]:", &x)

	passByValue(&x) // & means address of int

	// x value is changed because it was passed to function by reference
	fmt.Println("In main function after calling passByReference function value [x]:", x)
	// address is the same
	fmt.Println("In main function after calling passByReference function address [&x]:", x)
}

func passByValue(pointer *int) { // the value is a pointer to an int, we are passing an address, it's still passing by value (value of an address)
	// address is the same
	fmt.Println("In passByReference function, before change [pointer]:", pointer)   // prints address of a pointer
	fmt.Println("In passByReference function, before change [*pointer]:", *pointer) // prints value under that address

	// changing value at given address
	*pointer = 100

	// address is the same
	fmt.Println("In passByReference function, after change [pointer]:", pointer)
	fmt.Println("In passByReference function, after change [*pointer]:", *pointer)
}
