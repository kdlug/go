package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		given := "test"
		got, _ := dictionary.Search(given)
		want := "this is just a test"

		assertStrings(t, got, want, given)
	})

	t.Run("unknown word", func(t *testing.T) {
		given := "unknown"
		_, err := dictionary.Search(given)

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertError(t, err, ErrNotFound)
	})

}

func assertStrings(t *testing.T, got, want, given string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s' given, '%s'", got, want, given)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
