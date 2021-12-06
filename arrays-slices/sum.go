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
	}
	return sums
} //append function --> takes a slice (sums) and a new value (previous function Sum), and returns a new slice with both these items on it

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums =append(sums, 0)
		} else {
		tails:= numbers[1:]
		sums = append(sums, Sum(tails))
		}
		fmt.Println(numbers, sums)
	}
	return sums
}

func main()  {
	fmt.Println(SumAllTails([]int{}, []int{1, 1, 1}, []int{2, 2, 2}, []int{3, 3, 3}))
}






// range
// --lets you iterate over an array
// -- on each iteration, range returns 2 values, --> the index position and the value
// -- by using _, we are choosing to ignore the index position


//...
//variadic function
//allows multiple arguments


