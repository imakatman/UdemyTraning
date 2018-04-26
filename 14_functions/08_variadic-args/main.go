package main

import "fmt"

func main() {
	// The average function is expecting 0 or more items of a type float64 and data
	// is of a type float64 SLICE
	data := []float64{4, 90, 11, 37}

	// In order to pass in multiple parameters into the average function,
	// we need to follow the data variable with a ... because that will look
	// inside of data, which is a SLICE, and grab each individual item and pass
	// it into the function.
	n := average(data...)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	fmt.Printf("%T \n", sf)
	// Default value of float64 is 0.0
	var total float64
	// You use range to loop over a list of items, or collection of items
	// First parameter in this for loop is index but we don't need it so we provide a blank identifier
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}