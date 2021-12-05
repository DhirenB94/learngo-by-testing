package main

func Sum(numbers[5]int) int {

	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

// range
// --lets you iterate over an array
// -- on each iteration, range returns 2 values, --> the index position and the value
// -- by using _, we are choosing to ignore the index position