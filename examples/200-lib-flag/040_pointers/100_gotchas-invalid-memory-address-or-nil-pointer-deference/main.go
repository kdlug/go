// panic: runtime error: invalid memory address or nil pointer dereference
package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func main() {
	/*
		panic: runtime error: invalid memory address or nil pointer dereference
		var p *Point
		fmt.Println(p.Abs())
	*/

	// 1)
	// point := Point{X: 0, Y: 0}
	// p := &point

	// 2)
	// var p *Point // The uninitialized pointer p in the main function is nil
	// p = &Point{0, 0}

	// 3)
	// var p *Point = new(Point) // we have to create a Point first
	// fmt.Println(p.Abs())      // 0

	// Since methods with pointer receivers take either a value or a pointer, you could also skip the pointer
	var p Point          // has zero value Point{X:0, Y:0}
	fmt.Println(p.Abs()) // 0
}
