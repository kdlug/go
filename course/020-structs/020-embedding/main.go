// Structs
// Embedding
package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo // embedding
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Bean",
		contact: contactInfo{
			email:   "jim.bean@gmail.com",
			zipCode: 12345,
		},
	}

	fmt.Printf("%+v", jim)
}
