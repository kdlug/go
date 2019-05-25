// Rules:
// - File should have a name like xxx_test.go
// - The test function must start with the word Test
// - The test function takes one argument only t *testing.T

// godoc -http :8010
// http://localhost:8010/pkg/testing/
package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		// We are calling the Errorf method on our t which will print out a message and fail the test
		t.Errorf("got '%s' want %s", got, want)
	}
}

func TestHelloYou(t *testing.T) {
	got := HelloYou("Mike")
	want := "Hello, Mike"

	if got != want {
		t.Errorf("got '%s' want %s", got, want)
	}
}

func TestHelloSubsets(t *testing.T) {
	// subset of tests - it group tests with different scenarios
	t.Run("saying hello to you", func(t *testing.T) {
		got := HelloYou("Mike")
		want := "Hello, Mike"

		if got != want {
			t.Errorf("got '%s' want %s", got, want)
		}
	})

	t.Run("saying hello with empty string", func(t *testing.T) {
		got := HelloYou("")
		want := "Hello, stranger"

		if got != want {
			t.Errorf("got '%s' want %s", got, want)
		}

	})
}

func TestHelloSubsetsRefactored(t *testing.T) {
	// we removed duplicated code
	// added assertion into a function
	assertCorrectMsg := func(t *testing.T, got, want string) {
		// t.Helper() is needed to tell the test suite that this method is a helper. By doing this when it fails the line number reported will be in our function call rather than inside our test helper. This will help other developers track down problems easier.
		t.Helper()

		if got != want {
			t.Errorf("got '%s' want %s", got, want)
		}
	}

	t.Run("saying hello to you", func(t *testing.T) {
		got := HelloYou("Mike")
		want := "Hello, Mike"

		assertCorrectMsg(t, got, want)
	})

	t.Run("saying hello with empty string", func(t *testing.T) {
		got := HelloYou("")
		want := "Hello, stranger"

		assertCorrectMsg(t, got, want)
	})
}

func TestHelloLang(t *testing.T) {
	assertCorrectMsg := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got '%s' want %s", got, want)
		}
	}

	t.Run("saying hello to you in spanish", func(t *testing.T) {
		got := HelloYouLang("Mike", "spanish")
		want := "Hola, Mike"

		assertCorrectMsg(t, got, want)
	})
}
