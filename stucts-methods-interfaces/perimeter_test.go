package perimeter

import "testing"

func TestPerimeter(t *testing.T) {

	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("Got %f.2, but want %f.2", got, want)
	} // f for float 64 and .2 for 2 decimal places
}

