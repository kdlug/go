package main

import "fmt"

func main() {
	// we can pass as many int arguments as we want

	s := sum(1, 2, 3, 4, 5)
	fmt.Println("The sum is", s)
}

func sum(values ...int) int { // we have to specify a type of returning value
	// we get a slice of int :)
	fmt.Printf("%T %v\n", values, values)
	result := 0
	for _, v := range values {
		result += v
	}

	return result
}
