package main

import "fmt"

// type Actor struct {
// 	name string
// }
func main() {
	// using field names are optional before values is best practice
	actor := struct{ name string }{name: "Robert De Niro"} // the second {} is initializer
	fmt.Println(actor)
	fmt.Println(actor.name)

	// only data are passed (copy of data)
	anotherActor := actor
	anotherActor.name = "John Depp"

	fmt.Println(actor, anotherActor)
}
