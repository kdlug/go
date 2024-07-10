package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	var ErrClient = errors.New("error occurred in Client")

	//err := fmt.Errorf("%w: %s", ErrClient, errors.New("Client Error"))
	err := errors.Wrap(ErrClient, errors.New("Client Error").Error())

	u := errors.Unwrap(err)
	fmt.Println(u)
	fmt.Println(err)

	// doubleErr := errors.Wrap(err, "double error")

	switch errors.Cause(err) { // errors.Unwrap(err)
	case nil:
		fmt.Println("nil")
	case ErrClient:
		fmt.Println("ErrClient")
	default:
		fmt.Println("Default")
	}

	u2 := errors.Unwrap(errors.New("Client Error"))
	fmt.Println(u2)
}
