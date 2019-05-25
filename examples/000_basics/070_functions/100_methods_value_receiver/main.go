package main

import "fmt"

type greeter struct {
	greeting string
	name     string
}

func main() {
	g := greeter{
		greeting: "Hey",
		name:     "Joe",
	}

	g.greet() // method invocation

	fmt.Println("Name is", g.name)
}

// function which are executed in a known context
// (g greeter) part is a value reciver
func (g greeter) greet() { // we can use any type not only struct
	// we have copy of a greeting object so wehave access to its fields
	fmt.Println(g.greeting, g.name)
	// we can change the value of field
	// but it will be changed locally
	g.name = "Alex"
}
