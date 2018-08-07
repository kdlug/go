// Pointers
// We gets a pointer to name called namePointer and prints out the memory address of the pointer itself!
// Then in the function, we take the pointer that was passed as an argument and print out the address of that pointer too.
// The memory address printed by both Println calls will be different because everything in go is passed by value:-)
package main

import "fmt"

func main() {
	name := "bill"

	namePointer := &name

	fmt.Println(&namePointer) // 0x1040c130
	printPointer(namePointer) // 0x1040c140
}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
}
