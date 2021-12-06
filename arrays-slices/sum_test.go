package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	check:= func(t testing.TB, numbers []int, got, want int) {
		if got != want {
			t.Errorf("got %d, want %d, given slice %v", got, want, numbers)
		}
	}
	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15
		check(t, numbers, got, want)
	})

	t.Run("Collection of any size (slice)", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6
		check(t, numbers, got, want)

	})

	checkSums:= func(t testing.TB, got, want[]int) { //variable t assigned the testing package -- got and want of type []int
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got slice %v, want slice %v", got, want)
		}
	}

	t.Run("takes a varying number of slices and returns a new slice containing the totals for each slice passed", func(t *testing.T) {

		got := SumAll([]int{1,2}, []int{7, 3})
		want := []int{3, 10}
		checkSums(t, got, want)
	})

	t.Run("Summing all the tails", func(t *testing.T) {
		//tails are all parts of a slice apart from the 1st(head), here we are summing all the tails of each slice
		got := SumAllTails([]int{1,2}, []int{7, 3})
		want := []int{2, 3}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slicea", func(t *testing.T) {
		got:= SumAllTails([]int{}, []int{3, 4, 5})
		want:= []int{0, 9}

		checkSums(t, got, want)

	})


}


//reflect.DeepEqual
//NOT type safe
//If we changed the want:= to a string, the test would still compile! Eventhough you are comparing a string to a slice