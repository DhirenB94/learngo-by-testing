package perimeter

import "testing"

func TestPerimeter(t *testing.T) {

	rectangle := Rectangle{
		width:  10.0,
		height: 10.0,
	}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("Got %.2f, but want %.2f", got, want)
	} // f for float 64 and .2 for 2 decimal places
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12.0, 5.0}, 60.0},
		{Circle{10.0}, 314.1592653589793},
		{Triangle{base: 10.0, height: 20.0}, 100.00 },
		
	}

	for _, tt := range areaTests{
		got:= tt.shape.Area()
		if got != tt.want{
			 t.Errorf("got %g, want %g", got, tt.want)
		}
	}
}

// TDD - table driven tests
// created an anonymous struct --> areaTests
// this struct is a slice of structs ([]struct) that has 2 fields, shape and want
// Then we fill the slices with cases - eg rectangle and circle, and then iterate through the slice of structs in the areaTests struct
// you can easily add another shape and add to the test list
// useful when you want to build a list of test cases that can be tested in the same manner