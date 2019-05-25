package shape

import "math"

// Rectangle has a method called Area that returns a float64 so it satisfies the Shape interface
// Circle has a method called Area that returns a float64 so it satisfies the Shape interface
// string does not have such a method, so it doesn't satisfy the interface
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

func Area(r Rectangle) float64 {
	return r.Width + r.Height
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
