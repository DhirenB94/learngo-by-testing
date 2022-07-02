package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	//want to have called the sleep() 3 times
	if spySleeper.Calls != 3 {
		t.Errorf("not enough calls to sleeper, want 3 got %d", spySleeper.Calls)
	}

}

//Why do we need to mock?
//our tests take 3 seconds to run every time - think about if the number was bigger thn 3
//we have a dependency on sleep - need to extract to be able to control - mock it
//if we can mock time.Sleep
//we can use dependency injection to use it instead of a "real" time.Sleep
//and then we can spy on the calls to make assertions on them.
