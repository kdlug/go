package statements

import (
	"fmt"
	"testing"
)

func TestIfBasic(t *testing.T) {
	r := If(1)

	if r != true {
		t.Errorf("If(1) = %t; want true", r)
	}
}

func TestIfTableDriven(t *testing.T) {
	var tests = []struct {
		i    int
		want bool
	}{
		{1, true},
		{0, false},
		{2, false},
		{100, false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.i)
		t.Run(testname, func(t *testing.T) {
			r := If(tt.i)

			if r != tt.want {
				t.Errorf("got %t, want %t", r, tt.want)
			}
		})
	}
}

func TestSwitchBasic(t *testing.T) {
	r := Switch(1)

	if r != true {
		t.Errorf("Switch(1) = %t", r)
	}
}

func TestSwitchTableDriven(t *testing.T) {
	// we can do also as map; random order
	data := map[int]bool{
		0:   false,
		1:   true,
		2:   false,
		100: false,
	}

	for given, wanted := range data {
		testname := fmt.Sprintf("%d", given)
		r := Switch(given)

		t.Run(testname, func(t *testing.T) {
			if r != wanted {
				t.Errorf("got %t, want %t", r, wanted)
			}
		})
	}
}

func TestSwitchConst(t *testing.T) {
	r := SwitchConst(1)

	if r != true {
		fmt.Errorf("SwitchConst(1) = %t", r)
	}
}

func TestSwitchConstTableDriven(t *testing.T) {
	data := map[int]bool{
		0:   false,
		1:   true,
		2:   false,
		100: false,
	}

	for given, wanted := range data {
		testname := fmt.Sprintf("%d", given)
		r := SwitchConst(given)

		t.Run(testname, func(t *testing.T) {
			if r != wanted {
				t.Errorf("got %t, want %t", r, wanted)
			}
		})
	}
}

// Benchmark tests typically go in _test.go files and are named beginning with Benchmark.
// The testing runner executes each benchmark function several times,
// increasing b.N on each run until it collects a precise measurement.
func BenchmarkIf(b *testing.B) {
	for i := 0; i < b.N; i++ { // Typically the benchmark runs a function we’re benchmarking in a loop b.N times.
		If(1)
	}
}

func BenchmarkSwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Switch(1)
	}
}

func BenchmarkSwitchConst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SwitchConst(1)
	}
}
