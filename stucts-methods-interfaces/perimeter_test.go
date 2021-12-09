package perimeter

import "testing"

func TestPerimeter(t *testing.T) {

	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("Got %.2f, but want %.2f", got, want)
	} // f for float 64 and .2 for 2 decimal places
}

