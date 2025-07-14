package anonymous

import "fmt"

type transformFn func(int) int // receives: int ; returns: int

func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := transformNumbers(&numbers, func(number int) int { return number * 2 }) // example of an anonymous function - written just in time
	fmt.Println(doubled)
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	tNumbers := []int{} // new empty array to return later
	for _, val := range *numbers {
		tNumbers = append(tNumbers, transform(val)) // recursive!
	}
	return tNumbers
}
