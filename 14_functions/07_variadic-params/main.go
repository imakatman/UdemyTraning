package main

import "fmt"

func main() {
	n := average(4, 90, 11, 37)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	// sf the parameter is of a type float64 slice because it is expecting 0 or more float64 items
	fmt.Printf("%T \n", sf)
	// Default value of float64 is 0.0
	var total float64
	// You use range to loop over a list of items, or collection of items
	// First parameter in this for loop is index but we don't need it so we provide a blank identifier
	for _, v := range sf {
		fmt.Printf("%T \n", v)
		total += v
	}
	return total / float64(len(sf))
}