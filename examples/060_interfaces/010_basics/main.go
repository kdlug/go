// An interface type consists of a set of method signatures. A variable of interface type can hold any value that implements these methods.
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x MyStringer

	x = Temp(24)
	fmt.Println(x.String()) // 24 °C

	// Actually, *Temp also implements MyStringer, since the method set of a pointer type *T is the set of all methods with receiver *T or T.
	t := Temp(24)
	x = &t
	fmt.Println(x.String()) // 24 °C

	// 	But Point{0, 0} does not implement MyStringer (String method has pointer receiver)
	x = &Point{1, 2}
	fmt.Println(x.String()) // (1,2)
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
