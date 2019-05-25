// Upper case function is public (accessible from other packages), lowercase function is private
package main

import "fmt"

// function main() in pkg main
// entry point of our application
func main() { // opening curlu brace should be at the same line as func keyword
	// calling a function with passing an argument
	sayMessage("Hi there!")

	for i := 0; i < 5; i++ {
		sayBye("Bye", i)
	}

	sayGreeting("Ciao", "Bambino")

	// we are passing a value (value is copied)
	name := "Joe"
	// which means that original variable doesn't change
	sayGreeting("Hey", name)
	fmt.Println("Name from main function: " + name) // Joe
}

// function with parameters
func sayMessage(msg string) {
	fmt.Println(msg)
}

// function with 2 parameters
func sayBye(msg string, index int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is", index)
}

// function with 2 parameters, but all args are in the same type
// so we can specify type at the end
func sayGreeting(greeting, name string) {
	fmt.Println(greeting, name)
	// changing local variable, won't change a variable outside a function
	name = "Jack"
	fmt.Println("Name from sayGreeting function: " + name)
}
