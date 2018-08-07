// Structs
// Receiver functions

// Passing by value
//
// jim := person{}
// jim.updateName("Jimmy") -> func (p person) updateName(name string)
//
// RAM
//
// Address | Variable | Value
// -------------------------------------------
// 0000    |          |
// 0001    | jim      | person {"Jim", "Bean"}
// 0002    |          |
// 0003    | p        | person {"Jim", "Bean"} // copy of a value
// 0004    |          |

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

	// call receiver methods
	jim.updateName("jimmy") // it won't modify jim variable, because method receiver is passed by value, not by refference (pointer)
	jim.print()             // {firstName:Jim lastName:Bean contactInfo:{email:jim.bean@gmail.com zipCode:12345}}
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// person type is passed by value, so the function can't modify any of it's properties
func (p person) updateName(name string) {
	p.firstName = name
}
