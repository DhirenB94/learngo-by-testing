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
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)
	//make --> allows us to create a slice with a starting capacity of the len of numbersToSum
	fmt.Println(sums)


	//[]int{1,2}, []int{7, 3}
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}
// iterate over the varying number of slices in the argument --> add them using our existing Sum func, then add these to a new slice we made

func main() {
	fmt.Println(SumAll([]int{1,2}, []int{7, 3}))
}

// range
// --lets you iterate over an array
// -- on each iteration, range returns 2 values, --> the index position and the value
// -- by using _, we are choosing to ignore the index position


//...
//variadic function
//allows multiple arguments


