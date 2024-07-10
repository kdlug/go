package main

import "fmt"

// printer displays information
type printer interface {
	print()
}

// user defines a user in the program
type user struct {
	name string
}

// print displays the user's name
func (u user) print() {
	fmt.Printf("User Name: %s\n", u.name)
}

func main() {
	u := user{"Pablo"}

	// Add the values and pointers
	// to the slce of printer interface values
	entities := []printer{
		// Store a copy of user value in the interface value
		u,
		// Store a copy of the address og the user value in the interface
		&u,
	}

	// Change the name field on the user value
	u.name = "Pablo_CHG"
	// Iterate over the slice of entities and call
	// print against the copied interface value
	for _, e := range entities {
		e.print()
	}
}
