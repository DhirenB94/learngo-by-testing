package main

import "testing"

/* func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got,want)
	}
} */

//when function is called wih an empty string, to default to priniting "Hello, world", rather than "Hello"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"

		if got!= want {
			t.Errorf("got %q want %q", got, want)
		}
	})                     //this part of the test still passes

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello ("")
		want := "Hello, world"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})          //but this part does not, since this is a subtest under one test, the overall test still fails
}

