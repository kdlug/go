// interfaces are implicit, we don't have to link with interface keyword etc.
// Add bot interface with getGreeting() function
// Uncomment func (sb spanishBot) getGreeting() string
// Modify printGreeting and replace an argument to bot
package main

import (
	"fmt"
)

// every type with method called getGreeting() is a member of bot type
type bot interface {
	getGreeting() string // the same name for english and spanish bot, but different implementation
}

type englishBot struct{}

type spanishBot struct{}

func main() {
	eb := englishBot{} // is also of type bot (it implements bot interface)
	sb := spanishBot{} // is also of type bot (it implements bot interface)

	printGreeting(eb)
	printGreeting(sb)
}

// englishBot implements bot interface because it has a function call getGreeting() which returns a string
// english bot is automatically type of bot
func (eb englishBot) getGreeting() string {
	return "Hello"
}

// spanishBot implements bot interface because it has a function call getGreeting() which returns a string
// spanish bot is automatically type of bot
func (sb spanishBot) getGreeting() string {
	return "Hola"
}

// a common function for all bots
// because we added bot interface, we can use bot type as an argument
// if bot is english bot, printGreeting function will display Hello
func printGreeting(b bot) {
	fmt.Println(b.getGreeting()) // function is elastic, depends on
}
