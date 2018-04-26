package main

import "fmt"

func main() {
	data := []float64{4, 90, 11, 37}

	n := average(data)
	fmt.Println(n)
}

func average(sf []float64) float64 {
	fmt.Printf("%T \n", sf)

	var total float64

	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}