package main

func Sum(numbers []int) int {

	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
} //append function --> takes a slice (sums) and a new value (previous function Sum), and returns a new slice with both these items on it

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		tails:= numbers[1:] //slicing the numbers from and including index position 1 ---> The syntax is slice[low:high]
		sums = append(sums, Sum(tails))
	}
	return sums
}






// range
// --lets you iterate over an array
// -- on each iteration, range returns 2 values, --> the index position and the value
// -- by using _, we are choosing to ignore the index position


//...
//variadic function
//allows multiple arguments


