// Structs
// Receiver functions
// Passing by reference

// &variable - memory address of the value this variable is pointing at
// *variable - the value this memory address is pointing at
// *type - type description, a pointer to a given type (it may be confusing with *variable)
//
// transformation
// address -> value   *address
// value   -> address  value
//
// jimPointer := jim
//
// RAM
//                  Address | Variable | Value
//                  -------------------------------------------
//                | 0000    |          |                        |
// jimPointer --> | 0001    | jim      | person {"Jim", "Bean"} |
//                | 0002    |          |                        |
//                | 0003    |          |                        |
//                | 0004    |          |                        |

package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo // embedding - we can only use a type, it's equivalent to contact contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Bean",
		contactInfo: contactInfo{ // fieldname is the same as type
			email:   "jim.bean@gmail.com",
			zipCode: 12345,
		},
	}

	// pointer to person
	jimPointer := &jim // memory address to the jim variable (reference, pointer)
	print(jimPointer)  // 0xc420057f40

	jimPointer.updateName("Jimmy") // it will modify jim variable, because method receiver is passed by reference
	jim.print()                    // {firstName:Jim lastName:Bean contactInfo:{email:jim.bean@gmail.com zipCode:12345}}

	// we can use a shortcut with receivers (we dont't need to specify pointer)
	// if we specify a pointer receiver func (pointerToPerson *person)
	jim.updateName("George")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// person type is passed by reference *pointer
// *person means that we expect pointer to persons
func (pointerToPerson *person) updateName(name string) {
	// pointerToPerson is memory address
	// *pointerToPerson is a value stored in that address
	(*pointerToPerson).firstName = name
}
