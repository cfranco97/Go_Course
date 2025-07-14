package recursions

// RECURSIONS
func main() {
	println(factorial(5))
}

// Using factorial function as an example, it calls itself until a certain condition is met.
func factorial(number int) int {
	if number == 0 {
		return 1
	}

	return number * factorial(number-1)
}
