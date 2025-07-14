package main

import "fmt"

// create a custom type - to avoid verbose function lengths inside the generic function!
type transformFn func(int) int // receives: int ; returns: int

func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := transformNumbers(&numbers, double) // don't put '()' so we can use double function as an argument instead of executing it!
	fmt.Println(doubled)
}

// A Generic function that will return an array '[]int' - takes in a pointer to some numbers array and a function that will transform the numbers in said array.
func transformNumbers(numbers *[]int, transform transformFn) []int {
	tNumbers := []int{} // new empty array to return later
	for _, val := range *numbers {
		tNumbers = append(tNumbers, transform(val)) // recursive!
	}
	return tNumbers
}

func double(number int) int {
	return number * 2
}
