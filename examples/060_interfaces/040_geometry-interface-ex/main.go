package main

import (
	"fmt"
	"math"
)

func main() {
	// Interface has two types. A static type of interface is interface itself - geometry in our case.
	// A interface do not have a static value, rather it points to a dynamic value.
	var g geometry
	fmt.Printf("Value of g is: %v\n", g) // nil
	fmt.Printf("Type of g is: %T\n ", g) // nil (dynamic type of interface)

	r := rectangle{
		width:  3,
		height: 4,
	}

	c := circle{
		radius: 5,
	}

	measure(r)
	measure(c)

	// because rectangle implements geometry interface, we can assign it to g
	g = r
	// g = int(9) // annot use int(9) (type int) as type geometry in assignment: int does not implement geometry (missing area method)
	fmt.Printf("Value of g is: %v\n", g) //  {3 4}
	fmt.Printf("Type of g is: %T\n ", g) // main.rectangle (dynamic type of interface)

}

type geometry interface {
	area() float64
	perim() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
