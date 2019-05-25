// We don't have to use structs to implement interfaces
/// We can use any type of custom type
package main

import "fmt"

func main() {
	myInt := IntCounter(0)
	// IntCounter is also type of incrementer because it's implementing Inrement() method
	// so we can add a pointer
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (i *IntCounter) Increment() int {
	*i++
	return int(*i)
}
