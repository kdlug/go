// https://github.com/golang/go/wiki/TableDrivenTests
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
		shape Shape
		want  float64
	}{
		{Rectangle{5.0, 10.0}, 50.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %.2f want %.2f", got, tt.want)
		}
	}
}
