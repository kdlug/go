// Go automatically dereferences struct field pointers
package main

import "fmt"

type myStruct struct {
	foo int
}

func main() {
	var ms *myStruct // ms is a pointer to myStruct
	fmt.Println(ms)  // initialization value of pointer is <nil>
	// If we are acceptiong pointer as arguments it is the best practice to check the value if it's not nil

	ms = &myStruct{foo: 42} // gets address of myStruct
	fmt.Println(ms)         // &{42} - it shows that is a pointer and automatically prints value
	fmt.Println(*ms)        // {42}
	// fmt.Println(ms.foo) // automatically dereferencing, explained below

	// we can also use new operator (it's not the best practice)
	ms = new(myStruct)
	fmt.Println(ms) // &{0}
	// in order to set a foo field to struct we have to dereference a pointer first using *
	(*ms).foo = 17 // () are needed because * has lower precedence than . operator and we only want to dereference ms variable not ms.foo
	fmt.Println((*ms).foo)

	// Because (*ms) syntax is ugly, Go doing dereferencing automatically
	// the followng code works the same as above
	ms = new(myStruct)
	ms.foo = 17
	fmt.Println(ms.foo)

}
