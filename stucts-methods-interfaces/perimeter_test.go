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

	t.Run("Area of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{12.0, 5.0}
		got := Area(rectangle)
		want := 60.0

		if got != want {
			t.Errorf("Got area %.2f, but want area %.2f", got, want)
		}
	})

	t.Run("Area of a circle", func(t *testing.T) {

		circle := Circle {
			radius: 10.0,
		}
		got := Area(circle)
		want := 314.1592653589793

		if got != want {
			t.Errorf("Got area %g, but want area %g", got, want) //g prints a more precise decimal number
		}

	})

}
