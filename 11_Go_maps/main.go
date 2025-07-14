package main

import "fmt"

//TYPE ALIAS
type floatMap map[string]float64

// func for custom type
func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)
	// search map by key (case-sensitive!)
	fmt.Println(websites["Amazon Web Services"])
	// adding new items = target a key that doesn't exist.
	// works the same when replacing an existing value from an existing key!
	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)
	// deleting a map
	delete(websites, "LinkedIn")
	fmt.Println(websites)

	// using map to avoid aditional memory alocation
	//courseRatings := make(map[string]float64, 3)
	courseRatings := make(floatMap, 3)
	courseRatings["Go"] = 4.7
	courseRatings["React"] = 4.8
	courseRatings["Angular"] = 4.7
	courseRatings.output()

	// FOR loops with collections (arrays,slices,maps)

	for index, value := range courseRatings {
		fmt.Println("Index:", index)
		fmt.Println("Value:", value)
	}
}
