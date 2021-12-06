package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15


		if got != want {
			t.Errorf("Got %d, Want %d, Given array %v", got, want, numbers)
		}
	})

	t.Run("Collection of any size (slice)", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d, want %d, given slice %v", got, want, numbers)
		}
	})

	t.Run("takes a varying number of slices and returns a new slice containing the totals for each slice passed", func(t *testing.T) {

		got := SumAll([]int{1,2}, []int{7, 3})
		want := []int{8, 5}


		//can't do if got != want as cannot use this when comparing slices
		//reflect.DeepEqual allows for the comparison of ANY 2 variables
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got slice %v, want slice %v", got, want)
		}
	})

}
