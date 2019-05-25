package main

import "fmt"

func main() {
	// we can pass as many int arguments as we want
	sum(1, 2, 3, 4, 5)

	// the first argument have to be string
	sum2("The sum is", 1, 2, 3, 4, 5)
}

// all arguments have the same type
func sum(values ...int) {
	// we get a slice of int :)
	fmt.Printf("%T %v\n", values, values)
	result := 0
	for _, v := range values {
		result += v
	}

	fmt.Println("The sum is", result)
}

// First argument has a different type
// Variadic argument can be only one and placed at the end
func sum2(msg string, values ...int) {
	// we get a slice of int :)
	fmt.Printf("%T %v\n", values, values)
	result := 0
	for _, v := range values {
		result += v
	}

	fmt.Println(msg, result)
}
