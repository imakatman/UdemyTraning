package main

import "fmt"

var metersToYards float64 = 1.09361

func main(){
	var meters float64
	fmt.Print("Enter amount of meters: ")
	fmt.Scan(&meters) // Scan will point at that memory address and enter what is inputted
	yards := meters * metersToYards
	fmt.Println(meters, "meters is", yards, "yards")
}