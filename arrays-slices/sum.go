package main

func Sum(numbers []int) int {

	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	return
}

// range
// --lets you iterate over an array
// -- on each iteration, range returns 2 values, --> the index position and the value
// -- by using _, we are choosing to ignore the index position


//...
//variadic function
//allows multiple arguments


