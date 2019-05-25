// Why doesn’t this code compile?
// cannot assign to s[0]
package main

import "fmt"

func main() {
	// Go strings are immutable and behave like read-only byte slices (with a few extra properties).
	// s := "hello"

	// To update the data, use a rune slice instead.
	// If the string only contains ASCII characters, you could also use a byte slice.
	buffer := []rune("hello")
	buffer[0] = 'H'
	s := string(buffer)
	fmt.Println(s) // "Hello"
}
