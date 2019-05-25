package main

import "fmt"

func main() {
	var divide func(float64, float64) (float64, error)

	// function should be defined before we're trying to execute it
	// 	d, err := divide(5.0, 0.0) // this will cause an error
	divide = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0, fmt.Errorf("Division by 0 is forbidden")
		}
		return a / b, nil
	}

	// divide() returns 2 values
	d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}
