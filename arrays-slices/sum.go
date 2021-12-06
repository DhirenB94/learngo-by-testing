package main

import "fmt"

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
		fmt.Println(sums)
	}
	return sums
} //append function --> takes a slice (sums) and a new value (previous function Sum), and returns a new slice with both these items on it

func main() {
	fmt.Println(SumAll([]int{1,2}, []int{7, 3}, []int{4, 2}, []int{5, 9}))
}



// range
// --lets you iterate over an array
// -- on each iteration, range returns 2 values, --> the index position and the value
// -- by using _, we are choosing to ignore the index position


//...
//variadic function
//allows multiple arguments


