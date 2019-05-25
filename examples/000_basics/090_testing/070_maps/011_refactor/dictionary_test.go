package main

import "testing"

func TestSearch(t *testing.T) {
	// dictionary := map[string]string{"test": "this is just a test"}
	dictionary := Dictionary{"test": "this is just a test"}
	given := "test"
	got := dictionary.Search(given)
	want := "this is just a test"

	assertStrings(t, got, want, given)
}

func assertStrings(t *testing.T, got, want, given string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s' given, '%s'", got, want, given)
	}
}
