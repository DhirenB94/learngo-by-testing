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
		name  string
		shape Shape
		hasArea  float64
	}{
		{"Rectangle", Rectangle{12.0, 5.0}, 60.0},
		{"Circle", Circle{10.0}, 314.1592653589793},
		{"Triangle", Triangle{10.0, 20.0}, 100.00},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g, has area %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}

//changes to make test easier to decipher and debug
//1. changing our error message to include %#v --> this allows us to see exactly which case has the error, otherwise we would only know there was an error in one of these cases
// here we are printing out the struct with its values in its fields (eg Triangle{10.0, 20.0})
//2. including a name field in our struct
//3. renaming the want field to hasArea
//4. wrapping each case in a t.run gives us a clearer test output as it will print the name of the case eg FAIL: TestArea/Rectangle (0.00s)
//5. using tt.name from the case to use it as the t.run test name








// TDD - table driven tests
// created an anonymous struct --> areaTests
// this struct is a slice of structs ([]struct) that has 2 fields, shape and want
// Then we fill the slices with cases - eg rectangle and circle, and then iterate through the slice of structs in the areaTests struct
// you can easily add another shape and add to the test list
// useful when you want to build a list of test cases that can be tested in the same manner