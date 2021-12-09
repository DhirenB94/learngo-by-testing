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

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
	got := shape.Area()
	if got != want {
		t.Errorf("got %g, want %g", got, want)
	}

	}
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 5.0}
		checkArea(t, rectangle, 60.0)

	})

	t.Run("circles", func(t *testing.T) {
		circle:= Circle{10.0}
		checkArea(t, circle, 314.1592653589793)
	})

	//Interfaces
	//We want to take a collection of shapes, call the Area() method on them and then check the result
	// So we can write a checkArea function that we can pass both Rectangles and Circles to, but fail to compile if we try to pass something that isnt a shape
	//Interfaces allow you to make functions that can be used with different types whilst still maintaining type safety

	//We are creating a helper function where we are asking for a Shape to be passed in - if something that isnt a shape is passed it will not clompile
	//now we just need to tell Go what a Shape interface is!
}
