// https://github.com/golang/go/wiki/TableDrivenTests
//  tip with table driven tests: use t.Run and to name the test cases
// go test -v
package shape

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	// We are creating an "anonymous struct", areaTests.
	// We are declaring a slice of structs by using []struct with two fields, the shape and the want.
	// Then we fill the slice with cases.
	areaTests := []struct { // anonymous struct
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{5.0, 10.0}, hasArea: 50.0},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{2.0, 10.0}, hasArea: 10.0},
	}

	// using tt.name from the case to use it as the `t.Run` test name
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.hasArea) // The %#v format string will print out our struct with the values in its field
			}
		})

	}
}
