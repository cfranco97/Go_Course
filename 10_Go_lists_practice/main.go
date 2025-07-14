package main

import (
	"errors"
	"fmt"
)

type Product struct {
	title string
	id    string
	price int
}

func new(title, id string, price int) (Product, error) {
	if title == "" || id == "" || price <= 0 {
		return Product{}, errors.New("invalid input")
	}
	return Product{
		title: title,
		id:    id,
		price: price,
	}, nil
}

func (product Product) Display() {
	fmt.Printf("7)Your product %v with the id %v costs:%v€\n", product.title, product.id, product.price)
}

func main() {
	//	1
	hobbies := [3]string{"Gaming", "Reading", "Music"}
	fmt.Println("1)Hobbies: ", hobbies)
	//	2
	fmt.Println("2) first element:", hobbies[0])
	fmt.Println("2) second and third elements:", hobbies[1:])
	//	3
	slicedHobbies := hobbies[:2]
	fmt.Println("3) slicedHobbies:", slicedHobbies)
	//	4
	slicedHobbies = hobbies[1:]
	fmt.Println("4) updated slicedHobbies:", slicedHobbies)
	//	5 & 6
	dynamicHobbies := []string{}
	dynamicHobbies = append(dynamicHobbies, "Go")
	dynamicHobbies = append(dynamicHobbies, "Coding")
	dynamicHobbies = append(dynamicHobbies, "Coding")
	fmt.Println("5&6) dynamicHobbies:", dynamicHobbies)
	//	7
	userProduct, err := new("Bananas", "1", 5) // hardcoded product input

	if err != nil {
		fmt.Print(err)
		return
	}
	userProduct.Display()
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
