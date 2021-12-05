package main

import "testing"

func TestSum(t *testing.T) {

	t.Run("Collection of 5 numbers (array)", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15


		if got != want {
			t.Errorf("Got %d, Want %d, Given array %v", got, want, numbers)
		}
	})

	t.Run("Colletion of any size (slice)", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d, want %d, given slice %v", got, want, numbers)
		}

	})

}
