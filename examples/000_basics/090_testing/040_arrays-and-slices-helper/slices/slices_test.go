package slices

import (
	"reflect"
	"testing"
)

/*
Go's built-in testing toolkit features a coverage tool, which can help identify areas of your code you have not covered:
go test -cover
*/
func TestSum(t *testing.T) {
	t.Run("slice of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	/*
		./sum_test.go:26:9: invalid operation: got != want (slice can only be compared to nil)

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want)
		}
	*/

	// It's important to note that reflect.DeepEqual is not "type safe", the code will compile even if you did something a bit silly.
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	assertEqualSlices := func(t *testing.T, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("sum slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		assertEqualSlices(t, got, want)
	})

	t.Run("sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}

		assertEqualSlices(t, got, want)
	})
}

// go test -bench=.
func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 2}, []int{0, 9})
	}
}
