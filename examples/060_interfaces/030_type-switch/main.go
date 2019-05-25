// interface{} means any type
// we can change our behaviour depends on type
// A type switch performs several type assertions in series and runs the first case with a matching type.
package main

import (
	"fmt"
)

func main() {
	var i interface{} = "test"
	switch i.(type) {
	case nil:
		fmt.Println("x is nil") // here x has type interface{}
	case int:
		fmt.Println("i is integer")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("I dont know what i is")
	}

}
