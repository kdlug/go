// In most languages the only way is causing an exception
package main

import "fmt"

func main() {
	// divide() returns 2 values
	d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}

// function returns 2 values
func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0, fmt.Errorf("Division by 0 is forbidden")
	}
	return a / b, nil
}
