package main

import "fmt"

// Generic functions are good safety nets when you don't really know what type of input the user might input into your function (EX: library sharing).
func main() {
	result := add(1, 2.5)

	fmt.Println("Result:", result)
}

// This generic function, accepts the following types: int,float64 and string.
// Since those 3 types can make sum(+) calculatons, any argument introcuded into this generic function will work as long as their types are widthin those 3 said types.
func add[T int | float64 | string](a, b T) T {
	return a + b
}
