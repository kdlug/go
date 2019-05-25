package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	repeated := Repeat("a", 5)

	fmt.Println(repeated)
	// Output: aaaaa
}

// Benchmarking
// The testing.B gives you access to the cryptically named b.N.
// When the benchmark code is executed, it runs b.N times and measures how long it takes.
// go test -bench=.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
