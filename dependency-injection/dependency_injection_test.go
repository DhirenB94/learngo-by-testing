package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	//buffer type from the bytes package, implements the Writer interface
	//so we use it to send in as our Writer

	//then we can check what was written to it after calling Greet
	Greet(&buffer, "Ziggy")

	got := buffer.String()
	want := "Hello, Ziggy"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
