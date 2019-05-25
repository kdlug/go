// The error type is an interface type. An error variable represents any value that can describe itself as a string. Here is the interface's declaration:
// type error interface {
//     Error() string
// }
package main

import (
	"errors"
	"fmt"
)

type errorString struct {
	s string
}

func main() {
	// throw an error
	err := New("Ooops!")

	if err != nil {
		fmt.Println(err)
	}

	// throw an error
	i, err := MagicFunction(-5)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(i)
}

// implementing error interface
func (e *errorString) Error() string {
	return "Custom error: " + e.s
}

// New returns an error that formats as the given text.
func New(text string) error {
	// returns a reference to errorString structure with s: text
	return &errorString{s: text} // it has to be address because: func (e *errorString) Error() string {
}

// MagicFunction returns error when number is negative
func MagicFunction(i int) (int, error) {
	if i < 0 {
		return 0, errors.New("You can't use negative number")
	}

	return i, nil
}
