package integers

import "testing"

func TestAdder(t *testing.T) {
	got := Add(2, 2)
	expected := 4

	if got != expected {
		t.Errorf("Expected '%d' but got '%d'", expected, got)
	}
}
