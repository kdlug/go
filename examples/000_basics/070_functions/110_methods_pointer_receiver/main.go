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
// (g greeter) is a pointer reciver
func (g *greeter) greet() { // we can use any type not only struct
	// we have now a pointer to g
	// implicitly dereferencing pointers for struct
	// so we don't have to add * before (but we can if we want)
	fmt.Println(g.greeting, g.name)
	// we can change the value of field
	g.name = "Alex"
}
