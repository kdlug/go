package main

// Structs
// Defining, declaring and updating values

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

func main() {
	// #1 without defining fields
	james := person{"James", "Bond"}

	// #2 with fields (recommended)
	ethan := person{firstName: "Ethan", lastName: "Hunt"}

	fmt.Println(james, ethan)

	// #3 a value of zero will be assigned (empty string), not nil!
	var jack person

	// print
	fmt.Println(jack) // { }
	// a better way of printing
	fmt.Printf("%+v", jack) // {firstName: lastName:}

	// update values
	jack.firstName = "Jack"
	jack.lastName = "Reacher"
	fmt.Printf("%+v", jack)
}
