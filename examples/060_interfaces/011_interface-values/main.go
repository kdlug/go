// An interface value consists of a concrete value and a dynamic type: [Value, Type]
// The zero value of an interface type is nil, which is represented as [nil, nil].
// You can also use type assertions, type switches and reflection to access the dynamic type of an interface value.
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var y Temp
	fmt.Printf("%v %T\n", y, y) // 0 °C main.Temp

	var x MyStringer
	// The zero value of an interface type is nil, which is represented as [nil, nil].
	fmt.Printf("%v %T\n", x, x) // <nil> <nil>

	x = Temp(24)
	fmt.Printf("%v %T\n", x, x) // 24 °C main.Temp

	x = &Point{1, 2}
	fmt.Printf("%v %T\n", x, x) // (1,2) *main.Point

	x = (*Point)(nil)
	fmt.Printf("%v %T\n", x, x) // <nil> *main.Point
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
