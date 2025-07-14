package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

// used for reference
/* func main() {
	var productNames [4]string                      // an array with 4 empty slots
	productNames = [4]string{"A book"}              // only 1 slot is filled
	productNames[2] = "A bottle"                    // index 2 filled
	prices := [4]float64{10.99, 9.99, 12.50, 11.99} // an array of size 4 of float64
	fmt.Println(prices)
	fmt.Println(prices[3]) // print index 3
	fmt.Println(productNames)

	// how to slice arrays
	//featuredPrices := prices[1:3] // select all elements between index 1 to 3
	//featuredPrices := prices[:3] // select all elements from the start to index 3
	featuredPrices := prices[1:] // select all elements between index 1 to the end
	fmt.Println(featuredPrices)

} */

// dynamic array example
func main() {
	prices := []float64{}
	prices = append(prices, 5)
	fmt.Println(prices)

	// Append all items of a list to another list
	newPrices := []float64{10, 5, 2}
	prices = append(prices, newPrices...)
	fmt.Println(prices)

	// MAKE function
	// useful when you already know the size of the array/slice you'll be working with, since Go recreates the array/slice every time an item is added into it.
	userNames := make([]string, 2, 5) // string array with 2 empty slots with a max capacity of 5.

	userNames[0] = "Julie"
	userNames = append(userNames, "John")
	userNames = append(userNames, "Maria")
	fmt.Println(userNames)
}
