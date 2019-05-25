// Equality
// Two interface values are equal
// - if they have equal concrete values and identical dynamic types,
// - or if both are nil.
// A value t of interface type T and a value x of non-interface type X are equal if
// - it’s concrete value is equal to x
// -and t’s dynamic type is identical to X.
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x MyStringer
	fmt.Printf("T: %T, V: %v, x == nil: %v\n", x, x, x == nil) // true (because value and conrete type is nil)

	// The concrete value of x equals nil, but its dynamic type is *Point, which is not nil.
	x = (*Point)(nil)
	fmt.Printf("T: %T, V: %v, x == nil: %v\n", x, x, x == nil) // false because an interface value has a conrete type
}

type MyStringer interface {
	String() string
}

type Temp int

// A type implements an interface by implementing its methods. No explicit declaration is required.
func (t Temp) String() string {
	return strconv.Itoa(int(t)) + " °C"
}

type Point struct {
	x, y int
}

func (p *Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}
