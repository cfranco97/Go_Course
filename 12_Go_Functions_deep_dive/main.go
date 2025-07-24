package main

import (
	"fmt"
)

// Variadic functions
func main() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := sumup(0, 2, 3, 4, 5)        // first value will be 1, the remaining will go into the numbers ... array
	anotherSum := sumup(0, numbers...) // still uses numbers array as input
	fmt.Println(sum)
	fmt.Println(anotherSum)
}

// Accepts different input arguments as long as the type is correct!
// The ... essencially merge all incoming int values into a slice/array (mouse over numbers argument to verify it!)
func sumup(startingval int, numbers ...int) int {
	sum := startingval

	for _, val := range numbers {
		sum += val
	}

	return sum
}
