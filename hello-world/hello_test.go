package main

import "testing"


//when function is called wih an empty string, to default to priniting "Hello, world", rather than "Hello"


func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'Hello, world'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T) {
		got := Hello("Thierry", "French")
		want := "Bonjour, Thierry"
		assertCorrectMessage(t, got, want)
	})
}

// have assigned the test function (if got!= want part) to the assertCorrectMessage variable
// then used this variable instead of repeated if gotwant lines
//when using t.helper - use testing.TB as this gives access to both testing.T (test) and testing.B (benchmark)
//t.helper allows us to see exactly which line number the test fails at, comment it out to see the difference, it would fail within the test helper.

//addition of spanish requirement
// test fails because you pass 2 string arguements, but function only has one
// after making the hello func have 2 arguements, there now arent enough arguements in the other tests!
//