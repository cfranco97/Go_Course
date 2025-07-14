package transformers

import "fmt"

// create a custom type - to avoid verbose function lengths inside the generic function!
type transformFn func(int) int // receives: int ; returns: int

func main() {
	numbers := []int{1, 2, 3, 4}
	double := createTransformer(2)                    // created double function
	triple := createTransformer(3)                    // created triple function
	transformed := transformNumbers(&numbers, double) // don't put '()' so we can use double function as an argument instead of executing it!

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(transformed)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

// A Generic function that will return an array '[]int' - takes in a pointer to some numbers array and a function that will transform the numbers in said array.
func transformNumbers(numbers *[]int, transform transformFn) []int {
	tNumbers := []int{} // new empty array to return later
	for _, val := range *numbers {
		tNumbers = append(tNumbers, transform(val)) // recursive!
	}
	return tNumbers
}

// Function that creates an anonymous function - factor value is locked in once defined!
// fancy way of creating the multiplication function lol
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
