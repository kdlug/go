package main

import "fmt"

// & gives an address of the value
// * give a value of an address
// *int means the pointer to an int (this is a type different than int)

func main() {
	// Without pointers
	var a int = 10   // or in short: a := 10
	fmt.Println(a)   // getting value
	fmt.Println(&a)  // getting address in memory when the value is stored
	fmt.Println(*&a) // will print value :-)

	c := a          // copy a value from a to c and store in a new memory address
	fmt.Println(c)  // getting value
	fmt.Println(&c) // getting address in memory when the value is stored

	// changing vale of c
	c = 7
	fmt.Println(a, c) // a value wasn't changed

	// checking types
	fmt.Printf("%T\n", a)  // int
	fmt.Printf("%T\n", &a) // *int (type of address is a pointer to an int)

	// Pointers
	// var b int = &a // will cause an error, because those are 2 different types (int and *int)
	var b *int = &a // b is a pointer to integer and it points to a
	// b := &a
	fmt.Println(b)        // address
	fmt.Println(*b)       // value (* deferencing an address)
	fmt.Printf("%T\n", b) // *int

	fmt.Println(&*b) // will print address

	*b = 43        // we're changing value at address
	fmt.Println(a) // 43 (b is a pointer to a, we changed a value at original location)

}
