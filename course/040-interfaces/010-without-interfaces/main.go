// every value has a type
// every function has to specify types of it's arguments
// but it doesn't mean that we have to rewrite every function with the same logic to accomodate different types
// interfaces help us to reuse code
//
// We have 2 functions getGreeting() and printGreeting() for 2 languages
// Functions with the same name will be overwritten, so below code won't work
// Basically implementationn of getGreeting() will be different for each language
// But printGreeting() will be exactly the same
// We can use interfaces to fix this problem
package main

import (
	"fmt"
)

type englishBot struct{}

type spanishBot struct{}

func main() {
	eb := englishBot{}
	// sb := spanishBot{}

	printGreeting(eb)
	// printGreeting(sb)
}

func (eb englishBot) getGreeting() string {
	return "Hello"
}

// func (sb spanishBot) getGreeting() string {
// 	return "Hola"
// }

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

// func printGreeting(sb spanishBot) {
// 	fmt.Println(spanishBot.getGreeting())
// }
